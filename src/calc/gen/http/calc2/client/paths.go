// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the calc2 service.
//
// Command:
// $ goa gen calc/design

package client

import (
	"fmt"
)

// MinusCalc2Path returns the URL path to the calc2 service minus HTTP endpoint.
func MinusCalc2Path(a int, b int) string {
	return fmt.Sprintf("/minus/%v/%v", a, b)
}
