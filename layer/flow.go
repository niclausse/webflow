package layer

import (
	"context"
	"github.com/niclausse/webflow/utils"
)

type IFlowParam interface {
	Validate() error
}

type FlowParam struct {
}

func (p *FlowParam) Validate() error {
	return nil
}

type IFlow interface {
	SetContext(ctx context.Context)
	GetContext() context.Context
	OnCreate(param IFlowParam)
	Create(newFlow IFlow) IFlow
	CreateWithParam(newFlow IFlow, param IFlowParam) IFlow
	Copy(ctx context.Context) IFlow
}

type Flow struct {
	ctx   context.Context
	Param IFlowParam
}

func (f *Flow) SetContext(ctx context.Context) {
	f.ctx = ctx
}

func (f *Flow) GetContext() context.Context {
	return f.ctx
}

func (f *Flow) OnCreate(param IFlowParam) {
	f.Param = param
}

func (f *Flow) Create(newFlow IFlow) IFlow {
	return f.CreateWithParam(newFlow, nil)
}

func (f *Flow) CreateWithParam(newFlow IFlow, param IFlowParam) IFlow {
	newFlow.SetContext(f.ctx)
	newFlow.OnCreate(param)
	return newFlow
}

func (f *Flow) Copy(ctx context.Context) IFlow {
	newFlow := f.CreateWithParam(utils.NewObject(f).(IFlow), f.Param)
	if ctx != nil {
		newFlow.SetContext(ctx)
	}
	return newFlow
}
