package main

import (
	"fmt"
	"github.com/jonnyGit81/coba/testinterface/intf"
)

func main() {
	c := intf.NewImpl()
	calcAndPrint(c)
}

func calcAndPrint(i intf.IInterface) {
	for j := 0; j < 1e6; j++ {
		i.AddV1(j)
		i.AddV2(j - 1)
		i.Calc()
		fmt.Println(i.GetV3())
	}
}
