package logic

import (
	"context"
	"fmt"

	"github.com/penglin1995/webflow/examples/dto"
)

type UserLogic struct {
}

func NewUserLogic() *UserLogic {
	return &UserLogic{}
}

func (l *UserLogic) Add(ctx context.Context, req *dto.AddUserReq) error {
	fmt.Printf("%+v\n", req)
	return nil
}
