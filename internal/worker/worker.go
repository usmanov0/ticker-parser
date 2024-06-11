package worker

import (
	"awesomeProject/internal/service"
	"context"
	"sync"
)

func RunWorkers(services []*service.TickerService, ctx context.Context, wg *sync.WaitGroup) {
	for _, svc := range services {
		wg.Add(1)
		go svc.Run(ctx, wg)
	}
}
