package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	rc "github.com/Kourin1996/micro-service-example/pkg1/controller/rent"
	"github.com/Kourin1996/micro-service-example/pkg1/pb"
	rr "github.com/Kourin1996/micro-service-example/pkg1/repository/rent"
	rs "github.com/Kourin1996/micro-service-example/pkg1/service/rent"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

func NewTerminateSignalCh() chan error {
	errCh := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errCh <- fmt.Errorf("%s", <-c)
	}()

	return errCh
}

func main() {
	logger := log.NewJSONLogger(os.Stdout)

	rentRepo := rr.NewRentRepository()
	rentService := rs.NewService(logger, rentRepo)
	rentServer := rc.NewGRPCServer(logger, rentService)

	errCh := NewTerminateSignalCh()

	grpcListener, err := net.Listen("tcp", ":10000")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterRentServiceServer(baseServer, rentServer)

		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errCh)

}
