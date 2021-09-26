package main

import (
	"fmt"
	"github.com/jonnyGit81/coba/price_error_sim/subscribers"
	"sync"
	"time"
)

var subs sync.Map

func sendMessage(msg string, sub *subscribers.Subscriber) {
	sub.Queue <- msg
}

func main() {
	sub := &subscribers.Subscriber{Queue: make(chan string, 1000)}

	go func() {
		for {
			sendMessage("NEW PRICE", sub)
		}
	}()

	// clear price
	go func() {
		for {
			sendMessage("CLEAR", sub)
			time.Sleep(time.Second * 20)
		}
	}()

	// consumer
	for m := range sub.Queue {
		if m == "CLEAR" {
			fmt.Printf(m+" %v\n", time.Now().String())
		}
	}

	time.Sleep(time.Hour * 1)
}
