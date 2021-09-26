package benchmarksyncmap

import (
	"fmt"
	"sync"
	"testing"
)

// go test -run none -bench . -benchtime 3s
// go test -run none -bench . -benchtime 3s -benchmem
var s sync.Map

var p = &sync.Map{}

var gs string

// BenchmarkSprint tests the performance of using Sprint
func BenchmarkSprint(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}
	gs = s
}

// BenchmarkSprint tests the performance of using Sprintf
func BenchmarkSprintf(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}
	gs = s
}

func BenchmarkWriteSyncMapNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Store(i, i)
	}
}

func BenchmarkWriteSyncMapPointer(b *testing.B) {
	p = &sync.Map{}
	for i := 0; i < b.N; i++ {
		p.Store(i, i)
	}
}

func BenchmarkReadSyncMapNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Load(i)
	}
}

func BenchmarkReadSyncMapPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Load(i)
	}
}
