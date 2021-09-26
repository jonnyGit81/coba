package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int, 2)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1e6; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1e6; i++ {
			ch <- i
		}
	}()

	for {
		select {
		case i := <-ch:
			doSomeThing(i)
		}
	}

	wg.Wait()
}

func doSomeThing(i int) {
	fmt.Println(i)
	//time.Sleep(time.Hour * 1)
}
