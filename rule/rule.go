package rule

import (
	"github.com/CoachApplication/api"
)

// Rule an object which can test authorization
type Rule interface {
	Id() string
	AuthorizeOperation(operation api.Operation) Result
}

// Result The result of applying a rule
type Result interface {
	// Return the Id of the Rule that returned the result
	Rule() string
	Allowed() bool
	Denied() bool
	Message() ResultMessage
}

// ResultMessage Abstracted message for when a rule has applied
type ResultMessage interface {
	Message() string
}
