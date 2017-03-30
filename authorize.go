package security

import (
	"github.com/CoachApplication/api"
)

type AuthorizeConnector interface {
	AuthorizeOperation(operation api.Operation) AuthorizeResult
}

type AuthorizeResult interface {
	Message() string
	Allowed() bool
	Denied() bool
}
