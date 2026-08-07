package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go_project/config"
	"go_project/router"

	"github.com/gin-gonic/gin"
)

const addr = ":8008"

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.InitDB()

	srv := &http.Server{
		Addr:    addr,
		Handler: router.New(),
	}

	go func() {
		log.Println("服务启动:", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("服务启动失败:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务关闭失败:", err)
	}
	log.Println("服务已退出")
}
