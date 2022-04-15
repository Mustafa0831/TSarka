package counter

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Mustafa0831/TSarka/pkg/util"
	"github.com/go-redis/redis/v8"
)

//Service ...
type Service struct {
	Repository Counter
}

//NewService ...
func NewService(client *redis.Client) *Service {
	return &Service{
		Repository: NewCounterRepository(client),
	}
}

//CounterService ...
type CounterService interface {
	Get(ctx context.Context) (string, error)
	SetIncrement(ctx context.Context, num string) error
	SetDecrement(ctx context.Context, num string) error
}

//SetIncrement ...
func (s *Service) SetIncrement(ctx context.Context, num string) error {
	numUint, err := strconv.ParseUint(num, 10, 32)
	if err != nil {
		return err
	}
	numUint = util.Increment(numUint)
	return s.Repository.Set(ctx, numUint)
}

//SetDecrement ...
func (s *Service) SetDecrement(ctx context.Context, num string) error {
	numUint, err := strconv.ParseUint(num, 10, 32)
	if err != nil {
		return err
	}
	numUint = util.Decrement(numUint)
	return s.Repository.Set(ctx, numUint)
}

//Get ...
func (s *Service) Get(ctx context.Context) (string, error) {
	val, err := s.Repository.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("Get repository %s", err)
	}

	_, err = strconv.ParseUint(val, 10, 32)
	if err != nil {
		return "", fmt.Errorf("strconv.ParseUnit %s", err)
	}
	return val, nil
}
