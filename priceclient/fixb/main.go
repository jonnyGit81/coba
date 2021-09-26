package main

import (
	"context"
	"fmt"
	"github.com/jonnyGit81/coba/priceclient/pkg"
	"github.com/jonnyGit81/coba/priceclient/price_api"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"time"
)

var dc bool

var logger *zap.SugaredLogger

var symbols = []string{
	"STONEX:100G99.99/USD",
	"STONEX:CHF/USD",
	"STONEX:EUR/USD",
	"STONEX:GB 1KG 4X9/USD",
	"STONEX:USD/CHF",
	"STONEX:USD/HKD",
	"STONEX:XAG/EUR",
	"STONEX:XAG/GBP",
	"STONEX:XAG/USD",
	"STONEX:XAU/EUR",
	"STONEX:XAU/GBP",
	"STONEX:XAU/USD",
	"STONEX:XPD/USD",
	"STONEX:XPT/EUR",
	"STONEX:XPT/USD",
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

	INTERVAL := viper.GetInt("price_clearing_interval")

	fmt.Println(INTERVAL)
	time.Sleep(time.Hour * 1)
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
