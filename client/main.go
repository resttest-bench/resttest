package main

import (
	"log"

	pb "github.com/resttest-bench/resttest/transactions"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTransactionsClient(conn)

	r, err := c.GetTransactions(context.Background(), &pb.GetRequest{UserId: "1", Cursor: "asdf", Count: 10})
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	log.Printf("Response: %s", r)
}
