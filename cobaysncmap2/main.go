package main

import (
	"fmt"
	"log"
	"sync"
)

type User struct {
	i  int
	mu sync.Mutex
}

var m sync.Map

func main() {
	for {
		coba()
	}
}

func coba() {
	m.Store("count", &User{})

	var wg sync.WaitGroup

	for numOfThread := 0; numOfThread < 10; numOfThread++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				value, ok := m.Load("count")
				v, ok := value.(*User)
				if ok {
					v.mu.Lock()
					v.i++
					// m.Store("count", v)
					v.mu.Unlock()
				} else {
					log.Println("NOT FOUND?")
				}
			}
		}()

		// // Read
		go func() {
			for i := 0; i < 50000; i++ {
				value, ok := m.Load("count")
				v, ok := value.(*User)
				if ok {
					fmt.Printf("%+v\n", v)
				} else {
					log.Println("NOT FOUND?")
				}
			}
		}()
	}

	//log.Println("threads starts")
	wg.Wait()

	value, ok := m.Load("count")
	if ok {
		v, _ := value.(*User)
		if v.i != 50 {
			log.Printf("WRONG >>>: %+v", v)
		}
		//log.Printf("final count: %+v", v)
	}

	// log.Println("all done")
}

// func main() {
// 	var m *sync.Map
// 	m = &sync.Map{}
// 	m.Store("count", 0)
//
// 	var wg sync.WaitGroup
//
// 	for numOfThread := 0; numOfThread < 10; numOfThread++ {
// 		wg.Add(1)
// 		go func() {
//
// 			for i := 0; i < 5; i++ {
// 				value, ok := m.LoadOrStore("count", i+1)
// 				if !ok {
// 					log.Println("load count error")
// 				} else {
// 					v, _ := value.(int)
// 					m.Store("count", v+1)
// 				}
// 			}
// 			 wg.Done()
// 		}()
// 	}
//
// 	log.Println("threads starts")
// 	wg.Wait()
//
// 	value, ok := m.Load("count")
// 	if ok {
// 		v, _ := value.(int)
// 		log.Printf("final count: %d", v)
// 	}
//
// 	log.Println("all done")
// }
