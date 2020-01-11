package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	pb "grpc-python-go-sum/protos"
)

const (
	address = "computation:50051"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request 1 + 1")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Sum(ctx, &pb.SumRequest{Operand1: 1, Operand2: 1})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}
	log.Printf("Result 1 + 1 = %d", res.GetResult())

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	log.Print("Starting server")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
