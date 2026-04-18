# RPC 教程 - gRPC & Kitex 实战

本仓库包含 gRPC 和 Kitex 两个主流 Go RPC 框架的完整教程和实战代码。

---

## 📚 目录

- [项目简介](#项目简介)
- [教程文档](#教程文档)
- [项目结构](#项目结构)
- [快速开始](#快速开始)
- [gRPC 教程](#grpc-教程)
- [Kitex 教程](#kitex-教程)
- [RPC 核心概念](#rpc-核心概念)
- [常见问题](#常见问题)
- [许可证](#许可证)

---

## 项目简介

本仓库是一个完整的 RPC 教程项目，包含：

- ✅ **gRPC 实战** - Google 开源的跨语言 RPC 框架
- ✅ **Kitex 实战** - 字节跳动开源的高性能 Go RPC 框架
- ✅ **核心概念讲解** - RPC vs HTTP、序列化、服务治理等
- ✅ **完整代码示例** - 加减乘除计算器服务
- ✅ **详细教程文档** - 适合预习和复习

---

## 教程文档

| 文档 | 说明 |
|------|------|
| [RPC.md](./RPC.md) | RPC 核心概念与微服务实战|
| [gRPC/gRPC教程详解.md](./gRPC/gRPC教程详解.md) | gRPC 完整教程|
| [Kitex/Kitex教程详解.md](./Kitex/Kitex教程详解.md) | Kitex 完整教程|

---

## 项目结构

```
rpc/
├── README.md                          # 本文件
├── RPC.md                             # RPC 核心概念教程
├── .gitignore                         # Git 忽略文件
├── gRPC/                              # gRPC 教程目录
│   ├── gRPC教程详解.md               # gRPC 详细教程
│   └── demo/                          # gRPC 示例代码
│       ├── proto/                     # proto 文件
│       │   ├── calc.proto            # 接口定义
│       │   ├── calc.pb.go            # 生成的数据结构
│       │   └── calc_grpc.pb.go       # 生成的服务端/客户端代码
│       ├── server/                    # 服务端代码
│       │   └── server.go
│       ├── client/                    # 客户端代码
│       │   └── client.go
│       ├── go.mod
│       └── go.sum
└── Kitex/                             # Kitex 教程目录
    ├── Kitex教程详解.md              # Kitex 详细教程
    └── demo/                          # Kitex 示例代码
        ├── proto/                     # thrift 文件
        │   └── calc.thrift           # 接口定义
        ├── handler.go                  # 业务逻辑实现
        ├── main.go                     # 服务启动文件
        ├── client/                     # 客户端代码
        │   └── client.go
        ├── go.mod
        └── go.sum
```

---

## 快速开始

### 前置条件

- Go 1.19+
- Git
- protoc 编译器（gRPC 需要）
- kitex 工具（Kitex 需要）

---

### gRPC 快速开始

```bash
# 进入 gRPC 示例目录
cd gRPC/demo

# 启动服务端（终端 1）
go run ./server

# 运行客户端（终端 2）
go run ./client
```

---

### Kitex 快速开始

```bash
# 进入 Kitex 示例目录
cd Kitex/demo

# 启动服务端（终端 1）
go run .

# 运行客户端（终端 2）
go run ./client
```

---

## gRPC 教程

### 功能特性

- ✅ 跨语言支持（Java/Python/Go/...）
- ✅ 基于 HTTP/2，高性能
- ✅ Protobuf 序列化，体积小速度快
- ✅ 类型安全，编译时检查
- ✅ 双向流、客户端流、服务端流

### 代码生成

```bash
cd gRPC/demo/proto
protoc --go_out=. --go-grpc_out=. calc.proto
```

生成文件：
- `calc.pb.go` - 数据结构
- `calc_grpc.pb.go` - 服务端和客户端代码

---

## Kitex 教程

### 功能特性

- ✅ 高性能（字节跳动内部大规模使用）
- ✅ Go 语言专属，API 简洁
- ✅ 支持 Thrift 和 Protobuf
- ✅ 内置服务治理能力
- ✅ 代码生成便捷

### 代码生成

```bash
cd Kitex/demo
kitex -module demo -service calc proto/calc.thrift
```

生成文件：
- `kitex_gen/` - 自动生成的代码
- `handler.go` - 业务逻辑模板
- `main.go` - 服务启动文件

---

## RPC 核心概念

### 什么是 RPC？

**RPC（Remote Procedure Call）** = 远程过程调用

简单来说：**像调用本地函数一样，调用远程服务器上的函数**

### RPC vs HTTP

| 特性 | RPC（gRPC/Kitex） | HTTP + JSON |
|------|-------------------|-------------|
| 性能 | ⭐⭐⭐⭐⭐ 极高 | ⭐⭐ 一般 |
| 序列化 | Protobuf/Thrift（二进制）| JSON（文本）|
| 类型安全 | ✅ 编译时检查 | ❌ 运行时检查 |
| 浏览器支持 | ❌ 不直接支持 | ✅ 完美支持 |
| 适用场景 | 内部微服务调用 | 对外 API、前端 |

### 核心流程

```
1. 定义 IDL（.proto / .thrift）
         ↓
2. 用工具生成代码
         ↓
3. 实现服务端业务逻辑
         ↓
4. 实现客户端调用代码
         ↓
5. 先启动服务端，再运行客户端测试
```

---

## 常见问题

### gRPC 常见问题

**Q: `protoc-gen-go: unable to determine Go import path`**
**A:** 确保 proto 文件中有 `option go_package = "./;calc";`

**Q: `package demo/proto is not in std`**
**A:** 确保在 `demo` 目录下运行，不要在 `server` 或 `client` 目录下

---

### Kitex 常见问题

**Q: `kitex : The term 'kitex' is not recognized`**
**A:** 先安装 kitex：`go install github.com/cloudwego/kitex/tool/cmd/kitex@latest`

**Q: 重新生成后 handler.go 被覆盖了**
**A:** 重新生成前务必备份 handler.go
