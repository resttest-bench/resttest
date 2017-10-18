package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	proto "github.com/golang/protobuf/ptypes"
	pb "github.com/resttest-bench/resttest/transactions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func (s *server) GetTransactions(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {

	var transactions = make([]*pb.Transaction, in.Count)
	var i uint32

	// Generate transactions
	for ; i < in.Count; i++ {
		transactions[i] = &pb.Transaction{Date: proto.TimestampNow(), Ledger: "Some Ledger", Amount: float32(123.05), Company: "A Company"}
	}

	return &pb.GetReply{Transactions: transactions, NextCursor: in.Cursor + "123"}, nil
}

type server struct{}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransactionsServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
