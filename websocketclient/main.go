package main

import (
	"flag"
	"fmt"
	"github.com/jonnyGit81/coba/websocketclient/wsclient"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var addr = flag.String("addr", "price-metalor-uat.hydrax.io", "http service address")

type MsgInOld struct {
	Type    string   `json:"type"`
	Symbols []string `json:"symbols"`
}

func main() {

	clients := 1000

	for i := 0; i < clients; i++ {
		go connectToWebsocket()
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	fmt.Println("Goodbye")

	// client, err := wsclient.NewWebSocketClient(*addr, "stream_prices")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Connecting")
	//
	// time.Sleep(time.Second * 5)
	//
	// go func() {
	//
	// 	symbols := []string{"DUMMY", "INTL:XAU/USD", "JPM:HKD/SGD", "JPM:USD/HKD", "JPM:USD/SGD", "JPM:XAG/SGD", "JPM:XAG/USD", "JPM:XAU/SGD", "JPM:XAU/USD", "JPM:XPD/USD", "JPM:XPT/USD", "UBS:AUD/USD", "UBS:HKD/SGD", "UBS:USD/HKD", "UBS:USD/SGD", "UBS:XAG/SGD", "UBS:XAG/USD", "UBS:XAU/SGD", "UBS:XAU/USD", "UBS:XPD/SGD", "UBS:XPD/USD", "UBS:XPT/SGD", "UBS:XPT/USD", "XAU/USD-reg", "XAU/USD-vip", "XAG/USD-reg", "XAG/USD-vip", "XAG/SGD-reg", "XAG/SGD-vip", "XPD/USD-reg", "XPD/USD-vip", "XPT/SGD-reg", "XPT/SGD-vip", "XPT/USD-reg", "XPT/USD-vip", "XPD/SGD-vip", "XPD/SGD-reg", "XAU/SGD-reg", "XAU/SGD-vip"}
	// 	reqBody := &MsgInOld{Type: "set_symbols", Symbols: symbols}
	// 	err := client.Write(reqBody)
	// 	if err != nil {
	// 		fmt.Printf("error: %v, writing error\n", err)
	// 	}
	//
	// 	// write down data every 100 ms
	// 	// ticker := time.NewTicker(time.Millisecond * 1500)
	// 	// i := 0
	// 	// for range ticker.C {
	// 	// 	err := client.Write("ping")
	// 	// 	if err != nil {
	// 	// 		fmt.Printf("error: %v, writing error\n", err)
	// 	// 	}
	// 	// 	i++
	// 	// }
	// }()
	//
	// // Close connection correctly on exit
	// sigs := make(chan os.Signal, 1)
	//
	// // `signal.Notify` registers the given channel to
	// // receive notifications of the specified signals.
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//
	// // The program will wait here until it gets the
	// <-sigs
	// client.Stop()
	// fmt.Println("Goodbye")
}

func connectToWebsocket() {
	forever := make(chan struct{})
	{
		client, err := wsclient.NewWebSocketClient(*addr, "stream_prices")
		defer client.Stop()

		if err != nil {
			panic(err)
		}
		fmt.Println("Connecting")

		time.Sleep(time.Second * 5)

		go func() {

			symbols := []string{"DUMMY", "INTL:XAU/USD", "JPM:HKD/SGD", "JPM:USD/HKD", "JPM:USD/SGD", "JPM:XAG/SGD", "JPM:XAG/USD", "JPM:XAU/SGD", "JPM:XAU/USD", "JPM:XPD/USD", "JPM:XPT/USD", "UBS:AUD/USD", "UBS:HKD/SGD", "UBS:USD/HKD", "UBS:USD/SGD", "UBS:XAG/SGD", "UBS:XAG/USD", "UBS:XAU/SGD", "UBS:XAU/USD", "UBS:XPD/SGD", "UBS:XPD/USD", "UBS:XPT/SGD", "UBS:XPT/USD", "XAU/USD-reg", "XAU/USD-vip", "XAG/USD-reg", "XAG/USD-vip", "XAG/SGD-reg", "XAG/SGD-vip", "XPD/USD-reg", "XPD/USD-vip", "XPT/SGD-reg", "XPT/SGD-vip", "XPT/USD-reg", "XPT/USD-vip", "XPD/SGD-vip", "XPD/SGD-reg", "XAU/SGD-reg", "XAU/SGD-vip"}
			reqBody := &MsgInOld{Type: "set_symbols", Symbols: symbols}
			err := client.Write(reqBody)
			if err != nil {
				fmt.Printf("error: %v, writing error\n", err)
			}

			// write down data every 100 ms
			// ticker := time.NewTicker(time.Millisecond * 1500)
			// i := 0
			// for range ticker.C {
			// 	err := client.Write("ping")
			// 	if err != nil {
			// 		fmt.Printf("error: %v, writing error\n", err)
			// 	}
			// 	i++
			// }
		}()
	}
	forever <- struct{}{}
}
