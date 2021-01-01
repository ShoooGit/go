package calcapi

import (
	calc "calc/gen/calc"
	"context"
	"fmt"
	"log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	return p.A + p.B, nil
}

// Divide implements divide.
func (s *calcsrvc) Divide(ctx context.Context, p *calc.DividePayload) (res int, err error) {
	if p.B == 0 {
		// 生成された関数を利用してエラー結果を作成する
		return 0, calc.MakeDivByZero(fmt.Errorf("right operand cannot be 0"))
	}
	return p.A / p.B, nil
}
