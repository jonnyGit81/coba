package main

import (
	"context"
	"fmt"
	"github.com/jonnyGit81/coba/poc_price_publisher/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Client Starting Subscribe...")
	fmt.Println()

	opts := grpc.WithInsecure()
	con, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("Error connecting: %v \n", err)
	}

	defer con.Close()
	c := match_api.NewMatchAPIClient(con)

	subscribePrices(c)

}

func subscribePrices(c match_api.MatchAPIClient) error {
	req := &match_api.MatchEmpty{}

	// Get the stream and err
	stream, err := c.RequestPricesStreaming(context.Background(), req)
	if err != nil {
		return err
	}

	for {
		// Start receiving streaming messages
		resp, err := stream.Recv()

		// any error break for loop
		if err != nil {
			return err
		}
		// if err == io.EOF {
		// 	log.Printf("EOF: %v", err)
		// 	break
		// }

		log.Printf("Response from RequestPricesStreaming: %v", resp.GetPrices())
	}

}
