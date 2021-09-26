package server

import (
	"fmt"
	"github.com/jonnyGit81/coba/cobagrpc_clientstreaming/price_api"
	"io"
	"log"
)

type PriceServer struct{}

func (PriceServer) MEasClientOfPEAndStreamingSendPriceUpdateToPE(stream price_api.API_MEasClientOfPEAndStreamingSendPriceUpdateToPEServer) error {
	for {
		price, err := stream.Recv()

		if err == io.EOF {
			log.Println("EOF")
			return stream.SendAndClose(&price_api.Acknowledge{Ack: "DONE"})
		}

		if err != nil {
			log.Println("HAVE ERROR HERE", err)
			return err
		}

		fmt.Printf("PRICE_SERVER_RECEIVED_STREAM_DATA_FROM_ME_TO_PROCEED %+v", price)
	}
}
