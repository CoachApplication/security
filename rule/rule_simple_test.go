package rule_test

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/security"
	"testing"
)

type TestRule struct {
	id  string
	res security.Result
}

func NewTestRule(id string, res security.Result) *TestRule {
	return &TestRule{
		id:  id,
		res: res,
	}
}

func (tr *TestRule) Id() string {
	return tr.id
}
func (tr *TestRule) AuthorizeOperation(operation api.Operation) security.Result {
	return tr.res
}


func TestTestRule_Id(t *testing.T) {
	r := NewTestRule("test", 1, security.NewSimpleResult("test", security.NewResultMessage("Test Message", nil)))
}