package common

import (
	"github.com/bitwormhole/starter-security/auths"
	"github.com/bitwormhole/starter/lang"
)

type AuthToken struct {
	UserName string
	Secret   string
	AuthType string
}

func (inst *AuthToken) _impl_() auths.AuthenticationToken {
	return inst
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
