package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Kourin1996/micro-service-example/pkg1/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Printf("Start Server2\n")

	conn, err := grpc.Dial(
		":10000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
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
