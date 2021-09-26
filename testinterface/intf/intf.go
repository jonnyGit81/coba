package intf

import "fmt"

type IInterface interface {
	Calc()
	AddV1(v int)
	AddV2(v int)
	GetV3() int
}

type Impl struct {
	Val1 int
	Val2 int
	Val3 int
}

func NewImpl() *Impl {
	return &Impl{}
}

func (i *Impl) Calc() {
	i.plus()
}

func (i *Impl) AddV1(v int) {
	i.Val1 += v
}
func (i *Impl) AddV2(v int) {
	i.Val2 += v
}

func (i *Impl) GetV3() int {
	return i.Val3
}

func (i *Impl) plus() {
	fmt.Println("calling non interface method")
	i.Val3 = i.Val1 + i.Val2
}
