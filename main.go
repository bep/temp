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

//

type readCloserProvider interface {
	OpenReadCloser() (io.ReadCloser, error)
}

type openReadCloser func() (io.ReadCloser, error)

func (o openReadCloser) OpenReadCloser() (io.ReadCloser, error) {
	return o()
}

type contact struct {
	Name string
}

type myhttp struct {
}

func (myhttp) Get(s string) readCloserProvider {
	return openReadCloser(func() (io.ReadCloser, error) {
		resp, err := http.Get(s)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	})
}

// https://jsonplaceholder.typicode.com/todos

func main() {

	globals := map[string]interface{}{
		"http":    myhttp{},
		"Println": fmt.Println,
	}

	script1 := []byte(`
	


url = "https://s1.demo.opensourcecms.com/wordpress/wp-json/wp/v2/posts"


def getTodos():
	return http.Get(url)

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
	//var todos []interface{}
	for {
		var m interface{}
		if err := dec.Decode(&m); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		//todos = append(todos, m)
		fmt.Println(m)
	}

	//fmt.Println("TODOS:", todos)
}

func runFunc(fn *starlark.Function) {
	thread := &starlark.Thread{}
	v, err := starlark.Call(thread, fn, nil, nil)
	if err != nil {
		panic(err)
	}

	f := fromValue(v).(readCloserProvider)
	rrc, err := f.OpenReadCloser()
	if err != nil {
		panic(err)
	}

	decodeTODOs(rrc)
	rrc.Close()

}

func fromValue(v starlark.Value) interface{} {
	vv := convert.FromValue(v)

	return vv

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
