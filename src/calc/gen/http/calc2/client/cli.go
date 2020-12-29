// Code generated by goa v3.2.6, DO NOT EDIT.
//
// calc2 HTTP client CLI support package
//
// Command:
// $ goa gen calc/design

package client

import (
	calc2 "calc/gen/calc2"
	"fmt"
	"strconv"
)

// BuildMinusPayload builds the payload for the calc2 minus endpoint from CLI
// flags.
func BuildMinusPayload(calc2MinusA string, calc2MinusB string) (*calc2.MinusPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calc2MinusA, 10, 64)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calc2MinusB, 10, 64)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc2.MinusPayload{}
	v.A = a
	v.B = b

	return v, nil
}
