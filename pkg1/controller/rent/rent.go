package rent

import (
	"context"

	"github.com/Kourin1996/micro-service-example/pkg1/pb"
	"github.com/Kourin1996/micro-service-example/shared/domain/rent"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type rentGRPCServer struct {
	rent gt.Handler
}

func NewGRPCServer(logger log.Logger, rentService rent.IRentService) pb.RentServiceServer {
	return &rentGRPCServer{
		rent: gt.NewServer(
			// handler
			func(ctx context.Context, request interface{}) (response interface{}, err error) {
				return rentService.Rent(ctx, request.(*rent.RentDTO))
			},
			// request encoder
			func(_ context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.RentRequest)

				return &rent.RentDTO{ID: req.Id, Memo: req.Memo}, nil
			},
			// response decoder
			func(_ context.Context, response interface{}) (interface{}, error) {
				resp := response.(*rent.RentResult)

				return &pb.RentResponse{Status: resp.Status}, nil
			},
		),
	}
}

func (s *rentGRPCServer) Rent(ctx context.Context, req *pb.RentRequest) (*pb.RentResponse, error) {
	_, resp, err := s.rent.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.RentResponse), nil
}
