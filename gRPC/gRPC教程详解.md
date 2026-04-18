# gRPC 完整教程详解

本教程包含三个部分：接口定义、服务端实现、客户端实现，适合预习和复习使用。

---

## 目录
- [一、gRPC 接口定义文件详解](#section-1)
- [二、gRPC 服务端实现详解](#section-2)
- [三、gRPC 客户端实现详解](#section-3)
- [四、如何运行项目](#section-4)

---

<a id="section-1"></a>
## 一、gRPC 接口定义文件详解

### 1. 文件整体介绍
这是 **gRPC 的核心接口定义文件（后缀：`.proto`）**，全称 Protobuf IDL（接口定义语言）。

作用：**用统一的语法定义远程调用的「请求参数、返回结果、服务方法」**，跨语言通用，执行命令后可自动生成 Go 代码，不用手写数据结构。

本次代码功能：定义一个**加法远程调用服务**。

---

### 2. 完整代码
```protobuf
syntax = "proto3";

package calc;

option go_package = "./;calc";

message AddReq {
    int32 a = 1;
    int32 b = 2;
}

message AddResp {
    int32 sum = 1;
}

service Calc {
    rpc Add(AddReq) returns (AddResp);
}
```

---

### 3. 逐行逐段详细解释

#### 3.1 语法版本声明
```protobuf
syntax = "proto3";
```
- 作用：**指定 Protobuf 语法版本为 proto3**
- 说明：
  - Protobuf 主流版本：`proto2` / `proto3`
  - `proto3` 是最新版，语法更简洁、无必填字段校验，是 gRPC 开发的**标准选择**
  - 必须写在文件**第一行**，是整个文件的开头声明

---

#### 3.2 包名声明
```protobuf
package calc;
```
- 作用：**定义命名空间（包名），防止命名冲突**
- 说明：
  - 相当于 Java 的 `package`、Go 的 `package`
  - 这里命名为 `calc`，代表这是一个「计算器」相关的服务
  - 多个 proto 文件同时使用时，包名可以区分不同服务

---

#### 3.3 Go 代码生成配置
```protobuf
option go_package = "./;calc";
```
- 作用：**告诉编译器：生成的 Go 代码放在哪里、包名是什么**
- 格式固定：`「生成路径;Go包名」`
  - `./`：生成的 Go 代码文件放在**当前目录**
  - `calc`：生成的 Go 代码的包名为 `calc`
- 补充：
  - 这是**Go 语言专属配置**，不写无法生成 Go 代码
  - 标准写法也可以是 `./calc`，效果一致

---

#### 3.4 定义「请求消息体」AddReq
```protobuf
message AddReq {
    int32 a = 1;
    int32 b = 2;
}
```
- `message`：Protobuf 关键字，**定义消息结构体**（等价于 Go 的 struct）
- `AddReq`：消息名，全称 Add Request = **加法请求**
- 内部字段解释：

  | 代码      | 含义                                   |
  | :-------- | :------------------------------------- |
  | `int32`   | 字段类型：32位整数（等价Go的int32）|
  | `a` / `b` | 字段名：两个加数                       |
  | `= 1`     | 字段编号：**唯一标识，序列化时使用** |

✅ 核心重点：
字段编号 `1`、`2` **不能重复、不能修改**，是数据传输的唯一标识。

---

#### 3.5 定义「响应消息体」AddResp
```protobuf
message AddResp {
    int32 sum = 1;
}
```
- `AddResp`：全称 Add Response = **加法响应**
- `sum`：求和计算的**结果字段**
- `= 1`：字段唯一编号

---

#### 3.6 定义 gRPC 服务
```protobuf
service Calc {
    rpc Add(AddReq) returns (AddResp);
}
```
- `service`：关键字，**定义一个 gRPC 服务**
- `Calc`：服务名称（计算器服务）
- `rpc Add(...) returns(...)`：**定义一个远程调用方法**
  - `rpc`：远程过程调用（Remote Procedure Call）
  - `Add`：方法名（加法方法）
  - `AddReq`：**入参**（客户端传给服务端的数据）
  - `AddResp`：**返回值**（服务端返回给客户端的数据）

✅ 一句话总结：
客户端调用 `Add` 方法，传入两个数字，服务端计算后返回求和结果。

---

### 4. 这份代码的核心作用
1. 它是 **gRPC 的「接口说明书」**
2. 不依赖任何语言，Java/Python/Go 都能用
3. 执行命令后**自动生成代码**，不用手写结构体
4. 定义了：传什么参数 → 返回什么结果 → 调用什么方法

---

### 5. 生成 Go 代码（详细步骤）

#### 5.1 前置条件：安装必要工具

在生成代码之前，需要先安装以下工具：

**1. 安装 protoc 编译器**
- 下载地址：https://github.com/protocolbuffers/protobuf/releases
- 下载对应系统的版本（Windows 下载 `protoc-xx.x-win64.zip`）
- 解压后，把 `bin` 目录添加到系统 PATH 环境变量

**2. 安装 Go 插件**
打开终端，执行：
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

这两个插件的作用：
- `protoc-gen-go`：生成 Protobuf 数据结构代码
- `protoc-gen-go-grpc`：生成 gRPC 服务端和客户端代码

验证安装：
```bash
protoc --version
protoc-gen-go --version
protoc-gen-go-grpc --version
```

都能输出版本号说明安装成功！

---

#### 5.2 执行生成命令

进入 `gRPC/demo/proto` 目录，执行：

```bash
protoc --go_out=. --go-grpc_out=. calc.proto
```

命令参数详解：
- `protoc`：调用 protoc 编译器
- `--go_out=.`：生成 Go 数据结构代码，输出到当前目录
- `--go-grpc_out=.`：生成 gRPC 代码，输出到当前目录
- `calc.proto`：要编译的 proto 文件

---

#### 5.3 生成的文件说明

执行成功后，会在 `proto` 目录下生成两个文件：

**1. calc.pb.go**
- 作用：数据结构定义
- 包含内容：
  - `AddReq` 结构体
  - `AddResp` 结构体
  - 序列化/反序列化方法
  - Getter/Setter 方法
- **注意**：不要手动修改这个文件！

**2. calc_grpc.pb.go**
- 作用：gRPC 服务端和客户端代码
- 包含内容：
  - `CalcServer` 接口（服务端需要实现）
  - `UnimplementedCalcServer`（向前兼容用）
  - `RegisterCalcServer`（注册服务用）
  - `CalcClient` 接口（客户端调用用）
  - `NewCalcClient`（创建客户端用）
- **注意**：不要手动修改这个文件！

---

#### 5.4 常见问题

**Q1: 报错 `protoc-gen-go: program not found`**
**原因**：protoc-gen-go 没有安装或不在 PATH 中
**解决**：
1. 确认已执行 `go install` 命令
2. 确认 `GOPATH/bin` 在系统 PATH 中
3. Windows 下默认在 `C:\Users\你的用户名\go\bin`

**Q2: 报错 `proto file not found`**
**原因**：没有在正确的目录下执行命令
**解决**：确保在 `gRPC/demo/proto` 目录下执行

**Q3: 生成后修改了 proto 文件怎么办？**
**解决**：重新执行生成命令，会覆盖旧的生成文件

---

<a id="section-2"></a>
## 二、gRPC 服务端实现详解

### 1. 文件整体介绍
这是 **gRPC 服务端代码**，负责：
1. 监听网络端口
2. 实现 proto 文件中定义的服务接口
3. 处理客户端的远程调用请求

---

### 2. 完整代码
```go
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
```

---

### 3. 逐行逐段详细解释

#### 3.1 包声明和导入
```go
package main

import (
	"context"
	"log"
	"net"

	calc "demo/proto"

	"google.golang.org/grpc"
)
```
- `package main`：声明这是一个可执行程序（不是库）
- `import`：导入需要的包
  - `context`：上下文，用于控制请求的超时、取消等
  - `log`：日志输出
  - `net`：网络编程，用于创建监听器
  - `calc "demo/proto"`：导入生成的 proto 代码，别名为 `calc`
  - `google.golang.org/grpc`：gRPC 核心库

---

#### 3.2 定义服务实现结构体
```go
// CalcServer 实现了 gRPC 服务接口
// 嵌入 calc.UnimplementedCalcServer 以确保向前兼容性
type CalcServer struct {
	calc.UnimplementedCalcServer
}
```
- `CalcServer`：我们自己定义的服务实现结构体
- `calc.UnimplementedCalcServer`：**必须嵌入**
  - 这是 gRPC 生成的未实现服务接口
  - 作用：确保向前兼容性，当 proto 文件新增方法时不会编译报错
  - 即使我们不嵌入它，代码也能运行，但推荐嵌入

---

#### 3.3 实现 Add 方法
```go
func (s *CalcServer) Add(ctx context.Context, req *calc.AddReq) (*calc.AddResp, error) {
	sum := req.A + req.B
	return &calc.AddResp{Sum: sum}, nil
}
```
- 这是**真正的业务逻辑**
- 参数说明：
  - `ctx`：上下文，可以获取请求元数据、设置超时等
  - `req`：客户端传来的请求，包含 `A` 和 `B` 两个字段
- 返回值说明：
  - `*calc.AddResp`：响应结果，包含 `Sum` 字段
  - `error`：错误信息，如果有的话
- 业务逻辑：简单的加法运算 `req.A + req.B`

---

#### 3.4 main 函数 - 启动服务
```go
func main() {
	// 1. 创建 TCP 监听器，监听 8888 端口
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
```
- `net.Listen("tcp", ":8888")`：创建 TCP 监听器
  - `"tcp"`：使用 TCP 协议
  - `":8888"`：监听本机的 8888 端口
- 错误检查：如果监听失败（比如端口被占用），打印错误并退出

```go
	// 2. 创建 gRPC 服务器实例
	server := grpc.NewServer()
```
- `grpc.NewServer()`：创建一个新的 gRPC 服务器

```go
	// 3. 将我们的服务实现注册到 gRPC 服务器
	calc.RegisterCalcServer(server, &CalcServer{})
```
- `calc.RegisterCalcServer(...)`：这是自动生成的注册函数
- 把我们的 `CalcServer` 实例注册到 gRPC 服务器上
- 这样 gRPC 服务器就知道：收到请求后找谁处理

```go
	// 4. 启动 gRPC 服务，开始处理请求
	// Serve 会阻塞直到服务器停止或出错
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```
- `server.Serve(lis)`：**启动服务**
  - 这是一个**阻塞调用**，会一直运行直到服务器停止
  - 开始接受并处理客户端的请求

---

<a id="section-3"></a>
## 三、gRPC 客户端实现详解

### 1. 文件整体介绍
这是 **gRPC 客户端代码**，负责：
1. 连接到 gRPC 服务器
2. 调用远程服务方法
3. 接收并处理服务端返回的结果

---

### 2. 完整代码
```go
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
```

---

### 3. 逐行逐段详细解释

#### 3.1 包声明和导入
```go
package main

import (
	"context"
	"log"

	calc "demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
```
- 注意多了一个导入：`google.golang.org/grpc/credentials/insecure`
- 这个包用于创建**不安全的连接**（不使用 TLS 加密）
- 生产环境应该使用 TLS，但开发环境用 insecure 更方便

---

#### 3.2 建立连接
```go
	// 建立与 gRPC 服务器的连接
	// 服务器地址 (localhost:8888)
	// 使用不安全的凭证（不使用 TLS 加密，适合开发环境）
	conn, _ := grpc.NewClient(
		"localhost:8888",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
```
- `grpc.NewClient(...)`：创建 gRPC 客户端连接
- 参数 1：`"localhost:8888"` - 服务器地址
  - `localhost`：本机（如果服务器在其他机器，填 IP 地址）
  - `8888`：端口号，必须和服务端一致
- 参数 2：`grpc.WithTransportCredentials(insecure.NewCredentials())`
  - 配置使用**不安全的连接**（不加密）
  - 生产环境应该用 `tls.NewCredentials()`

```go
	// 确保在函数退出时关闭连接
	defer conn.Close()
```
- `defer conn.Close()`：**非常重要**
- `defer` 会在函数退出时自动执行
- 确保连接被正确关闭，防止资源泄漏

---

#### 3.3 创建客户端实例
```go
	// 创建 gRPC 客户端实例
	client := calc.NewCalcClient(conn)
```
- `calc.NewCalcClient(conn)`：这是自动生成的客户端构造函数
- 传入连接对象，得到一个可以调用远程方法的客户端

---

#### 3.4 调用远程方法
```go
	// 调用远程的 Add 方法
	// context.Background(): 创建一个空的上下文
	resp, _ := client.Add(context.Background(), &calc.AddReq{A: 5, B: 10})
```
- `client.Add(...)`：**调用远程的 Add 方法**
  - 看起来像调用本地方法，实际上是通过网络远程调用
- 参数 1：`context.Background()` - 上下文
  - 可以用于设置超时、取消请求等
  - `Background()` 创建一个空的上下文
- 参数 2：`&calc.AddReq{A: 5, B: 10}` - 请求参数
  - 创建一个请求对象，设置 `A=5`，`B=10`
- 返回值：
  - `resp`：响应结果
  - `_`：忽略错误（实际开发中应该处理错误）

---

#### 3.5 处理结果
```go
	// 打印服务器返回的结果
	log.Println(resp.Sum)
```
- `resp.Sum`：从响应中取出计算结果
- `log.Println(...)`：打印结果，应该输出 `15`（5 + 10）

---

<a id="section-4"></a>
## 四、如何运行项目

### 1. 项目结构
```
gRPC/demo/
├── proto/
│   ├── calc.proto          # 接口定义文件
│   ├── calc.pb.go          # 自动生成的数据结构
│   └── calc_grpc.pb.go     # 自动生成的服务端/客户端代码
├── server/
│   └── server.go           # 服务端代码
├── client/
│   └── client.go           # 客户端代码
├── go.mod                  # Go 模块文件
└── go.sum                  # 依赖版本锁定文件
```

---

### 2. 前置条件
确保已安装：
- Go 1.19+
- protoc 编译器
- protoc-gen-go 和 protoc-gen-go-grpc 插件

---

### 3. 运行步骤

#### 步骤 1：启动服务端
打开一个终端，执行：
```powershell
cd e:\Go\Works\rpc\gRPC\demo
go run ./server
```
看到输出：
```
2026/04/15 XX:XX:XX server listening at [::]:8888
```
说明服务端启动成功！**不要关闭这个终端**。

---

#### 步骤 2：运行客户端
打开**另一个新终端**，执行：
```powershell
cd e:\Go\Works\rpc\gRPC\demo
go run ./client
```
看到输出：
```
2026/04/15 XX:XX:XX 15
```
说明客户端调用成功，结果是 15（5 + 10）！

---

### 4. 常见问题

#### Q1: 报错 `package demo/proto is not in std`
**原因**：在错误的目录下运行了命令
**解决**：确保在 `demo` 目录下运行，不要在 `server` 或 `client` 目录下运行

---

#### Q2: 报错 `listen tcp :8888: bind: address already in use`
**原因**：端口 8888 已被占用
**解决**：
1. 找到并关闭占用端口的程序，或
2. 修改代码中的端口号（比如改成 8889）

---

#### Q3: 客户端报错 `connection refused`
**原因**：服务端没有启动
**解决**：先启动服务端，再运行客户端

---

## 五、总结

### gRPC 工作流程
1. **定义接口**：写 `.proto` 文件，定义请求、响应、服务方法
2. **生成代码**：用 protoc 命令生成 Go 代码
3. **实现服务端**：写服务端代码，实现业务逻辑
4. **实现客户端**：写客户端代码，调用远程方法
5. **运行测试**：先启动服务端，再运行客户端

### 核心概念
- **proto 文件**：接口定义语言，跨语言通用
- **service**：定义 gRPC 服务
- **message**：定义请求和响应的数据结构
- **服务端**：实现服务接口，监听端口，处理请求
- **客户端**：连接服务端，调用远程方法

---

祝你学习愉快！🎉
