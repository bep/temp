package main

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

type B struct {
	foo string
}

func (b *B) Foo() string {
	return b.foo
}

func (b *B) String() string {
	return b.Foo()
}

func TestNil(t *testing.T) {
	c := qt.New(t)

	var a *B

	c.Assert(a, qt.Not(qt.IsNil))
}
