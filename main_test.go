package main

import (
	"sync"
	"testing"
)

var global int
var o sync.Once

func something() {
	o.Do(func() {
		global++
	})
	x := global
	for i := 0; i < 1e5; i++ {
		x += global
	}
}

func TestRace(t *testing.T) {
	n := 3
	var wg sync.WaitGroup
	wg.Add(n)
	for j := 0; j < n; j++ {
		go func() {
			something()
			wg.Done()
		}()
	}
	wg.Wait()
}
