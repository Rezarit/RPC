# Kitex 完整教程详解

本教程包含三个部分：接口定义、服务端实现、客户端实现，适合预习和复习使用。

---

## 目录
- [一、Thrift 接口定义文件详解](#section-1)
- [二、Kitex 服务端实现详解](#section-2)
- [三、Kitex 客户端实现详解](#section-3)
- [四、如何运行项目](#section-4)

---

<a id="section-1"></a>
## 一、Thrift 接口定义文件详解

### 1. 文件整体介绍
这是 **Kitex 的核心接口定义文件（后缀：`.thrift`）**，全称 Thrift IDL（接口定义语言）。

作用：**用统一的语法定义远程调用的「请求参数、返回结果、服务方法」**，跨语言通用，执行命令后可自动生成 Go 代码，不用手写数据结构。

本次代码功能：定义一个**加减乘除四则运算远程调用服务**。

---

### 2. 完整代码
```thrift
namespace go demo.calc

// 加法请求
struct AddReq{
    1:required i64 A
    2:required i64 B
}

// 加法响应
struct AddResp{
    1:required i64 Sum
}

// 减法请求
struct SubReq{
    1:required i64 A
    2:required i64 B
}

// 减法响应
struct SubResp{
    1:required i64 Difference
}

// 乘法请求
struct MulReq{
    1:required i64 A
    2:required i64 B
}

// 乘法响应
struct MulResp{
    1:required i64 Product
}

// 除法请求
struct DivReq{
    1:required i64 A
    2:required i64 B
}

// 除法响应
struct DivResp{
    1:required i64 Quotient  // 商
    2:required i64 Remainder // 余数
}

service CalcService{
    AddResp Add(1:AddReq req)
    SubResp Sub(1:SubReq req)
    MulResp Mul(1:MulReq req)
    DivResp Div(1:DivReq req)
}
```

---

### 3. 逐行逐段详细解释

#### 3.1 Go 命名空间声明
```thrift
namespace go demo.calc
```
- 作用：**指定生成的 Go 代码的包路径**
- 格式：`namespace go 「包路径」`
  - `demo.calc`：生成的代码包路径为 `demo/calc`
- 说明：这告诉编译器，生成的 Go 代码应该放在哪个包下面

---

#### 3.2 定义「加法请求消息体」AddReq
```thrift
// 加法请求
struct AddReq{
    1:required i64 A
    2:required i64 B
}
```
- `struct`：Thrift 关键字，**定义消息结构体**（等价于 Go 的 struct）
- `AddReq`：消息名，全称 Add Request = **加法请求**
- 内部字段解释：

  | 代码          | 含义                                   |
  | :------------ | :------------------------------------- |
  | `1:`          | 字段编号：**唯一标识，序列化时使用** |
  | `required`    | 字段是必填的（proto3 没有这个概念）|
  | `i64`         | 字段类型：64位整数（等价Go的int64）|
  | `A` / `B`     | 字段名：两个加数                       |

✅ 核心重点：
字段编号 `1`、`2` **不能重复、不能修改**，是数据传输的唯一标识。

---

#### 3.3 定义「加法响应消息体」AddResp
```thrift
// 加法响应
struct AddResp{
    1:required i64 Sum
}
```
- `AddResp`：全称 Add Response = **加法响应**
- `Sum`：求和计算的**结果字段**

---

#### 3.4 定义「减法请求/响应」
```thrift
// 减法请求
struct SubReq{
    1:required i64 A
    2:required i64 B
}

// 减法响应
struct SubResp{
    1:required i64 Difference
}
```
- `SubReq`：减法请求
- `SubResp`：减法响应
- `Difference`：差值（A - B）

---

#### 3.5 定义「乘法请求/响应」
```thrift
// 乘法请求
struct MulReq{
    1:required i64 A
    2:required i64 B
}

// 乘法响应
struct MulResp{
    1:required i64 Product
}
```
- `MulReq`：乘法请求
- `MulResp`：乘法响应
- `Product`：乘积（A × B）

---

#### 3.6 定义「除法请求/响应」
```thrift
// 除法请求
struct DivReq{
    1:required i64 A
    2:required i64 B
}

// 除法响应
struct DivResp{
    1:required i64 Quotient  // 商
    2:required i64 Remainder // 余数
}
```
- `DivReq`：除法请求
- `DivResp`：除法响应
- `Quotient`：商（A ÷ B 的整数部分）
- `Remainder`：余数（A % B）

---

#### 3.7 定义 Kitex 服务
```thrift
service CalcService{
    AddResp Add(1:AddReq req)
    SubResp Sub(1:SubReq req)
    MulResp Mul(1:MulReq req)
    DivResp Div(1:DivReq req)
}
```
- `service`：关键字，**定义一个 Kitex 服务**
- `CalcService`：服务名称（计算器服务）
- `AddResp Add(1:AddReq req)`：**定义一个远程调用方法**
  - `AddResp`：**返回值**（服务端返回给客户端的数据）
  - `Add`：方法名（加法方法）
  - `1:AddReq req`：**入参**（客户端传给服务端的数据）
    - `1:`：参数编号
    - `AddReq`：参数类型
    - `req`：参数名

✅ 一句话总结：
客户端调用 `Add/Sub/Mul/Div` 方法，传入两个数字，服务端计算后返回结果。

---

### 4. 这份代码的核心作用
1. 它是 **Kitex 的「接口说明书」**
2. 不依赖任何语言，Java/Python/Go 都能用
3. 执行命令后**自动生成代码**，不用手写结构体
4. 定义了：传什么参数 → 返回什么结果 → 调用什么方法

---

### 5. 生成 Go 代码（详细步骤）

#### 5.1 前置条件：安装必要工具

在生成代码之前，需要先安装以下工具：

**1. 安装 Kitex 工具**
打开终端，执行：
```bash
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```

验证安装：
```bash
kitex --version
```

能输出版本号说明安装成功！

注意：
- 如果提示 `kitex : The term 'kitex' is not recognized`，说明 kitex 不在 PATH 中
- Windows 下默认安装在 `C:\Users\你的用户名\go\bin`
- 可以用完整路径运行：`C:\Users\你的用户名\go\bin\kitex.exe`

---

#### 5.2 执行生成命令

进入 `Kitex/demo` 目录，执行：

```bash
kitex -module demo -service calc proto/calc.thrift
```

命令参数详解：
- `kitex`：调用 kitex 工具
- `-module demo`：指定 Go module 名称（必须和 go.mod 中的 module 一致）
- `-service calc`：指定服务名称，会自动生成 main.go
- `proto/calc.thrift`：要编译的 thrift 文件路径

---

#### 5.3 生成的文件说明

执行成功后，会生成以下目录和文件：

**1. kitex_gen 目录**
这是自动生成的代码目录，结构如下：
```
kitex_gen/
└── demo/
    └── calc/
        ├── calc.go              # 数据结构定义
        ├── k-calc.go            # Kitex 内部用
        ├── k-consts.go          # 常量定义
        └── calcservice/
            ├── calcservice.go    # 服务接口定义
            ├── client.go         # 客户端代码
            └── server.go         # 服务端代码
```

**2. calc.go（数据结构）**
- 作用：Thrift 数据结构定义
- 包含内容：
  - `AddReq` / `AddResp` 结构体
  - `SubReq` / `SubResp` 结构体
  - `MulReq` / `MulResp` 结构体
  - `DivReq` / `DivResp` 结构体
  - 序列化/反序列化方法
  - Getter/Setter 方法
- **注意**：不要手动修改这个文件！

**3. calcservice/ 目录**
- `calcservice.go`：服务接口定义
  - `CalcService` 接口（服务端需要实现）
- `client.go`：客户端调用代码
  - `NewClient` / `MustNewClient`：创建客户端
  - 各方法的客户端调用实现
- `server.go`：服务端代码
  - `NewServer`：创建服务器
  - 服务注册逻辑
- **注意**：不要手动修改这些文件！

**4. 其他自动生成的文件**
- `handler.go`：业务逻辑模板（**这个文件可以修改！**）
- `main.go`：服务启动文件（**这个文件可以修改！**）
- `kitex_info.yaml`：Kitex 信息文件
- `build.sh` / `script/bootstrap.sh`：构建脚本

---

#### 5.4 重新生成代码的注意事项

**重要提示：**
1. **kitex_gen 目录下的所有文件都不要手动修改！**
2. **handler.go 和 main.go 可以修改**，但重新生成时会被覆盖！
3. 重新生成前，建议备份 handler.go 和 main.go

**如果修改了 thrift 文件，需要重新生成：**
```bash
# 1. 备份你的 handler.go 和 main.go
copy handler.go handler.go.backup
copy main.go main.go.backup

# 2. 重新生成代码
kitex -module demo -service calc proto/calc.thrift

# 3. 从备份中恢复你的业务逻辑
# 把 handler.go.backup 中的业务逻辑复制回新的 handler.go
```

---

#### 5.5 常见问题

**Q1: 报错 `kitex : The term 'kitex' is not recognized`**
**原因**：kitex 没有安装或不在 PATH 中
**解决**：
1. 确认已执行 `go install` 命令
2. 确认 `GOPATH/bin` 在系统 PATH 中
3. 或者使用完整路径：`C:\Users\你的用户名\go\bin\kitex.exe`

**Q2: 报错 `go.mod file not found`**
**原因**：没有在正确的目录下执行命令
**解决**：确保在 `Kitex/demo` 目录下执行（有 go.mod 的目录）

**Q3: 报错 `module name does not match`**
**原因**：`-module` 参数和 go.mod 中的 module 不一致
**解决**：检查 go.mod 第一行，用正确的 module 名

**Q4: 重新生成后 handler.go 被覆盖了**
**原因**：kitex 重新生成时会重置 handler.go 为模板
**解决**：
- 重新生成前备份 handler.go
- 或者用 git 管理代码，重新生成后 git diff 查看变化
- 把你的业务逻辑复制回去

**Q5: 生成的代码在哪里？**
**答案**：在 `kitex_gen` 目录下，不要手动修改！

---

<a id="section-2"></a>
## 二、Kitex 服务端实现详解

### 1. 文件整体介绍
Kitex 服务端由两个文件组成：
- **handler.go**：业务逻辑实现，实现四则运算
- **main.go**：启动服务器，监听端口

---

### 2. handler.go - 业务逻辑实现

#### 2.1 完整代码
```go
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
```

---

#### 2.2 逐行逐段详细解释

##### 2.2.1 包声明和导入
```go
package main

import (
	"context"
	calc "demo/kitex_gen/demo/calc"
)
```
- `package main`：声明这是一个可执行程序（不是库）
- `import`：导入需要的包
  - `context`：上下文，用于控制请求的超时、取消等
  - `calc "demo/kitex_gen/demo/calc"`：导入生成的 thrift 代码，别名为 `calc`

---

##### 2.2.2 定义服务实现结构体
```go
// CalcServiceImpl implements the last service interface defined in the IDL.
type CalcServiceImpl struct{}
```
- `CalcServiceImpl`：我们自己定义的服务实现结构体
- 这个结构体需要实现 thrift 文件中定义的 `CalcService` 接口的所有方法

---

##### 2.2.3 实现 Add 方法
```go
// Add 实现加法
func (s *CalcServiceImpl) Add(ctx context.Context, req *calc.AddReq) (resp *calc.AddResp, err error) {
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

##### 2.2.4 实现 Sub 方法
```go
// Sub 实现减法
func (s *CalcServiceImpl) Sub(ctx context.Context, req *calc.SubReq) (resp *calc.SubResp, err error) {
	difference := req.A - req.B
	return &calc.SubResp{Difference: difference}, nil
}
```
- 减法运算：`req.A - req.B`
- 返回差值 `Difference`

---

##### 2.2.5 实现 Mul 方法
```go
// Mul 实现乘法
func (s *CalcServiceImpl) Mul(ctx context.Context, req *calc.MulReq) (resp *calc.MulResp, err error) {
	product := req.A * req.B
	return &calc.MulResp{Product: product}, nil
}
```
- 乘法运算：`req.A * req.B`
- 返回乘积 `Product`

---

##### 2.2.6 实现 Div 方法
```go
// Div 实现除法，返回商和余数
func (s *CalcServiceImpl) Div(ctx context.Context, req *calc.DivReq) (resp *calc.DivResp, err error) {
	if req.B == 0 {
		return &calc.DivResp{Quotient: 0, Remainder: 0}, nil
	}
	quotient := req.A / req.B
	remainder := req.A % req.B
	return &calc.DivResp{Quotient: quotient, Remainder: remainder}, nil
}
```
- 除法运算，需要特别处理**除零**的情况
- 如果 `req.B == 0`，返回默认值
- 否则：
  - `quotient = req.A / req.B`：商（整数除法）
  - `remainder = req.A % req.B`：余数
- 返回商和余数两个结果

---

### 3. main.go - 启动服务器

#### 3.1 完整代码
```go
package main

import (
	calc "demo/kitex_gen/demo/calc/calcservice"
	"log"
)

func main() {
	svr := calc.NewServer(new(CalcServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
```

---

#### 3.2 逐行逐段详细解释

##### 3.2.1 包声明和导入
```go
package main

import (
	calc "demo/kitex_gen/demo/calc/calcservice"
	"log"
)
```
- `calc "demo/kitex_gen/demo/calc/calcservice"`：导入生成的服务端代码
- `log`：日志输出

---

##### 3.2.2 创建服务器
```go
	svr := calc.NewServer(new(CalcServiceImpl))
```
- `calc.NewServer(...)`：创建 Kitex 服务器
- `new(CalcServiceImpl)`：创建我们的服务实现实例
- 把服务实现注册到服务器上

---

##### 3.2.3 启动服务器
```go
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
```
- `svr.Run()`：**启动服务**
  - 这是一个**阻塞调用**，会一直运行直到服务器停止
  - 默认监听 `:8888` 端口
- 如果启动出错，打印错误信息

---

<a id="section-3"></a>
## 三、Kitex 客户端实现详解

### 1. 文件整体介绍
这是 **Kitex 客户端代码**，负责：
1. 连接到 Kitex 服务器
2. 调用远程服务方法（加减乘除）
3. 接收并处理服务端返回的结果

---

### 2. 完整代码
```go
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
```

---

### 3. 逐行逐段详细解释

#### 3.1 包声明和导入
```go
package main

import (
	"context"
	"demo/kitex_gen/demo/calc"
	"demo/kitex_gen/demo/calc/calcservice"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
)
```
- `demo/kitex_gen/demo/calc`：导入生成的数据结构
- `demo/kitex_gen/demo/calc/calcservice`：导入生成的客户端代码
- `github.com/cloudwego/kitex/client`：Kitex 客户端库

---

#### 3.2 建立连接
```go
	// 创建客户端连接
	client := calcservice.MustNewClient("calc",
		client.WithHostPorts("127.0.0.1:8888"),
	)
```
- `calcservice.MustNewClient(...)`：创建 Kitex 客户端
- 参数 1：`"calc"` - 服务名称
- 参数 2：`client.WithHostPorts("127.0.0.1:8888")`
  - 配置服务器地址
  - `127.0.0.1`：本机（如果服务器在其他机器，填 IP 地址）
  - `8888`：端口号，必须和服务端一致
- `MustNewClient`：如果创建失败会 panic（还有 `NewClient` 会返回 error）

---

#### 3.3 测试加法
```go
	// 测试加法
	fmt.Println("=== 测试加法 ===")
	addReq := &calc.AddReq{A: 10, B: 5}
	addResp, err := client.Add(context.Background(), addReq)
	if err != nil {
		log.Fatalf("Add failed: %v", err)
	}
	fmt.Printf("%d + %d = %d\n\n", addReq.A, addReq.B, addResp.Sum)
```
- `&calc.AddReq{A: 10, B: 5}`：创建加法请求
- `client.Add(...)`：**调用远程的 Add 方法**
  - 看起来像调用本地方法，实际上是通过网络远程调用
- `context.Background()`：创建一个空的上下文
- 错误检查：如果调用失败，打印错误并退出
- 打印结果：`10 + 5 = 15`

---

#### 3.4 测试减法
```go
	// 测试减法
	fmt.Println("=== 测试减法 ===")
	subReq := &calc.SubReq{A: 10, B: 5}
	subResp, err := client.Sub(context.Background(), subReq)
	if err != nil {
		log.Fatalf("Sub failed: %v", err)
	}
	fmt.Printf("%d - %d = %d\n\n", subReq.A, subReq.B, subResp.Difference)
```
- 调用 `client.Sub(...)` 方法
- 打印结果：`10 - 5 = 5`

---

#### 3.5 测试乘法
```go
	// 测试乘法
	fmt.Println("=== 测试乘法 ===")
	mulReq := &calc.MulReq{A: 10, B: 5}
	mulResp, err := client.Mul(context.Background(), mulReq)
	if err != nil {
		log.Fatalf("Mul failed: %v", err)
	}
	fmt.Printf("%d × %d = %d\n\n", mulReq.A, mulReq.B, mulResp.Product)
```
- 调用 `client.Mul(...)` 方法
- 打印结果：`10 × 5 = 50`

---

#### 3.6 测试除法
```go
	// 测试除法
	fmt.Println("=== 测试除法 ===")
	divReq := &calc.DivReq{A: 10, B: 3}
	divResp, err := client.Div(context.Background(), divReq)
	if err != nil {
		log.Fatalf("Div failed: %v", err)
	}
	fmt.Printf("%d ÷ %d = %d 余 %d\n\n", divReq.A, divReq.B, divResp.Quotient, divResp.Remainder)
```
- 调用 `client.Div(...)` 方法
- 打印结果：`10 ÷ 3 = 3 余 1`

---

#### 3.7 完成提示
```go
	fmt.Println("所有测试完成！🎉")
```

---

<a id="section-4"></a>
## 四、如何运行项目

### 1. 项目结构
```
Kitex/demo/
├── proto/
│   └── calc.thrift         # 接口定义文件
├── kitex_gen/              # 自动生成的代码
│   └── demo/calc/
│       ├── calc.go         # 数据结构
│       └── calcservice/    # 服务端/客户端代码
├── handler.go              # 业务逻辑实现
├── main.go                 # 服务端启动文件
├── client/
│   └── client.go           # 客户端代码
├── go.mod                  # Go 模块文件
└── go.sum                  # 依赖版本锁定文件
```

---

### 2. 前置条件
确保已安装：
- Go 1.19+
- Kitex 工具

安装 Kitex：
```bash
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```

---

### 3. 重新生成代码（如果修改了 thrift 文件）
```bash
cd e:\Go\Works\rpc\Kitex\demo
kitex -module demo -service calc proto/calc.thrift
```

注意：重新生成代码后，`handler.go` 可能会被重置为默认模板，需要重新添加业务逻辑！

---

### 4. 运行步骤

#### 步骤 1：启动服务端
打开一个终端，执行：
```powershell
cd e:\Go\Works\rpc\Kitex\demo
go run .
```
看到服务启动（默认监听 8888 端口），说明服务端启动成功！**不要关闭这个终端**。

---

#### 步骤 2：运行客户端
打开**另一个新终端**，执行：
```powershell
cd e:\Go\Works\rpc\Kitex\demo
go run ./client
```
看到输出：
```
=== 测试加法 ===
10 + 5 = 15

=== 测试减法 ===
10 - 5 = 5

=== 测试乘法 ===
10 × 5 = 50

=== 测试除法 ===
10 ÷ 3 = 3 余 1

所有测试完成！🎉
```
说明客户端调用成功！

---

### 5. 常见问题

#### Q1: 报错 `undefined: calc.SubResp`
**原因**：修改了 thrift 文件，但没有重新生成 kitex 代码
**解决**：重新生成代码
```bash
kitex -module demo -service calc proto/calc.thrift
```
注意：重新生成后需要重新实现 handler.go 中的业务逻辑！

---

#### Q2: 报错 `kitex : The term 'kitex' is not recognized`
**原因**：kitex 没有安装或不在 PATH 中
**解决**：
1. 先安装 kitex：`go install github.com/cloudwego/kitex/tool/cmd/kitex@latest`
2. 或者使用完整路径：`C:\Users\你的用户名\go\bin\kitex.exe`

---

#### Q3: 客户端报错 `connection refused`
**原因**：服务端没有启动
**解决**：先启动服务端，再运行客户端

---

#### Q4: 重新生成代码后 handler.go 被覆盖了
**原因**：kitex 重新生成时会重置 handler.go 为模板
**解决**：
- 重新生成前备份 handler.go
- 或者重新生成后，重新添加业务逻辑

---

## 五、总结

### Kitex 工作流程
1. **定义接口**：写 `.thrift` 文件，定义请求、响应、服务方法
2. **生成代码**：用 kitex 命令生成 Go 代码
3. **实现服务端**：写 handler.go 实现业务逻辑
4. **实现客户端**：写 client.go 调用远程方法
5. **运行测试**：先启动服务端，再运行客户端

### 核心概念
- **thrift 文件**：接口定义语言，跨语言通用
- **service**：定义 Kitex 服务
- **struct**：定义请求和响应的数据结构
- **handler.go**：服务端业务逻辑实现
- **main.go**：服务端启动入口
- **client.go**：客户端调用代码

---

祝你学习愉快！🎉
