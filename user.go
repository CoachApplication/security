package security

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
)

const (
	OPERATION_ID_USER = "security.user"
)

type User interface {
	Id() string
	Name() string
}

type UserOperation struct{}

func (up *UserOperation) Id() string {
	return OPERATION_ID_USER
}

func (up *UserOperation) Ui() api.Ui {
	return base.NewUi(
		up.Id(),
		"User",
		"Active user",
		"",
	).Ui()
}

func (up *UserOperation) Usage() api.Usage {
	return (&base.ExternalOperationUsage{}).Usage()
}

func (up *UserOperation) Properties() api.Properties {
	return base.NewProperties().Properties()
}
