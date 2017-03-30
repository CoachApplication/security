package configwrapper

import (
	"context"
	"errors"
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	base_errors "github.com/CoachApplication/base/errors"
	base_property "github.com/CoachApplication/base/property"
	config "github.com/CoachApplication/config"
	security "github.com/CoachApplication/security"
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

type UserOperation struct {
	OperationBase
	security.UserOperation
}

func (uo *UserOperation) Operation() api.Operation {
	return api.Operation(uo)
}

func (uo *UserOperation) Validate(props api.Properties) api.Result {
	if uo.wr == nil {
		return base.MakeFailedResult()
	} else {
		return base.MakeSuccessfulResult()
	}
}

func (uo *UserOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	ctx := context.Background()
	if ctxProp, err := props.Get(base_property.PROPERTY_ID_CONTEXTLIMIT); err == nil {
		ctx = ctxProp.Get().(context.Context)
	}

	go func() {
		if usc, err := uo.wr.Get(CONFIG_ID_SECURITY_USER); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else if uc, err := usc.Get(config.CONFIG_SCOPE_DEFAULT); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {
			var cu ConfigUser
			getRes := uc.Get(&cu)

			select {
			case <-getRes.Finished():
				if getRes.Success() {
					p := security.NewUserProperty(cu.User())
					res.AddProperty(p.Property())
					res.MarkSucceeded()
				} else {
					res.AddError(errors.New("Failed to revtrieve User from config wrapper"))
					res.MarkFailed()
				}
			case <-ctx.Done():
				res.AddError(error(base_errors.OperationTimedOut{Ctx: ctx}))
				res.MarkFailed()
			}
		}

		res.MarkFinished()
	}()

	return res.Result()
}
