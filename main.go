package main

import (
	"context"
	"fmt"
	"gin-app-start/app/config"
	"gin-app-start/app/database/mysql"
	"gin-app-start/app/database/redis"
	"gin-app-start/app/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 配置文件初始化
	config.Init()

	// mysql初始化
	db := mysql.Init()
	defer db.Close()

	if config.Conf.Redis.Use {
		// 初始化redis服务
		redis.Init()
	}

	RunServer()
}

func RunServer() {
	router := router.InitRouter()

	// router.Run(config.Conf.Server.Port)

	// 优雅关停
	server := &http.Server{
		Addr:         config.Conf.Server.Port,
		Handler:      router,
		ReadTimeout:  config.Conf.Server.ReadTimeout,
		WriteTimeout: config.Conf.Server.WriteTimeout,
	}

	log.Println(fmt.Sprintf("Listening and serving HTTP on Port: %s, Pid: %d", config.Conf.Server.Port, os.Getpid()))

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	log.Println("shutdown server...")

	// 创建5s的超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting...")
}
