package common

import (
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/security"
)

type AuthToken struct {
	UserName string
	Secret   string
	AuthType string
}

func (inst *AuthToken) _Impl() security.AuthenticationToken {
	return nil
}

func (inst *AuthToken) Principal() lang.Object {
	return inst.UserName
}

func (inst *AuthToken) Credentials() lang.Object {
	return inst.Secret
}

func (inst *AuthToken) Mechanism() string {
	return inst.AuthType
}
