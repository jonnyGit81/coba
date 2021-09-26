package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jonnyGit81/coba/cobagrpc_clientstreaming/price_api"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

var (
	serverAddr = flag.String("server_addr", "localhost:9111", "The server address in the format of host:port")
)

type MePricePublisher struct {
	ch chan *price_api.Prices
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := price_api.NewAPIClient(conn)

	streamPriceUpdateToPE(client)
}

func initDummyData() []*price_api.Prices {
	p := make([]*price_api.Prices, 100)
	for i := 0; i < 100; i++ {
		p = append(p, &price_api.Prices{
			Symbol: "USD/CET" + strconv.Itoa(i),
			Time:   time.Now().String(),
			Bid:    strconv.Itoa(i),
			Ask:    strconv.Itoa(i),
		})
	}
	return p
}

func streamPriceUpdateToPE(c price_api.APIClient) {
	pub := &MePricePublisher{}
	pub.ch = make(chan *price_api.Prices, 10)
	defer close(pub.ch)

	dummyPrices := initDummyData()
	go func(prices []*price_api.Prices) {
		for _, price := range prices {
			pub.ch <- price
		}
	}(dummyPrices)

	stream, err := c.MEasClientOfPEAndStreamingSendPriceUpdateToPE(context.Background())
	if err != nil {
		log.Println("Error on Sending Price: %v", err)
	}

	for price := range pub.ch {
		fmt.Printf("ME streaming price Update request: %+v \n", price)
		fmt.Println()
		stream.Send(price)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("Error when closing the stream and receiving the response: %v", err)
	}
	fmt.Printf("Ack %+v\n", res)
}
