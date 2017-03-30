package security

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	"errors"
)

type AuthorizedDecorationFactory struct {
	authOp              api.Operation
}

func (adf *AuthorizedDecorationFactory) DecorateOperation(api.Operation) api.Operation {

}

type AuthorizingDecoratorOperation struct {
	tOp api.Operation
	con AuthorizeConnector
	authorizeOnValidate bool
}

func (ado *AuthorizingDecoratorOperation) Operation() api.Operation {
	return api.Operation(ado)
}

func (ado *AuthorizingDecoratorOperation) Id() string {
	return ado.tOp.Id()
}

func (ado *AuthorizingDecoratorOperation) Ui() api.Ui {
	return ado.tOp.Ui()
}

func (ado *AuthorizingDecoratorOperation) Usage() api.Usage {
	return ado.tOp.Usage()
}

func (ado *AuthorizingDecoratorOperation) Properties() api.Properties {
	return ado.tOp.Properties()
}

func (ado *AuthorizingDecoratorOperation) Validate(props api.Properties) api.Result {
	mainRes := base.NewResult()

	valRes := base.NewResult()
	go func () {
		if ado.con == nil {
			valRes.AddError(errors.New("Missing authorization connector"))
			valRes.MarkFinished()
		} else if ado.authorizeOnValidate {
			ado.con.AuthorizeOperation(ado.tOp)
		}
	}()

	mainRes.Merge(valRes.Result())
	mainRes.Merge(ado.tOp.Validate(props))
	return mainRes.Result()
}
