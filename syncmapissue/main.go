package main

import (
	"fmt"
	"sync"
	"time"
)

type Prices struct {
	Symbol  string
	Balance int
}

var allPrices sync.Map

func main() {
	p := &Prices{Symbol: "USD/AMX", Balance: 10000}

	allPrices.Store("USD/AMX", p)

	go func() {
		for {
			d, ok := allPrices.Load("USD/AMX")
			if ok {
				fmt.Printf("balance %d\n", d.(*Prices).Balance)
			} else {
				fmt.Printf("d is nil %+v", d)
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)
		allPrices.Range(func(key, value interface{}) bool {
			prices := value.(*Prices)
			newPrices := &Prices{Symbol: prices.Symbol}
			allPrices.Delete(key)
			allPrices.Store(key, newPrices)
			return true
		})
	}()

	time.Sleep(time.Hour * 1)
}
