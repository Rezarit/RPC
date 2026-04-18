package main

import (
	"context"
	"demo/kitex_gen/demo/calc"
	"demo/kitex_gen/demo/calc/calcservice"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
)

func main() {
	// 创建客户端连接
	client := calcservice.MustNewClient("calc",
		client.WithHostPorts("127.0.0.1:8888"),
	)

	// 测试加法
	fmt.Println("=== 测试加法 ===")
	addReq := &calc.AddReq{A: 10, B: 5}
	addResp, err := client.Add(context.Background(), addReq)
	if err != nil {
		log.Fatalf("Add failed: %v", err)
	}
	fmt.Printf("%d + %d = %d\n\n", addReq.A, addReq.B, addResp.Sum)

	// 测试减法
	fmt.Println("=== 测试减法 ===")
	subReq := &calc.SubReq{A: 10, B: 5}
	subResp, err := client.Sub(context.Background(), subReq)
	if err != nil {
		log.Fatalf("Sub failed: %v", err)
	}
	fmt.Printf("%d - %d = %d\n\n", subReq.A, subReq.B, subResp.Difference)

	// 测试乘法
	fmt.Println("=== 测试乘法 ===")
	mulReq := &calc.MulReq{A: 10, B: 5}
	mulResp, err := client.Mul(context.Background(), mulReq)
	if err != nil {
		log.Fatalf("Mul failed: %v", err)
	}
	fmt.Printf("%d × %d = %d\n\n", mulReq.A, mulReq.B, mulResp.Product)

	// 测试除法
	fmt.Println("=== 测试除法 ===")
	divReq := &calc.DivReq{A: 10, B: 3}
	divResp, err := client.Div(context.Background(), divReq)
	if err != nil {
		log.Fatalf("Div failed: %v", err)
	}
	fmt.Printf("%d ÷ %d = %d 余 %d\n\n", divReq.A, divReq.B, divResp.Quotient, divResp.Remainder)

	fmt.Println("所有测试完成！🎉")
}
