package main

import (
	"context"
	"hello-blog/internal/router"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	api := router.NewApiHandler(":8080")
	go api.Server()

	// 创建信号通道监听中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 等待中断信号
	<-quit
	log.Println("Shutting down server...")

	// 优雅关闭服务器，设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	api.Stop(ctx)

	log.Println("Server exited")
}
