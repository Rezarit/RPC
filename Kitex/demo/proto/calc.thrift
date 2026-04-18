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
