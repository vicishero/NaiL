// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源（开发环境）
	},
}

// WsMessage WebSocket 消息结构
type WsMessage struct {
	Type string      `json:"type"` // 消息类型: notification, status, heartbeat
	Data interface{} `json:"data"` // 消息数据
	Time int64       `json:"time"` // 时间戳
}

// WsClient WebSocket 客户端
type WsClient struct {
	ID     string
	Conn   *websocket.Conn
	Send   chan []byte
	Hub    *WsHub
	mu     sync.Mutex
}

// WsHub WebSocket 连接管理中心
type WsHub struct {
	clients    map[*WsClient]bool
	broadcast  chan []byte
	register   chan *WsClient
	unregister chan *WsClient
	mu         sync.RWMutex
	onlineCount int
}

var (
	hubInstance *WsHub
	hubOnce     sync.Once
)

// GetWsHub 获取WebSocket Hub单例
func GetWsHub() *WsHub {
	hubOnce.Do(func() {
		hub := &WsHub{
			clients:    make(map[*WsClient]bool),
			broadcast:  make(chan []byte, 256),
			register:   make(chan *WsClient),
			unregister: make(chan *WsClient),
		}
		hubInstance = hub
		go hub.Run()
	})
	return hubInstance
}

// Run Hub运行循环
func (h *WsHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.onlineCount = len(h.clients)
			h.mu.Unlock()
			logrus.Infof("WebSocket client connected: %s (total: %d)", client.ID, h.onlineCount)

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
			h.onlineCount = len(h.clients)
			h.mu.Unlock()
			logrus.Infof("WebSocket client disconnected: %s (total: %d)", client.ID, h.onlineCount)

		case message := <-h.broadcast:
			h.mu.RLock()
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// BroadcastMessage 广播消息到所有客户端
func (h *WsHub) BroadcastMessage(msgType string, data interface{}) {
	msg := WsMessage{
		Type: msgType,
		Data: data,
		Time: time.Now().Unix(),
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		logrus.Errorf("Failed to marshal ws message: %v", err)
		return
	}
	h.broadcast <- msgBytes
}

// GetOnlineCount 获取在线客户端数量
func (h *WsHub) GetOnlineCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.onlineCount
}

// readPump 从WebSocket连接读取消息
func (c *WsClient) readPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("WebSocket error: %v", err)
			}
			break
		}
		// 处理客户端消息（心跳回复等）
		var msg WsMessage
		if err := json.Unmarshal(message, &msg); err == nil && msg.Type == "heartbeat" {
			// 心跳应答
			resp := WsMessage{Type: "heartbeat", Data: "pong", Time: time.Now().Unix()}
			respBytes, _ := json.Marshal(resp)
			c.Send <- respBytes
		}
	}
}

// writePump 向WebSocket连接写入消息
func (c *WsClient) writePump() {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// WebSocketHandler WebSocket连接处理
// @Summary WebSocket连接
// @Tags WebSocket
// @Description 建立WebSocket连接，接收实时通知和状态更新
// @Router /ws [get]
func (s *AuthServant) WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Errorf("WebSocket upgrade failed: %v", err)
		return
	}

	hub := GetWsHub()
	clientID := c.Query("client_id")
	if clientID == "" {
		clientID = c.Request.RemoteAddr
	}

	client := &WsClient{
		ID:   clientID,
		Conn: conn,
		Send: make(chan []byte, 64),
		Hub:  hub,
	}

	hub.register <- client

	// 发送欢迎消息
	welcome := WsMessage{
		Type: "connected",
		Data: gin.H{
			"client_id":    clientID,
			"online_count": hub.GetOnlineCount(),
			"server_time":  time.Now().Unix(),
		},
		Time: time.Now().Unix(),
	}
	welcomeBytes, _ := json.Marshal(welcome)
	client.Send <- welcomeBytes

	// 启动读写协程
	go client.writePump()
	go client.readPump()
}

// NotifyWebSocket 向所有WebSocket客户端发送通知
func NotifyWebSocket(msgType string, data interface{}) {
	if hub := GetWsHub(); hub != nil {
		hub.BroadcastMessage(msgType, data)
	}
}
