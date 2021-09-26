package main

import (
	"fmt"
	"github.com/jonnyGit81/coba/cobasyncmap3/balancecheck"
	"strconv"
	"sync"
	"time"
)

var m sync.Map

func main() {
	fmt.Print(m.Load("ad"))
	b := balancecheck.NewBalCheck(&m)
	go func() {
		for i := 0; i < 100; i++ {
			m.Store(i, strconv.Itoa(i))
		}
		fmt.Println("finish")
	}()

	time.Sleep(time.Second * 2)
	b.Print()
	// me := mem.Mem{}
	//  me.Pair = pair.GetPair()
	//  me.Pair.Range(func(key, value interface{}) bool {
	// 	 fmt.Println(value)
	//  	return true
	//  })
}
