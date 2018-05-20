package main

import (
	"fmt"
	"log"

	"reflect"

	"github.com/mitchellh/mapstructure"
)

type T struct {
	M1 map[string]interface{}
	S1 []interface{}
	S2 []interface{}
}

func main() {
	dest := &T{}

	hook := func(t1 reflect.Type, t2 reflect.Type, v interface{}) (interface{}, error) {
		fmt.Printf("%v\n", v)
		return v, nil
	}

	input := map[string]interface{}{
		"M1": map[string]interface{}{
			"V1": "v1",
			"M2": map[string]interface{}{
				"V2": "v2",
			},
		},
		"S1": []interface{}{"VS1_1", "VS1_2"},
		"S2": []interface{}{map[string]interface{}{
			"VMS": "vms",
		}},
	}

	conf := &mapstructure.DecoderConfig{
		DecodeHook:       hook,
		ErrorUnused:      false,
		WeaklyTypedInput: false,
		Metadata:         nil,
		Result:           dest,
		ZeroFields:       true,
	}

	dec, err := mapstructure.NewDecoder(conf)
	if err != nil {
		log.Fatal(err)
	}
	err = dec.Decode(input)
	if err != nil {
		log.Fatal(err)
	}

}
