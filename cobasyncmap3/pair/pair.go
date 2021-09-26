package pair

import "sync"

func GetPair() sync.Map {
	var m sync.Map
	m.Store("A", 1)
	m.Store("B", 2)
	m.Store("C", 3)
	return m
}
