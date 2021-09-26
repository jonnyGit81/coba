package main

import (
	"os"
	"os/signal"
	"time"
)

var disconnected = false

type MessageOut struct {
	Msg string
}

type Subsribers struct {
	ID  int
	Que chan *MessageOut
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		// every CTRL+C we will receive sigChan
		<-sigChan
		disconnected = !disconnected
	}
}

func clearPrice() {
	for {
		send()
		time.Sleep(time.Second * 3)
	}
}

func subscriber() {

	for {

	}

}
