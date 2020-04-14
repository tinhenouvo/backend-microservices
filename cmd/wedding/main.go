package main

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"agungdwiprasetyo.com/backend-microservices/config"
	"agungdwiprasetyo.com/backend-microservices/internal/app"
	"agungdwiprasetyo.com/backend-microservices/internal/services/wedding"
)

const (
	appLocation = "cmd/wedding"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer func() {
		cancel()
		if r := recover(); r != nil {
			fmt.Println("Failed to start wedding service:", r)
			fmt.Printf("Stack trace: \n%s\n", debug.Stack())
		}
	}()

	cfg := config.Init(ctx, appLocation)
	defer cfg.Exit(ctx)

	service := wedding.NewService(cfg)
	app.New(service).Run(ctx)
}
