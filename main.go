package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime/debug"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Sprintf("main has error %s and stack %s", err, string(debug.Stack())))
			os.Exit(1)
		}
		os.Exit(0)
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Shutdown Server ...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}
