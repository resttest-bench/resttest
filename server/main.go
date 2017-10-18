package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"

	"github.com/Pallinder/go-randomdata"
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
	var today = time.Now()
	// Generate transactions
	for ; i < in.Count; i++ {
		// Make some random data
		var randDate, _ = proto.TimestampProto(today.Add(-time.Duration(rand.Intn(10000)) * time.Minute))
		var randAmount = rand.Float32() + float32(rand.Intn(1000))
		var randLeger = randomdata.SillyName() + " " + randomdata.SillyName()
		var randCompany = randomdata.SillyName() + " Corp"
		transactions[i] = &pb.Transaction{Date: randDate, Ledger: randLeger, Amount: randAmount, Company: randCompany}
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
