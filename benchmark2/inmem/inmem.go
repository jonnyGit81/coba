package inmem

import "sync"

type User struct {
	Balance int
}

type Mem struct {
	AllUsers sync.Map
}
