package configwrapper

import (
	config "github.com/CoachApplication/config"
)

const (
	CONFIG_ID_SECURITY_USER = "user"
)

type OperationBase struct {
	wr config.Wrapper
}

func NewOperationBase(wr config.Wrapper) *OperationBase {
	return &OperationBase{
		wr: wr,
	}
}

func (ob *OperationBase) wrapper() config.Wrapper {
	return ob.wr
}
