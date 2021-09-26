package channel

import "fmt"

type Request struct {
	ID     int
	Thread string
}

type Channel struct {
	CH chan *Request
}

func (c *Channel) ProcessQue() {
	go func() {
		for chn := range c.CH {
			fmt.Println(len(c.CH))
			fmt.Printf("Q1 - %+v\n", chn)
		}
	}()
}

func (c *Channel) ProcessQue2() {
	go func() {
		for chn := range c.CH {
			fmt.Println(len(c.CH))
			fmt.Printf("Q2 - %+v\n", chn)
		}
	}()
}
