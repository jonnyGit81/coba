package main

import (
	"fmt"
	"github.com/jonnyGit81/coba/coba_mutex/user"
)

func main() {
	u := &user.User{}

	u.UpdateBalance(1)
	fmt.Printf("%+v", u)
}
