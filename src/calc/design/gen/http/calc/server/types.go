// Code generated by goa v3.2.6, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen calc/design

package server

import (
	calc "calc/design/gen/calc"
)

// NewAddPayload builds a calc service add endpoint payload.
func NewAddPayload(a int, b int) *calc.AddPayload {
	v := &calc.AddPayload{}
	v.A = a
	v.B = b

	return v
}
