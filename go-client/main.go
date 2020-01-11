package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "grpc-python-go-sum/protos"
)

const (
	address = "localhost:50051"
)

func main() {
	log.Printf("Request 1 + 1")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Sum(ctx, &pb.SumRequest{Operand1: 1, Operand2: 1})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}
	log.Printf("Sum 1 + 1 = %d", r.GetResult())
}
