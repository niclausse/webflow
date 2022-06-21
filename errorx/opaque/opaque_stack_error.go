package opaque

import (
	"fmt"
	"io"
)

type stackBizError struct {
	bizNo  int
	bizMsg string
	*stack
}

func (sbe *stackBizError) Error() string {
	return fmt.Sprintf("err_no: %d, err_msg: %s", sbe.bizNo, sbe.bizMsg)
}

func StackError(bizNo int, bizMsg string) error {
	return &stackBizError{
		bizNo:  bizNo,
		bizMsg: bizMsg,
		stack:  callers(),
	}
}

func (sbe *stackBizError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, sbe.Error())
			sbe.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, sbe.Error())
	case 'q':
		fmt.Fprintf(s, "%q", sbe.Error())
	}
}
