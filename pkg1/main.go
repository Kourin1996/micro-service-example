package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
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
	"google.golang.org/grpc/credentials"
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

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// CAの証明書をロード
	pemClientCA, err := ioutil.ReadFile("./cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// サーバー側の証明書と秘密鍵をロード
	serverCert, err := tls.LoadX509KeyPair("./cert/server-cert.pem", "./cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	logger := log.NewJSONLogger(os.Stdout)

	rentRepo := rr.NewRentRepository()
	rentService := rs.NewService(logger, rentRepo)
	rentServer := rc.NewGRPCServer(logger, rentService)

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		logger.Log("cannot load TLS credentials: ", err)

		return
	}

	errCh := NewTerminateSignalCh()

	grpcListener, err := net.Listen("tcp", ":10000")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer(
			grpc.Creds(tlsCredentials),
			// grpc.UnaryInterceptor(interceptor.Unary()),
			// grpc.StreamInterceptor(interceptor.Stream()),
		)
		pb.RegisterRentServiceServer(baseServer, rentServer)

		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errCh)

}
