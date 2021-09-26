package main

import (
	"fmt"
	"github.com/jonnyGit81/coba/testsyncmap/structs"
	"sync"
	"time"
)

func main() {
	s := structs.NewData()

	go func() {
		changeValue("A", "A", s.SM)
	}()

	go func() {
		changeValue("A", "B", s.SM)
	}()

	time.Sleep(time.Second * 3)

	s.SM.Range(func(key, value interface{}) bool {
		fmt.Printf("%s", value)
		return true
	})
	time.Sleep(time.Second * 100)

	var a map[string]string
	fmt.Println(a)
}

func changeValue(key string, val string, s sync.Map) {
	s.Store(key, val)
}
