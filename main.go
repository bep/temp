package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/starlight-go/starlight"
	"github.com/starlight-go/starlight/convert"
	"go.starlark.net/starlark"
)

type contact struct {
	Name string
}

// https://jsonplaceholder.typicode.com/todos

func main() {
	get := func(url string) io.Reader {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		return resp.Body
	}

	globals := map[string]interface{}{
		"Get":     get,
		"Println": fmt.Println,
	}

	script1 := []byte(`
	
def idiv(x, y):
  return x // y


def getTodos():
	return Get("https://jsonplaceholder.typicode.com/todos")



`)

	out, err := starlight.Eval(script1, globals, nil)
	if err != nil {
		panic(err)
	}

	f1 := out["getTodos"].(*starlark.Function)

	runFunc(f1)

}

func decodeTODOs(r io.Reader) {
	dec := json.NewDecoder(r)
	var todos []interface{}
	for {
		var m interface{}
		if err := dec.Decode(&m); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		todos = append(todos, m)
	}

	fmt.Println("TODOS:", todos)
}

func runFunc(fn *starlark.Function) {
	thread := &starlark.Thread{}
	v, err := starlark.Call(thread, fn, nil, nil)
	if err != nil {
		panic(err)
	}

	rrc := convert.FromValue(v).(io.ReadCloser)

	decodeTODOs(rrc)
	rrc.Close()

}

func toValue(args ...interface{}) starlark.Tuple {
	var tuple starlark.Tuple
	for _, v := range args {
		vv, err := convert.ToValue(v)
		if err != nil {
			panic(err)
		}
		tuple = append(tuple, vv)
	}

	return tuple
}
