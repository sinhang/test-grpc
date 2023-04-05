package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"test_grpc/pb"
)

func main() {
	//conn, err := g
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := pb.NewKaspawalletdClient(conn)

	ret, err := c.NewAddress(context.Background(), &pb.NewAddressRequest{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ret.Address)
}
