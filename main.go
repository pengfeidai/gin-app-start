package main

import (
	"context"
	"fmt"
	"gin-app-start/app/database/mysql"
	"gin-app-start/app/router"
	"gin-app-start/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := router.InitRouter()

	// mysql初始化
	db := mysql.Init()
	defer db.Close()

	// router.Run(config.Port)

	// 优雅关停
	server := &http.Server{
		Addr:         ":9060",
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	log.Println(fmt.Sprintf("Listening and serving HTTP on Port: %s, Pid: %d", config.Port, os.Getpid()))

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	log.Println("shutdown server...")

	// 创建10s的超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting...")
}
