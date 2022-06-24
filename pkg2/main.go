package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/Kourin1996/micro-service-example/pkg1/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// クライアント側はサーバーの証明書に対するCAの証明書をロード
	pemServerCA, err := ioutil.ReadFile("./cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// クライアント側の証明書と秘密鍵をロード
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	fmt.Printf("Start Server2\n")

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	conn, err := grpc.Dial(
		":10000",
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cleint := pb.NewRentServiceClient(conn)

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := cleint.Rent(timeoutCtx, &pb.RentRequest{
		Id:   1,
		Memo: "hello",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result=%d\n", res.Status)
}
