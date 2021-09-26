package main

import (
	"github.com/jonnyGit81/coba/poc_price_publisher/proto"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

var allPrices = &sync.Map{}
var max int = 5
var min int = 0

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomSleep() time.Duration {
	return time.Duration(rand.Intn(max-min) + min)
}

func randomSymbol() int {
	return rand.Intn(3)
}

type server struct{}

type subscriber struct {
	subCh chan *match_api.Prices
}

var subscribers = &sync.Map{}

// var sub *subscriber

func main() {
	log.Println("Starting server...")

	// sub = &subscriber{}
	// sub.subCh = make(chan *match_api.Prices, 1000)
	// subscribers.Store(sub, sub) // XXX
	// defer subscribers.Delete(sub)

	simulateCustomerMadeTransaction()

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()
	match_api.RegisterMatchAPIServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func simulateCustomerMadeTransaction() {
	symbolMap := make(map[int]string)
	symbolMap[0] = "USD/CAD"
	symbolMap[1] = "USD/CET"
	symbolMap[2] = "JPY/USD"

	// Simulate that customer do transactions
	go func() {
		for i := 0; i < 1e6; i++ {
			symb := symbolMap[randomSymbol()]
			p := &match_api.Prices{
				Symbol: symbolMap[randomSymbol()],
				Bid:    strconv.Itoa(i),
				Ask:    strconv.Itoa(i),
				Time:   time.Now().String(),
			}
			allPrices.Store(symb, p)

			subscribers.Range(func(key, value interface{}) bool {
				value.(*subscriber).subCh <- p
				return true
			})

			log.Printf("GR-01 customer made transaction %+v", p)
			time.Sleep(randomSleep() * time.Second)
		}
	}()

	// Simulate that customer do transactions
	// go func() {
	// 	for i := 0; i < 1e6; i++ {
	// 		symb := symbolMap[randomSymbol()]
	// 		p := &match_api.Prices{
	// 			Symbol: symbolMap[randomSymbol()],
	// 			Bid:    strconv.Itoa(i),
	// 			Ask:    strconv.Itoa(i),
	// 			Time:   time.Now().String(),
	// 		}
	// 		allPrices.Store(symb, p)
	//
	// 		subscribers.Range(func(key, value interface{}) bool {
	// 			value.(*subscriber).subCh <- p
	// 			return true
	// 		})
	//
	// 		log.Printf("GR-02 customer made transaction %+v", p)
	// 		time.Sleep(randomSleep() * time.Second)
	// 	}
	// }()
}

func (s *server) RequestPricesStreaming(req *match_api.MatchEmpty, stream match_api.MatchAPI_RequestPricesStreamingServer) error {
	sub := &subscriber{}
	sub.subCh = make(chan *match_api.Prices, 1000)
	subscribers.Store(sub, sub) // XXX
	defer close(sub.subCh)
	defer subscribers.Delete(sub)

	// for resume on the first subscribe.
	// subsequently any transaction happen need to push to the channel.
	allPrices.Range(func(key interface{}, val interface{}) bool {
		streamPrices := val.(*match_api.Prices)
		log.Printf("streamPrices = %+v\n", streamPrices)
		sub.subCh <- streamPrices
		return true
	})

	log.Println("1 Client Subscribed : size of channel:", len(sub.subCh))

	for price := range sub.subCh {
		if price == nil {
			log.Println("streamPrices is nil")
			continue
		}
		resp := &match_api.PriceStreamResponse{
			Prices: price,
		}
		if err := stream.Send(resp); err != nil {
			log.Println("Send to subscriber err:", err)
			break
		}
		log.Printf("Response stream prices %+v\n", price)
	}

	log.Println("Client Unsubscribed : size of channel:", len(sub.subCh))
	return nil
}
