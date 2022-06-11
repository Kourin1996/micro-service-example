package rent

import (
	"context"
)

type IRentRepository interface {
	Rent(context.Context, int32, string) (int32, error)
}
