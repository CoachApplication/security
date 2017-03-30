package security

import (
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	base_errors "github.com/CoachApplication/base/errors"
)

const (
	PROPERTY_ID_USER = "security.user"
)

type UserProperty struct {
	user User
}

func NewUserProperty(us User) *UserProperty {
	return &UserProperty{
		user: us,
	}
}

func (up *UserProperty) Property() api.Property {
	return api.Property(up)
}

func (up *UserProperty) Id() string {
	return PROPERTY_ID_USER
}

func (up *UserProperty) Type() string {
	return "coach/security/user"
}

func (up *UserProperty) Ui() api.Ui {
	return base.NewUi(
		up.Id(),
		"User",
		"User",
		"",
	).Ui()
}

func (up *UserProperty) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}
}

func (up *UserProperty) Validate() bool {
	return up.user != nil
}

func (up *UserProperty) Get() interface{} {
	return interface{}(up.user)
}

func (up *UserProperty) Set(val interface{}) error {
	if conv, good := val.(User); good {
		up.user = conv
		return nil
	} else {
		return error(base_errors.PropertyWrongValueTypeError{
			Id:           up.Id(),
			ExpectedType: up.Type(),
		})
	}
}
