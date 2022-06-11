package rent

import (
	"context"

	"github.com/Kourin1996/micro-service-example/shared/domain/rent"
	"github.com/go-kit/kit/log"
)

type rentService struct {
	logger   log.Logger
	rentRepo rent.IRentRepository
}

func NewService(
	logger log.Logger,
	rentRepo rent.IRentRepository,
) rent.IRentService {
	return &rentService{
		logger:   logger,
		rentRepo: rentRepo,
	}
}

func (s *rentService) Rent(ctx context.Context, dto *rent.RentDTO) (*rent.RentResult, error) {
	status, err := s.rentRepo.Rent(ctx, dto.ID, dto.Memo)
	if err != nil {
		return nil, err
	}

	return &rent.RentResult{
		Status: status,
	}, nil
}
