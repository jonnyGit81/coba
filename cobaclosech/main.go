package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan struct{}, 1000)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		i := 0
		for {
			if i == 1000 {
				break
			}

			ch <- struct{}{}
			i++
		}
		wg.Done()
	}()

	time.Sleep(time.Second * 5)
	go func() {
		for {
			fmt.Println("print 1", len(ch))
			if len(ch) == 0 {
				break
			}
			fmt.Println("print 2", len(ch))
			<-ch
		}
		wg.Done()
	}()

	wg.Wait()
}
