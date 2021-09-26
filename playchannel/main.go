package main

import (
	"fmt"
	"sync"
)

//var ch chan struct{}

var gm sync.Map

type Price struct {
	Symbol  string
	Balance int
}

var ch = make(chan struct{})

func main() {
	// p:= &Price{
	// 	Symbol:"ABC/BCA",
	// 	Balance: 10000,
	// }
	// gm.Store(p.Symbol, p)
	//
	// go func() {
	// 	for  {
	// 		v, ok := gm.Load("ABC/BCA")
	// 		if ok {
	// 			if v.(*Price).Balance == 0 {
	// 				break
	// 			}
	// 			v.(*Price).Balance++
	// 		}
	// 		fmt.Println(v.(*Price).Balance)
	// 	}
	// }()
	//ch <- struct{}{}
	ch <- struct{}{}
	go clear()
}

func clear() {
	for range ch {
		// time.Sleep(time.Second * 3)
		gm.Range(func(key, value interface{}) bool {
			fmt.Println("DELETED")
			p := &Price{
				Symbol: value.(*Price).Symbol,
			}
			gm.Delete(p.Symbol)
			gm.Store(p.Symbol, p)
			fmt.Println("Deleted")
			return true
		})
	}
	ch <- struct{}{}
	go clear()
}
