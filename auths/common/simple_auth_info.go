package common

import (
	"github.com/bitwormhole/starter/security"
)

type AuthInfo struct {
	StatusOK     bool
	StatusText   string
	AccountID    string
	AccountUUID  string
	AccountRoles string
}

func (inst *AuthInfo) _Impl() security.AuthenticationInfo {
	return nil
}
