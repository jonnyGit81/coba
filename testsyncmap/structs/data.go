package structs

import "sync"

type Data struct {
	SM sync.Map
}

func NewData() *Data {
	return &Data{}
}
