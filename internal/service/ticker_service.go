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
	count   int
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
			for _, symbol := range s.Symbols {
				price, err := s.repo.FetchPrice(symbol)
				if err != nil {
					fmt.Printf("Error fetching price for %s: %v\n", symbol, err)
					continue
				}
				s.count++
				if prevPrice, exists := s.prices[symbol]; exists {
					if prevPrice != price {
						fmt.Printf("%s price: %s changed\n", symbol, price)
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
