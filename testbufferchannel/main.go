package main

import (
	"github.com/jonnyGit81/coba/testbufferchannel/channel"
	"time"
)

func main() {

	chn := make(chan *channel.Request, 100)

	go func() {
		for i := 0; i < 1e6; i++ {
			chn <- &channel.Request{ID: i, Thread: "Th-1"}
		}
	}()
	go func() {
		for i := 0; i < 1e6; i++ {
			chn <- &channel.Request{ID: i, Thread: "Th-2"}
		}
	}()
	// go func() {
	// 	for i := 0; i < 1e6; i++ {
	// 		chn <- &channel.Request{ID: i, Thread:"Th-3"}
	// 	}
	// }()

	c := channel.Channel{CH: chn}
	c.ProcessQue()
	c.ProcessQue2()
	time.Sleep(time.Hour * 1)
}
