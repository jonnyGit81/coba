package main

import (
	"context"
	"github.com/jonnyGit81/coba/priceclient/pkg"
	"github.com/jonnyGit81/coba/priceclient/price_api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"os/signal"
)

var dc bool

var logger *zap.SugaredLogger

var symbols = []string{
	"USD/USD",
	"XAG/EUR",
	"XAG/EUR-LP",
	"XAG/EUR-reg",
	"XAG/EUR-vip",
	"XAG/USD",
	"XAG/USD-reg",
	"XAG/USD-vip",
	"XAU/EUR",
	"XAU/EUR-reg",
	"XAU/EUR-vip",
	"XAU/USD",
	"XAU/USD-reg",
	"XAU/USD-vip",
	"XPD/USD",
	"XPD/USD-reg",
	"XPD/USD-vip",
	"XPT/USD",
	"XPT/USD-reg",
	"XPT/USD-vip",
	"UBS:100G99.99/USD",
	"UBS:EUR/USD",
	"UBS:GB 1KG 4X9/USD",
	"UBS:XAG/EUR",
	"UBS:XAG/USD",
	"UBS:XAU/EUR",
	"UBS:XAU/USD",
	"UBS:XPT/EUR",
	"UBS:XPT/USD",
}

var dummyPrices []*price_api.Prices

func initDummyPrices() {
	dummyPrices = make([]*price_api.Prices, 0, len(symbols))

	for _, v := range symbols {
		bids := []*price_api.PriceLevel{
			&price_api.PriceLevel{
				Price:      "1800.59",
				Qty:        "500",
				OrderCount: "0",
			},
			&price_api.PriceLevel{
				Price:      "1800.52",
				Qty:        "1000",
				OrderCount: "0",
			},
			&price_api.PriceLevel{
				Price:      "1800.34",
				Qty:        "2500",
				OrderCount: "0",
			},
			&price_api.PriceLevel{
				Price:      "1800.05",
				Qty:        "3500",
				OrderCount: "0",
			},
		}
		asks := []*price_api.PriceLevel{
			&price_api.PriceLevel{
				Price:      "1801.8",
				Qty:        "3500",
				OrderCount: "0",
			},
		}
		price := &price_api.Prices{
			Symbol: v,
			Time:   "2021-09-15T07:15:00.454Z",
			Bids:   bids,
			Asks:   asks,
			Bid:    "1800.59",
			Ask:    "1801.8",
		}
		dummyPrices = append(dummyPrices, price)
	}
}

func main() {

	logger := applog.InitLogger()

	initDummyPrices()

	con, conErr := grpc.Dial("localhost:5566", grpc.WithInsecure())
	if conErr != nil {
		logger.Fatalf("Failed to connect to matching engine %v", conErr)
	}
	defer con.Close()

	client := price_api.NewAPIClient(con)

	go func() {
		for {
			if !dc {
				for {
					if dc {
						break
					}
					for _, p := range dummyPrices {
						res, err := client.NewPrices(context.Background(), p)
						if err != nil {
							logger.Infof("error calling NewPrices : %v\n", err)
						}
						logger.Infof("response from server: %+v\n", res)
					}
				}
			}
		}
	}()

	// to simulate bind to os and get signal if we hit CTRL+C
	// hit CTRL+Z to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		// every CTRL+C we will receive sigChan
		<-sigChan
		dc = !dc
	}

}
