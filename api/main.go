package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc"
	pb "grpc-python-go-sum/protos"
)

const (
	address = "computation:50051"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s", err)
		return
	}
	operand1Param := r.Form.Get("operand1")
	operand1, err := strconv.ParseInt(operand1Param, 10, 32)
	if err != nil {
		log.Printf("Error parsing operand1: %s", err)
		return
	}

	operand2Param := r.Form.Get("operand2")
	operand2, err := strconv.ParseInt(operand2Param, 10, 32)
	if err != nil {
		log.Printf("Error parsing operand2: %s", err)
		return
	}

	log.Printf("Request sum %d + %d", operand1, operand2)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Sum(ctx, &pb.SumRequest{Operand1: int32(operand1), Operand2: int32(operand2)})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}
	log.Printf("Result %d + %d = %d", operand1, operand2, res.GetResult())

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
