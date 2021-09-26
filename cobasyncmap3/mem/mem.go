package mem

import "sync"

type Mem struct {
	Pair sync.Map
}
