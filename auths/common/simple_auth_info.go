package common

import (
	"github.com/bitwormhole/starter-security/auths"
	"github.com/bitwormhole/starter/lang"
)

type AuthInfo struct {
	StatusOK     bool
	StatusText   string
	AccountID    string
	AccountUUID  string
	AccountRoles string
}

func (inst *AuthInfo) _impl_() auths.AuthenticationInfo {
	return inst
}

func (inst *AuthInfo) Principal() lang.Object {
	return inst.AccountID
}

func (inst *AuthInfo) Roles() string {
	return inst.AccountRoles
}

func (inst *AuthInfo) Success() bool {
	return inst.StatusOK
}

func (inst *AuthInfo) Message() string {
	return inst.StatusText
}
