package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	pb "github.com/resttest-bench/resttest/transactions"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

// Just need this because the sort library doesn't have a built-in function for int64
type sortableInt64List []int64

func (a sortableInt64List) Len() int           { return len(a) }
func (a sortableInt64List) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortableInt64List) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTransactionsClient(conn)

	cursor := "-1"
	var totalAmount float32
	dailyAmounts := make(map[int64]float32)

	fmt.Printf("Raw data:\n")

	// Loop while we have data. The cursor will eventually be empty.
	for cursor != "" {
		r, err := c.GetTransactions(context.Background(), &pb.GetRequest{UserId: "1", Cursor: cursor, Count: 10})
		if err != nil {
			log.Fatalf("Request failed: %v", err)
		}
		cursor = r.NextCursor // Advance the cursor

		// For each transaction, push it into a map keyed by day
		for _, transaction := range r.Transactions {
			totalAmount += transaction.Amount

			// Massage the dates
			date := time.Unix(transaction.Date.GetSeconds(), int64(transaction.Date.GetNanos()))
			day := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
			dailyAmounts[day.Unix()] += transaction.Amount

			fmt.Printf("%v\n", transaction)
		}
	}

	// Sort the keys of the map
	var keys sortableInt64List
	for k := range dailyAmounts {
		keys = append(keys, k)
	}
	sort.Sort(keys)

	// Print the results
	fmt.Printf("\n\nDaily Amounts:\n")
	var runningTotal float32
	for _, k := range keys {
		runningTotal += dailyAmounts[k]
		date := time.Unix(k, 0).Format("Mon Jan 2")
		fmt.Printf("%s - %s (Running total: %s)\n", date, strconv.FormatFloat(float64(dailyAmounts[k]), 'f', 2, 32), strconv.FormatFloat(float64(runningTotal), 'f', 2, 32))
	}
	fmt.Println("Total: " + strconv.FormatFloat(float64(totalAmount), 'f', 2, 32))
}
