package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-app-start/internal/config"
	"gin-app-start/internal/controller"
	"gin-app-start/internal/model"
	"gin-app-start/internal/repository"
	"gin-app-start/internal/router"
	"gin-app-start/internal/service"
	"gin-app-start/pkg/database"
	"gin-app-start/pkg/logger"

	_ "gin-app-start/docs"

	"go.uber.org/zap"
)

//	@title			Gin App API
//	@version		1.0
//	@description	This is a RESTful API server built with Gin framework.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:9060
//	@BasePath	/

//	@schemes	http https

var Version string

func main() {
	log.Printf("Version: %s\n", Version)

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := logger.Init(cfg.Server.Mode, cfg.Log.FilePath); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Application starting", zap.String("version", Version), zap.String("mode", cfg.Server.Mode))

	db, err := database.NewPostgresDB(&database.PostgresConfig{
		Host:         cfg.Database.Host,
		Port:         cfg.Database.Port,
		User:         cfg.Database.User,
		Password:     cfg.Database.Password,
		DBName:       cfg.Database.DBName,
		SSLMode:      cfg.Database.SSLMode,
		MaxIdleConns: cfg.Database.MaxIdleConns,
		MaxOpenConns: cfg.Database.MaxOpenConns,
		MaxLifetime:  cfg.Database.MaxLifetime,
		LogLevel:     cfg.Database.LogLevel,
	})
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer database.Close()

	logger.Info("Database connected successfully")

	if cfg.Database.AutoMigrate {
		if err := db.AutoMigrate(&model.User{}); err != nil {
			logger.Fatal("Database migration failed", zap.Error(err))
		}
		logger.Info("Database migration completed")
	}

	_, err = database.NewRedisClient(&database.RedisConfig{
		Addr:         cfg.Redis.Addr,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
		MaxRetries:   cfg.Redis.MaxRetries,
	})
	if err != nil {
		logger.Warn("Failed to initialize Redis", zap.Error(err))
	} else {
		defer database.CloseRedis()
		logger.Info("Redis connected successfully")
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	healthController := controller.NewHealthController()

	r := router.SetupRouter(healthController, userController, cfg)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}

	go func() {
		appURL := fmt.Sprintf("http://localhost:%d", cfg.Server.Port)
		swaggerURL := fmt.Sprintf("http://localhost:%d/swagger/index.html", cfg.Server.Port)

		logger.Info("Server started", zap.String("url", appURL))
		logger.Info("Swagger documentation", zap.String("url", swaggerURL))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", zap.Error(err))
	}

	logger.Info("Server stopped")
}
