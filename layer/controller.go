package layer

import (
	"github.com/gin-gonic/gin/binding"
)

type IController interface {
	GetBindingType() binding.Binding
	GetRequest() interface{}
	Action() (interface{}, error)
}

type defaultCTL struct {
	req *struct{}
}

func (d *defaultCTL) GetBindingType() binding.Binding {
	return binding.JSON
}

func (d *defaultCTL) GetRequest() interface{} {
	return struct{}{}
}

func (d *defaultCTL) Action() (interface{}, error) {
	return "implement me", nil
}
