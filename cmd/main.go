package main

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"spell-checker/internal/config"
	"syscall"
	"time"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// configure logger
	logger := zap.NewExample()
	defer logger.Sync()

	// configure router
	bootstrap := config.Bootstrap{}
	ginRouter := bootstrap.Setup(logger)

	logger.Sugar().Info("Try to start server...")
	server := http.Server{Addr: ":8080", Handler: ginRouter}

	// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
	go func() {
		logger.Sugar().Infof("Server started working on port: %s", "8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Sugar().Fatalf("Error occurred. Server closed. Error: %s", err)
		}
	}()
	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	logger.Sugar().Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Sugar().Fatal("Server forced to shutdown: ", err)
	}
	logger.Sugar().Info("Server exiting")
}
