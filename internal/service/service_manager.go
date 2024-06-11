package service

import "awesomeProject/internal/domain"

func CreateServices(repo domain.TickerRepository, symbols []string, maxWorkers int) []*TickerService {
	if maxWorkers > 2 {
		maxWorkers = 2
	}

	symbolsPerWorker := len(symbols) / maxWorkers
	services := make([]*TickerService, maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		start := i * symbolsPerWorker
		end := start + symbolsPerWorker
		if i == maxWorkers-1 {
			end = len(symbols)
		}
		services[i] = NewTickerService(repo)
		services[i].Symbols = symbols[start:end]
	}
	return services
}
