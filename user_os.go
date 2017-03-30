package security

import (
	"os/user"
)

// OsUser create a User implementation based on the os/user User interface
type OsUser struct {
	user user.User
}

// NewOsUser Constructor for OsUser
func NewOsUser(u user.User) *OsUser {
	return &OsUser{
		user: u,
	}
}

func (ou *OsUser) User() User {
	return User(ou)
}

func (ou *OsUser) Id() string {
	return ou.user.Username
}

func (ou *OsUser) Name() string {
	return ou.user.Name
}
