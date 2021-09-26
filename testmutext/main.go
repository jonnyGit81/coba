package main

import (
	"fmt"
	"github.com/jonnyGit81/coba/testmutext/user"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		u := &user.User{}

		wg.Add(3)

		go func() {
			defer wg.Done()
			for i := 1; i <= 100; i++ {
				u.UpdateBal(i) // 1 + 2
			}

		}()

		go func() {
			defer wg.Done()
			for i := 1; i <= 100; i++ {
				u.UpdateBal(i) // 1 + 2
			}

		}()

		go func() {
			defer wg.Done()
			for i := 1; i <= 100; i++ {
				u.UpdateBal(i) // 1 + 2
			}
		}()

		wg.Wait()

		fmt.Println(u.Bal)
	}

	//time.Sleep(time.Hour * 1)

}
