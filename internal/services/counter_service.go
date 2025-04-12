package services

import (
	"context"

	"github.com/thyms-c/be-memo-app/internal/models"
	"github.com/thyms-c/be-memo-app/internal/repositories"
)

type CounterService interface {
	GetCounterByUserRole(ctx context.Context, name string) (*models.Counter, error)
}
type counterServiceImpl struct {
	counterRepository repositories.CounterRepository
}

func NewCounterService(counterRepository repositories.CounterRepository) CounterService {
	return &counterServiceImpl{
		counterRepository: counterRepository,
	}
}

func (c *counterServiceImpl) GetCounterByUserRole(ctx context.Context, name string) (*models.Counter, error) {
	counter, err := c.counterRepository.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if counter == nil {
		counter = &models.Counter{
			Name:  name,
			Value: 0,
		}
	}

	return counter, nil
}
