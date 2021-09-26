package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var m sync.Map

type Data struct {
	ID      string
	Balance int
}

func main() {
	for i := 0; i < 1000; i++ {
		m.Store(i, &Data{ID: strconv.Itoa(i), Balance: i})
	}
	go anotherGoroutine("SM1-1")
	go anotherGoroutine2("SM1-2")

	go readGo("Main 1 T1")
	go readGo("Main 1 T2")
	go readGo("Main 1 T3")
	go readGo("Main 1 T4")
	go readGo("Main 1 T5")
	go readGo("Main 1 T6")
	go readGo("Main 1 T7")
	go readGo("Main 1 T8")
	go readGo("Main 1 T9")
	go readGo("Main 1 T10")
	time.Sleep(time.Hour * 10)
}

func anotherGoroutine2(s string) {
	if "a" == "b" {
		fmt.Println("SAMA")
	}
	fmt.Println("TIDAK SAMA MAIN 1", s)
}

func anotherGoroutine(s string) {
	i := 0
	for {
		i++
		if i >= 1000 {
			i = 0
		}
		d := &Data{ID: strconv.Itoa(i), Balance: i}
		m.Delete(i)
		m.Store(i, d)
		fmt.Printf("Update Main 1 %s -- %+v\n", s, d)
	}
}

func readGo(t string) {
	for {
		m.Range(func(key, value interface{}) bool {
			fmt.Printf("READ %s %+v\n", t, value.(*Data))
			return true
		})
	}
}
