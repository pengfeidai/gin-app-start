package main

import (
	"gin-app-start/config"
	"gin-app-start/router"
)

func main() {
	router := router.InitRouter()

	router.Run(config.Port)

	// 优雅关停
	// server := &http.Server{
	// 	Addr:    ":9060",
	// 	Handler: router,
	// }

	// go func() {
	// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Println("shutdown server...")

	// // 创建10s的超时上下文
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Fatal("server shutdown:", err)
	// }
	// log.Println("server exiting...")
}
