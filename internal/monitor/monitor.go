package monitor

import (
	"awesomeProject/internal/service"
	"context"
	"fmt"
	"time"
)

func MonitorRequests(services []*service.TickerService, ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			totalRequests := 0
			for _, svc := range services {
				totalRequests += svc.GetRequestsCount()
			}
			fmt.Printf("workers requests total: %d\n", totalRequests)
		}
	}
}
