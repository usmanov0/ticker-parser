package worker

import (
	"awesomeProject/internal/service"
	"context"
	"sync"
)

// не понял зачем отдельный пакет для этого - это же обязанность сервисам - у тебя есть тикер сервис
func RunWorkers(services []*service.TickerService, ctx context.Context, wg *sync.WaitGroup) {
	for _, svc := range services {
		wg.Add(1)
		go svc.Run(ctx, wg)
	}
}
