// Code generated by goa v3.2.6, DO NOT EDIT.
//
// calc2 gRPC client types
//
// Command:
// $ goa gen calc/design

package client

import (
	calc2 "calc/gen/calc2"
	calc2pb "calc/gen/grpc/calc2/pb"
)

// NewMinusRequest builds the gRPC request type from the payload of the "minus"
// endpoint of the "calc2" service.
func NewMinusRequest(payload *calc2.MinusPayload) *calc2pb.MinusRequest {
	message := &calc2pb.MinusRequest{
		A: int32(payload.A),
		B: int32(payload.B),
	}
	return message
}

// NewMinusResult builds the result type of the "minus" endpoint of the "calc2"
// service from the gRPC response type.
func NewMinusResult(message *calc2pb.MinusResponse) int {
	result := int(message.Field)
	return result
}
