package errorx

import (
	"fmt"

	"github.com/pkg/errors"
)

type errorX struct {
	bizNo   int
	bizMsg  string
	details []string
}

func (e *errorX) Error() string {
	return fmt.Sprintf("err_no: %d, err_msg: %s", e.bizNo, e.bizMsg)
}

func New(errNo int, errMsg string, details ...string) error {
	return &errorX{bizNo: errNo, bizMsg: errMsg, details: details}
}

func AppendDetails(err error, details ...string) error {
	ex, ok := errors.Cause(err).(*errorX)
	if !ok {
		ex = &errorX{bizNo: NoSystemError, bizMsg: ErrMSG[NoSystemError]}
	}

	ex.details = append(ex.details, details...)
	return ex
}

func Biz(err error) (bizNo int, bizMsg string, details []string) {
	ex, ok := errors.Cause(err).(*errorX)
	if !ok {
		ex = &errorX{bizNo: NoSystemError, bizMsg: ErrMSG[NoSystemError], details: []string{fmt.Sprintf("[%s] is not an errorX", err.Error())}}
	}

	return ex.bizNo, ex.bizMsg, ex.details
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
