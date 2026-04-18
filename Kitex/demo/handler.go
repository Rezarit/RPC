package main

import (
	"context"
	calc "demo/kitex_gen/demo/calc"
)

// CalcServiceImpl implements the last service interface defined in the IDL.
type CalcServiceImpl struct{}

// Add 实现加法
func (s *CalcServiceImpl) Add(ctx context.Context, req *calc.AddReq) (resp *calc.AddResp, err error) {
	sum := req.A + req.B
	return &calc.AddResp{Sum: sum}, nil
}

// Sub 实现减法
func (s *CalcServiceImpl) Sub(ctx context.Context, req *calc.SubReq) (resp *calc.SubResp, err error) {
	difference := req.A - req.B
	return &calc.SubResp{Difference: difference}, nil
}

// Mul 实现乘法
func (s *CalcServiceImpl) Mul(ctx context.Context, req *calc.MulReq) (resp *calc.MulResp, err error) {
	product := req.A * req.B
	return &calc.MulResp{Product: product}, nil
}

// Div 实现除法，返回商和余数
func (s *CalcServiceImpl) Div(ctx context.Context, req *calc.DivReq) (resp *calc.DivResp, err error) {
	if req.B == 0 {
		return &calc.DivResp{Quotient: 0, Remainder: 0}, nil
	}
	quotient := req.A / req.B
	remainder := req.A % req.B
	return &calc.DivResp{Quotient: quotient, Remainder: remainder}, nil
}
