// Code generated by goa v3.2.6, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen calc/design

package client

import (
	calc "calc/gen/calc"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildAddPayload builds the payload for the calc add endpoint from CLI flags.
func BuildAddPayload(calcAddA string, calcAddB string) (*calc.AddPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcAddA, 10, 64)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
		if a < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("a", a, 1, true))
		}
		if a > 10 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("a", a, 10, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcAddB, 10, 64)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc.AddPayload{}
	v.A = a
	v.B = b

	return v, nil
}

// BuildDividePayload builds the payload for the calc divide endpoint from CLI
// flags.
func BuildDividePayload(calcDivideA string, calcDivideB string) (*calc.DividePayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcDivideA, 10, 64)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcDivideB, 10, 64)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc.DividePayload{}
	v.A = a
	v.B = b

	return v, nil
}
