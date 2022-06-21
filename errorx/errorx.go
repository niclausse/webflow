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
	x := &ErrorX{
		BizNo:   e.BizNo,
		BizMsg:  e.BizMsg,
		Details: e.Details,
	}

	x.Details = append(x.Details, details...)
	return x
}
