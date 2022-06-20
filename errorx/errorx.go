package errorx

import (
	"fmt"
)

type ErrorX struct {
	BizNo   int
	BizMsg  string
	Details []string
}

func (e *ErrorX) Error() string {
	return fmt.Sprintf("err_no: %d, err_msg: %s, details: %v", e.BizNo, e.BizMsg, e.Details)
}

func New(errNo int, errMsg string, details ...string) *ErrorX {
	return &ErrorX{BizNo: errNo, BizMsg: errMsg, Details: details}
}

func (e *ErrorX) WithDetails(details ...string) *ErrorX {
	e.Details = append(e.Details, details...)
	return e
}

const (
	NoParamInvalid   = 1  //参数错误
	NoSystemError    = 2  //服务内部错误
	NoUserNotLogin   = 3  //用户未登录
	NoInvalidRequest = 14 //无效请求
)

var ErrMSG = map[int]string{
	NoParamInvalid:   "请求参数不合理",
	NoSystemError:    "服务异常， 请稍后重试",
	NoUserNotLogin:   "用户未登录，请登录后重试",
	NoInvalidRequest: "请求无效，请稍后再试",
}

var (
	ParamInvalid = New(NoParamInvalid, ErrMSG[NoParamInvalid])
	SystemError  = New(NoSystemError, ErrMSG[NoSystemError])
)
