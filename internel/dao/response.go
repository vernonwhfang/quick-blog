package dao

type Response struct {
	Code RespCodeEnum `json:"code"`
	Msg  string       `json:"msg,omitempty"`
	Data interface{}  `json:"data,omitempty"`
}

const (
	CodeRespFail    = -1
	CodeRespSuccess = 200

	MsgSuccess   = "success"
	MsgFailValid = "校验不通过"
)

// RespCodeEnum 内置错误码（暂未启用，后续可能会考虑启用）
type RespCodeEnum int

// TODO
//
//	尽量使用标准化的 HTTP HEAD 状态码，来规范大类的错误；如果他是业务自定义的错误，我们再使用 JSON Code，他只是一种业务提示（Error Code Hint）
//	请求过于频繁可以使用 http.StatusTooManyRequests
const (
	CodeSuccess            RespCodeEnum = 0    // 200 http.StatusOK 接口成功响应数据
	CodeFailBadRequest     RespCodeEnum = 1001 // 400 http.StatusBadRequest 参数 基础校验、业务校验不通过
	CodeFailUnauthorized   RespCodeEnum = 2001 // 401 http.StatusUnauthorized 要求授权
	CodeFailExpected       RespCodeEnum = 3001 // 500 http.StatusInternalServerError 业务模块因业务数据条件报出的错误 TODO 前端根据具体的 code 做具体的提示
	CodeFailUnexpectedData RespCodeEnum = 4001 // 500 http.StatusInternalServerError 查询数据出错（数据库中的数据不是预期的状态） TODO 结合 dao 模块（不好搞）
	CodeFailPanic          RespCodeEnum = 9999 // 500 http.StatusInternalServerError 顶级错误 TODO 自定义 panic handler，gin 有考虑这个，可以直接基于 gin 的进行拓展
)
