package rule

import "fmt"

/**
 * Authorization Rules
 */

// Simple integer based Result
type SimpleResult struct {
	id      string
	status  int
	message ResultMessage
}

// Constructor for SimpleResult
func NewSimpleResult(id string, status int, message ResultMessage) *SimpleResult {
	if message == nil {
		message = *SimpleResultMessageFromStatus(status)
	}

	return &SimpleResult{
		id:      id,
		message: message,
		status:  status,
	}
}

// Convert this into a Result interface
func (result *SimpleResult) Result() Result {
	return Result(result)
}

// Does the result explicitly pass
func (result *SimpleResult) Rule() string {
	return result.id
}

// Does the result explicitly pass
func (result *SimpleResult) Message() ResultMessage {
	return result.message
}

// Does the result explicitly pass
func (result *SimpleResult) Allow() bool {
	return result.status > 0
}

// Does the result explicitly fail
func (result *SimpleResult) Deny() bool {
	return result.status < 0
}

// Use this for rule messages as we may need to translate in the future
type SimpleResultMessage struct {
	// Think of these as fmt.Sprintf() arguments
	message    string
	parameters []interface{}
}

func NewResultMessage(m string, p []interface{}) *SimpleResultMessage {
	return &SimpleResultMessage{
		message:    m,
		parameters: p,
	}
}

func SimpleResultMessageFromStatus(s int) *SimpleResultMessage {
	if s > 0 {
		return SimpleResultMessageFromString("Authorized")
	} else if s < 0 {
		return SimpleResultMessageFromString("Denied")
	} else {
		return SimpleResultMessageFromString("Ignored")
	}
}
func SimpleResultMessageFromString(m string) *SimpleResultMessage {
	return &SimpleResultMessage{
		message: m,
	}
}

func (srm *SimpleResultMessage) Message() string {
	return fmt.Sprintf(srm.message, srm.parameters...)
}
