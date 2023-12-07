package main

import (
	"context"
	"fmt"
	"log"
	"prototut/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	conn   *grpc.ClientConn
	client pb.CalculatorClient
}

func (c *client) Close() {
	defer c.conn.Close()
}

func New(serverAddr string) *client {
	i := insecure.NewCredentials()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(i),
	}

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalln("Failed to dial:", err)
	}

	c := pb.NewCalculatorClient(conn)

	return &client{
		conn:   conn,
		client: c,
	}
}

func (c *client) Sum(nums []int64) int64 {
	result, err := c.client.Sum(context.Background(), &pb.NumbersRequest{Numbers: nums})
	if err != nil {
		log.Fatalln("Failed to sum:", err)
	}

	return result.GetResult()
}

func main() {
	client := New("localhost:8080")
	defer client.Close()

	r := client.Sum([]int64{1, 2, 3, 4, 5})

	fmt.Println(r)
}
