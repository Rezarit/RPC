package main

import (
	"context"
	"log"
	"net"

	calc "demo/proto"

	"google.golang.org/grpc"
)

// CalcServer 实现了 gRPC 服务接口
// 嵌入 calc.UnimplementedCalcServer 以确保向前兼容性
type CalcServer struct {
	calc.UnimplementedCalcServer
}

func (s *CalcServer) Add(ctx context.Context, req *calc.AddReq) (*calc.AddResp, error) {
	sum := req.A + req.B
	return &calc.AddResp{Sum: sum}, nil
}

func main() {
	// 创建 TCP 监听器，监听 8888 端口
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	// 创建 gRPC 服务器实例
	server := grpc.NewServer()

	// 将我们的服务实现注册到 gRPC 服务器
	calc.RegisterCalcServer(server, &CalcServer{})

	// 启动 gRPC 服务，开始处理请求
	// Serve 会阻塞直到服务器停止或出错
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
