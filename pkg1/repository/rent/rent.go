package rent

import (
	"context"
	"fmt"

	"github.com/Kourin1996/micro-service-example/shared/domain/rent"
)

type RentRepository struct {
	count int32
}

func NewRentRepository() rent.IRentRepository {
	return &RentRepository{
		count: 0,
	}
}

func (r *RentRepository) Rent(ctx context.Context, id int32, memo string) (int32, error) {
	fmt.Printf("New Rent id=%d, memo=%s\n", id, memo)
	r.count++

	return r.count, nil
}
