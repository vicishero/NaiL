// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// dailyRotateWriter 按日期自动轮转的日志写入器。
// 每天零点自动创建新文件，旧文件通过 lumberjack 按大小和保留天数管理。
type dailyRotateWriter struct {
	mu       sync.Mutex
	dir      string
	prefix   string
	ext      string
	maxSize  int  // MB
	maxAge   int  // days
	currentDay string
	writer    io.WriteCloser
}

func newDailyRotateWriter(dir, prefix, ext string, maxSize, maxAge int) io.Writer {
	w := &dailyRotateWriter{
		dir:     dir,
		prefix:  prefix,
		ext:     ext,
		maxSize: maxSize,
		maxAge:  maxAge,
	}
	_ = w.rotate()
	return w
}

func (w *dailyRotateWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	today := time.Now().Format("2006-01-02")
	if today != w.currentDay {
		_ = w.rotate()
	}

	return w.writer.Write(p)
}

func (w *dailyRotateWriter) rotate() error {
	if w.writer != nil {
		w.writer.Close()
	}

	today := time.Now().Format("2006-01-02")
	w.currentDay = today

	if err := os.MkdirAll(w.dir, 0o755); err != nil {
		return fmt.Errorf("dailyRotateWriter: create log dir %s: %w", w.dir, err)
	}

	filename := filepath.Join(w.dir, fmt.Sprintf("%s-%s%s", w.prefix, today, w.ext))

	w.writer = &lumberjack.Logger{
		Filename:  filename,
		MaxSize:   w.maxSize,
		MaxAge:    w.maxAge,
		LocalTime: true,
	}
	return nil
}
