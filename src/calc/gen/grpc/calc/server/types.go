// Code generated by goa v3.2.6, DO NOT EDIT.
//
// calc gRPC server types
//
// Command:
// $ goa gen calc/design

package server

import (
	calc "calc/gen/calc"
	calcpb "calc/gen/grpc/calc/pb"

	goa "goa.design/goa/v3/pkg"
)

// NewAddPayload builds the payload of the "add" endpoint of the "calc" service
// from the gRPC request type.
func NewAddPayload(message *calcpb.AddRequest) *calc.AddPayload {
	v := &calc.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewAddResponse builds the gRPC response type from the result of the "add"
// endpoint of the "calc" service.
func NewAddResponse(result int) *calcpb.AddResponse {
	message := &calcpb.AddResponse{}
	message.Field = int32(result)
	return message
}

// NewDividePayload builds the payload of the "divide" endpoint of the "calc"
// service from the gRPC request type.
func NewDividePayload(message *calcpb.DivideRequest) *calc.DividePayload {
	v := &calc.DividePayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewDivideResponse builds the gRPC response type from the result of the
// "divide" endpoint of the "calc" service.
func NewDivideResponse(result int) *calcpb.DivideResponse {
	message := &calcpb.DivideResponse{}
	message.Field = int32(result)
	return message
}

// ValidateAddRequest runs the validations defined on AddRequest.
func ValidateAddRequest(message *calcpb.AddRequest) (err error) {
	if message.A < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.a", message.A, 1, true))
	}
	if message.A > 10 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.a", message.A, 10, false))
	}
	return
}
