package configwrapper

import "github.com/CoachApplication/security"

type ConfigUser struct {
	FromId   string `json:"id" yaml:"Id"`
	FromName string `json:"name" yaml:"Name"`
}

func (cu *ConfigUser) User() security.User {
	return security.User(cu)
}

func (cu *ConfigUser) Id() string {
	return cu.FromId
}

func (cu *ConfigUser) Name() string {
	return cu.FromName
}
