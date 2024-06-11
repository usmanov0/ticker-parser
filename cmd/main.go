package main

import (
	"awesomeProject/internal/adapter"
	"awesomeProject/internal/config"
	"awesomeProject/internal/input"
	"awesomeProject/internal/monitor"
	"awesomeProject/internal/service"
	"awesomeProject/internal/worker"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	repo := adapter.NewBinanceRepository()
	services := service.CreateServices(repo, config.Symbols, config.MaxWorkers)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	worker.RunWorkers(services, ctx, &wg)

	go monitor.MonitorRequests(services, ctx)
	go input.WaitForStop(ctx, cancel)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()

	wg.Wait()
}
