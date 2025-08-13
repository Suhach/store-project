package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Suhach/fasion-store/backend/internal/database/postgres"
	"github.com/Suhach/fasion-store/backend/pkg/logger"
	"github.com/Suhach/fasion-store/backend/pkg/prometheus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	// Catch panic if we have
	defer func() {
		if err := recover(); err != nil {
			logger.Log.Warn("recover from panic", zap.Any("error", err))
		}
	}()
	// init Gin
	r := gin.New()

	// init Zap
	if err := logger.Init(); err != nil {
		log.Fatalf("logger init error: %v", err)
	}
	defer logger.Log.Sync()
	logger.Log.Info("logger init success", zap.String("app", "main"))

	// init prometheus
	prometheus.InitMetrics()
	r.Use(prometheus.PrometheusMiddleware())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// TODO: init KAFKA

	// TODO: init Redis

	// init PostgreSQL
	err = postgres.Init()
	if err != nil {
		logger.Log.Fatal("PostgreSQL init error", zap.Error(err))
	}

	// TODO: init dependencies
	/*
		DO wire
	*/

	// TODO: init openapi RegisterHandlers
	// openapi.RegisterHandlers(r, handler)

	// init server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		logger.Log.Info("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Server error", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logger.Log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Error("Server forced to shutdown", zap.Error(err))
	} else {
		logger.Log.Info("Server exited gracefully")
	}
}
