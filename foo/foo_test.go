package foo

import (
	"testing"

	"github.com/fortytw2/leaktest"
)

func TestGoRoutines(t *testing.T) {
	defer leaktest.Check(t)()
	Run()
}
