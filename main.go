package main

import (
	"fmt"
	"sync"
)

type dependencies struct {
	inits []initer
}

type initer interface {
	init() error
}

type A struct {
	deps    *dependencies
	once    sync.Once
	onceErr error
	name    string
	// fn
}

func (a *A) init() error {
	a.once.Do(func() {
		// Init the dependencies first
		for _, dep := range a.deps.inits {
			if dep == a {
				break
			}

			if err := dep.init(); err != nil {
				return
			}
		}

		fmt.Println("init", a.name)
	})
	return nil
}

func main() {

	deps := &dependencies{}

	a := &A{name: "a", deps: deps}
	b := &A{name: "b", deps: deps}
	c := &A{name: "c", deps: deps}

	deps.inits = []initer{b, c, a}

	c.init()

}
