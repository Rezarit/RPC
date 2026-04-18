package main

import (
	"context"
	"log"

	calc "demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 建立与 gRPC 服务器的连接
	// 服务器地址 (localhost:8888)
	// 使用不安全的凭证（不使用 TLS 加密，适合开发环境）
	conn, _ := grpc.NewClient(
		"localhost:8888",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	// 确保在函数退出时关闭连接
	defer conn.Close()

	// 创建 gRPC 客户端实例
	client := calc.NewCalcClient(conn)

	// 调用远程的 Add 方法
	// context.Background(): 创建一个空的上下文
	resp, _ := client.Add(context.Background(), &calc.AddReq{A: 5, B: 10})

	// 打印服务器返回的结果
	log.Println(resp.Sum)
}
