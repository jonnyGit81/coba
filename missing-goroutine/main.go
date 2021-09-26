package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

func main() {
	var i int
	go func() {
		for {
			i++
			fmt.Println(i)
		}
	}()

	go get_cpu_stats()
	time.Sleep(time.Hour * 3)
}

var cpu_stats = []float64{0}

func get_cpu_stats() {
	for {
		cpu_stats, _ = cpu.Percent(time.Second, true)
	}
}
