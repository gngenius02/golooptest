package main

import (
	"testing"
)

func BenchmarkGetHashUsingLoop10Million(b *testing.B) {
	b.ReportAllocs()
	h := &Hasher{bs: []byte("abc"), b: make([]byte, 64)}
	for i := 0; i < b.N; i++ {
		h.HashFn()
	}
}
