// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package chain

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicishero/NaiL/server/internal/conf"
	"github.com/vicishero/NaiL/server/pkg/app"
	"github.com/vicishero/NaiL/server/pkg/xerror"
)

// RateLimit 限流中间件
func RateLimit() gin.HandlerFunc {
	rc := redisCache()
	cfg := conf.RateLimitSetting

	return func(c *gin.Context) {
		// 如果限流未启用，直接放行
		if cfg == nil || !cfg.Enable {
			c.Next()
			return
		}

		// 获取请求路径
		path := c.Request.URL.Path

		// OSS/静态资源请求不做任何限制
		if strings.HasPrefix(path, "/oss/") {
			c.Next()
			return
		}

		// 获取客户端IP
		clientIP := getClientIP(c)

		// 获取用户ID
		var userID int64
		if uid, exist := c.Get("UID"); exist {
			userID = uid.(int64)
		}

		// 匹配限流规则，默认使用全局配置
		limit := cfg.DefaultLimit
		duration := cfg.DefaultDur
		dimension := "ip+path" // 默认维度：IP+路径

		// 查找自定义规则
		for _, rule := range cfg.Rules {
			// 支持通配符匹配，例如"/v1/*"匹配所有/v1开头的路径
			if matchRule(rule.Key, path) {
				limit = rule.Limit
				duration = rule.Duration
				dimension = rule.Dimension
				break
			}
		}

		// 构建限流Key
		var key string
		switch dimension {
		case "ip":
			key = fmt.Sprintf("ip:%s", clientIP)
		case "user":
			if userID > 0 {
				key = fmt.Sprintf("user:%d", userID)
			} else {
				// 用户未登录时回退到IP限流
				key = fmt.Sprintf("ip:%s", clientIP)
			}
		case "path":
			key = fmt.Sprintf("path:%s", path)
		case "ip+path":
			key = fmt.Sprintf("ip:%s:path:%s", clientIP, path)
		case "user+path":
			if userID > 0 {
				key = fmt.Sprintf("user:%d:path:%s", userID, path)
			} else {
				// 用户未登录时回退到IP+路径限流
				key = fmt.Sprintf("ip:%s:path:%s", clientIP, path)
			}
		default:
			// 未知维度默认使用IP+路径限流
			key = fmt.Sprintf("ip:%s:path:%s", clientIP, path)
		}

		// 获取限流许可
		ctx := context.Background()
		allowed, current, err := rc.RateLimitAcquire(ctx, key, limit, duration)

		if err != nil {
			// Redis出错时放行，避免影响业务
			c.Next()
			return
		}

		if !allowed {
			response := app.NewResponse(c)
			response.ToErrorResponse(xerror.TooManyRequests)
			c.Abort()
			return
		}

		// 设置限流相关响应头
		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", limit))
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", limit-current))
		c.Header("X-RateLimit-Reset", fmt.Sprintf("%d", time.Now().Add(duration).Unix()))

		c.Next()
	}
}

// getClientIP 获取客户端真实IP，防止伪造
func getClientIP(c *gin.Context) string {
	// 先获取直连IP
	remoteIP, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		remoteIP = c.Request.RemoteAddr
	}

	// 只有直连IP是可信代理时，才信任X-Forwarded-For头
	if isTrustedProxy(remoteIP) {
		// 优先从X-Forwarded-For获取
		if ip := c.GetHeader("X-Forwarded-For"); ip != "" {
			parts := strings.Split(ip, ",")
			if len(parts) > 0 {
				return strings.TrimSpace(parts[0])
			}
		}

		// 其次从X-Real-IP获取
		if ip := c.GetHeader("X-Real-IP"); ip != "" {
			return ip
		}
	}

	// 否则直接返回直连IP
	return remoteIP
}

// isTrustedProxy 判断是否是可信代理
func isTrustedProxy(ip string) bool {
	// 可信代理列表，默认包括内网IP段和本地回环地址
	trustedProxies := []string{
		"127.0.0.1",
		"::1",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	for _, cidr := range trustedProxies {
		_, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			// 如果不是CIDR格式，直接比较IP
			if cidr == ip {
				return true
			}
			continue
		}
		if ipNet.Contains(parsedIP) {
			return true
		}
	}
	return false
}

// matchRule 规则匹配，支持简单的通配符*
func matchRule(pattern, path string) bool {
	// 完全匹配
	if pattern == path {
		return true
	}

	// 前缀匹配，例如"/v1/*"匹配所有/v1开头的路径
	if strings.HasSuffix(pattern, "/*") {
		prefix := strings.TrimSuffix(pattern, "/*")
		return strings.HasPrefix(path, prefix)
	}

	// 后缀匹配，例如"*.jpg"匹配所有.jpg结尾的路径
	if strings.HasPrefix(pattern, "*") {
		suffix := strings.TrimPrefix(pattern, "*")
		return strings.HasSuffix(path, suffix)
	}

	return false
}
