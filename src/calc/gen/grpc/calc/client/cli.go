// Code generated by goa v3.2.6, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen calc/design

package client

import (
	calc "calc/gen/calc"
	calcpb "calc/gen/grpc/calc/pb"
	"encoding/json"
	"fmt"
)

// BuildAddPayload builds the payload for the calc add endpoint from CLI flags.
func BuildAddPayload(calcAddMessage string) (*calc.AddPayload, error) {
	var err error
	var message calcpb.AddRequest
	{
		if calcAddMessage != "" {
			err = json.Unmarshal([]byte(calcAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"a\": 1698882017578366363,\n      \"b\": 6747375795581831989\n   }'")
			}
		}
	}
	v := &calc.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}

	return v, nil
}

// BuildMinusPayload builds the payload for the calc minus endpoint from CLI
// flags.
func BuildMinusPayload(calcMinusMessage string) (*calc.MinusPayload, error) {
	var err error
	var message calcpb.MinusRequest
	{
		if calcMinusMessage != "" {
			err = json.Unmarshal([]byte(calcMinusMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"a\": 686605435966370186,\n      \"b\": 8228676432890045784\n   }'")
			}
		}
	}
	v := &calc.MinusPayload{
		A: int(message.A),
		B: int(message.B),
	}

	return v, nil
}
