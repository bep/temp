package foo

import (
	"fmt"
	"testing"
)

func BenchmarkStructState(b *testing.B) {
	as := createAs(500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := handleWithStructState(as)
		if res == nil {
			b.Fatal(res)
		}
	}
}

func BenchmarkMapState(b *testing.B) {
	as := createAs(500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := handleWithMapState(as)
		if res == nil {
			b.Fatal(res)
		}
	}
}

func createAs(num int) []*A {
	as := make([]*A, num)
	for i := 0; i < num; i++ {
		as[i] = &A{name: fmt.Sprintf("name%d", i)}
	}

	return as
}
