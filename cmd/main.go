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

	repo := adapter.NewBinanceRepository()                                      // не понял почему это репо, если это клиент
	services := service.CreateServices(repo, config.Symbols, config.MaxWorkers) // неудачное название - непонятно какие сервисы создаем

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	worker.RunWorkers(services, ctx, &wg)

	// в целом слишком много пакетов - разве воркер и мониторинг это не сервисы - как будто была плохая попытка замутить DDD (но потом некоторые сущности поехали в отдельные пакеты)
	go monitor.MonitorRequests(services, ctx)
	go input.WaitForStop(ctx, cancel)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) //YAGNY
	go func() {
		<-sigChan
		cancel()
	}()

	wg.Wait()
}
