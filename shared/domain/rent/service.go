package rent

import "context"

type IRentService interface {
	Rent(context.Context, *RentDTO) (*RentResult, error)
}
