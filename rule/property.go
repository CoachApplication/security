package rule

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	"github.com/CoachApplication/base/errors"
)

const (
	PROPERTY_ID_RULERESULT = "ssecurity.rule.result"
)

type ResultProperty struct {
	res Result
}

func (rrp *ResultProperty) Id() string {
	return PROPERTY_ID_RULERESULT
}

func (rrp *ResultProperty) Ui() api.Ui {
	return base.NewUi(
		rrp.Id(),
		"Rule Result",
		"The result of a security check",
		"",
	).Ui()
}

func (rrp *ResultProperty) Usage() api.Usage {
	return (&base.ReadonlyPropertyUsage{}).Usage()
}

func (rrp *ResultProperty) Validate() bool {
	return rrp.res != nil
}

func (rrp *ResultProperty) Get() interface{} {
	return interface{}(rrp.res)
}

func (rrp *ResultProperty) Set(val interface{}) error {
	if conv, good := val.(Result); good {
		rrp.res = conv
		return nil
	} else {
		return errors.PropertyWrongValueTypeError{Id: rrp.Id()}
	}
}
