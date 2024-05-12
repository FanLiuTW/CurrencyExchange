package main

import (
	"Asiayo/config"
	"Asiayo/routes"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "./log/log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
	})
	log.SetOutput(mw)
	config.Initialize(".env")
}

func main() {
	gin.SetMode(config.RunMode)
	port := config.Port
	routesInit := routes.InitRouter()
	server := &http.Server{
		Addr:           port,
		Handler:        routesInit,
		ReadTimeout:    time.Duration(config.ReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(config.WriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	idleConnections := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Server shutdown failed: %v", err)
		}
		log.Println("Server gracefully shutdown")
		close(idleConnections)
	}()

	log.Println("[HttpServer]", port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server not gracefully shutdown: %v", err)
	}

	<-idleConnections
}
