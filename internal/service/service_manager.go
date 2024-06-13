package service

import "awesomeProject/internal/domain"

func CreateServices(repo domain.TickerRepository, symbols []string, maxWorkers int) []*TickerService {
	if maxWorkers > 2 { // мы же должны сравнить с количеством ЦПУ а не с 2
		maxWorkers = 2
	}

	symbolsPerWorker := len(symbols) / maxWorkers // если указать макс воркерс = 0 то словишь панику
	services := make([]*TickerService, maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		start := i * symbolsPerWorker
		end := start + symbolsPerWorker
		if i == maxWorkers-1 {
			end = len(symbols) // если символов 11, а воркеров 7 - распределит неравномерно
		}
		services[i] = NewTickerService(repo) // я бы все же разделил деление символов на группы и создание сервисов - как минимум легче писать юнит тесты отдельно
		services[i].Symbols = symbols[start:end]
	}
	return services
}
