package service

import (
	"awesomeProject/internal/domain"
	"context"
	"fmt"
	"log"
	"sync"
)

type TickerService struct {
	repo    domain.TickerRepository
	Symbols []string
	count   int //  у тебя конкурентный доступ - стоит скрывать за мьютексом или использовать атомик
	prices  map[string]string
}

func NewTickerService(repo domain.TickerRepository) *TickerService {
	return &TickerService{
		repo:   repo,
		prices: make(map[string]string),
	}
}

func (s *TickerService) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping worker")
			return
		default:
			for _, symbol := range s.Symbols { // лучше селект ставить внутри цикла, а то если символов много, придется ждать пока они все отработают прежде чем выйдешь по контектсу
				price, err := s.repo.FetchPrice(symbol)
				if err != nil {
					fmt.Printf("Error fetching price for %s: %v\n", symbol, err)
					continue
				}
				s.count++                                          // в условии задачи сказано, что каунтер считает и неудачные запросы
				if prevPrice, exists := s.prices[symbol]; exists { // это можно и одним if else написать
					if prevPrice != price {
						fmt.Printf("%s price: %s changed\n", symbol, price) // в условии задачи сказано, что нельзя писать в консоль тут (кроме ошщибок)
					} else {
						fmt.Printf("%s price: %s\n", symbol, price)
					}
				} else {
					fmt.Printf("%s price: %s\n", symbol, price)
				}
				s.prices[symbol] = price
			}
		}
	}
}

func (s *TickerService) GetRequestsCount() int {
	return s.count
}
