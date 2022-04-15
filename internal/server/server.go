package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mustafa0831/TSarka/internal/counter"
	"github.com/Mustafa0831/TSarka/internal/email"
	gohttp "github.com/Mustafa0831/TSarka/internal/server/delivery/http"
	"github.com/Mustafa0831/TSarka/pkg/redisclient"
)

var flagPort = flag.String("port", "8080", "specify the port")

//Run initializes whole application
func Run() {
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	redisClient := redisclient.NewClient("localhost:7172", "")
	defer redisClient.Close()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal(fmt.Errorf("redis not connection: %s", err))
	}

	counterService := counter.NewService(redisClient)
	emailService := email.NewService()

	handler := gohttp.NewHandler(counterService, emailService)
	
	httpServer := NewServer(handler.Init())

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	//Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("failed to stop server: %v", err)
	}
}

//NewServer ...
func NewServer(handler http.Handler) http.Server {
	return http.Server{
		Addr:           fmt.Sprintf(":%s", *flagPort),
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
