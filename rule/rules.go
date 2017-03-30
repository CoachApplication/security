package rule

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/utils"
)

/**
 * The Rules struct is provided to simplify processing a set of rules.
 */

type OrderedRules struct {
	m utils.OrderedMap
}

func (or *OrderedRules) safe() {
	if or.m == nil {
		or.m = utils.NewOrderedMap()
	}
}

func (or *OrderedRules) Set(id string, rule Rule) error {
	or.safe()
	return or.m.Set(id, interface{}(rule))
}

func (or *OrderedRules) Get(id string) (Rule, error) {
	or.safe()
	r, err := or.m.Get(id)
	return r.(Rule), err
}

func (or *OrderedRules) Order() []string {
	or.safe()
	return or.m.Order()
}

func (or *OrderedRules) AuthorizeOperation(op api.Operation) Result {
	for _, id := range or.Order() {
		rule, _ := or.Get(id)

		res := rule.AuthorizeOperation(op)
		if res.Allowed() || res.Denied() {
			return res
		}
	}
	return (&NewSimpleResult(
		"security.rule.none",
		0,
		NewResultMessage("No rules applied.", nil),
	)).Result()
}
