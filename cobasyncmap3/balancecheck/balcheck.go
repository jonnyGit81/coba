package balancecheck

import (
	"fmt"
	"sync"
)

type BalCheck struct {
	m *sync.Map
}

func NewBalCheck(m *sync.Map) *BalCheck {
	return &BalCheck{m: m}
}

func (b *BalCheck) Print() {
	b.m.Range(func(key, value interface{}) bool {
		fmt.Println(value.(string))
		return true
	})
}
