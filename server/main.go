package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vuluu2k/remember_fullstack/server/handler"
)

func main() {
	log.Println("Starting server...")
	router := gin.Default()

	handler.NewHandler(&handler.Config{R: router})

	svr := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Fail to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening in http://localhost%v", svr.Addr)

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	log.Println("Shutting down server...")

	if err := svr.Shutdown(ctx); err != nil {
		log.Fatalf("Fail to shutdown server: %v\n", err)
	}
}
