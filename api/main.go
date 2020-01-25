package main

import (
	"context"
	"encoding/json"
	"io"
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

func handlerSum(w http.ResponseWriter, r *http.Request) {
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

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	stream, err := c.Sum(ctx, &pb.SumRequest{Operand1: int32(operand1), Operand2: int32(operand2)})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.Sum(_) = _, %v", c, err)
		}

		if res.GetFinished() {
			log.Printf("Result %d + %d = %d", operand1, operand2, res.GetResult())

			js, err := json.Marshal(res)
			if err != nil {
				log.Fatalf("%v.Sum(_) = _, %v", c, err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		} else {
			log.Printf("Awaiting for computation result")
		}
	}

}

func main() {
	log.Print("Starting server")
	http.HandleFunc("/sum", handlerSum)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
