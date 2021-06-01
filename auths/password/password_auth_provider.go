package password

import (
	"errors"

	"github.com/bitwormhole/starter-security/auths"
	"github.com/bitwormhole/starter-security/auths/common"
)

type AuthProvider struct {
	DAO       common.AuthDAO
	Mechanism string
}

func (inst *AuthProvider) _impl_() auths.AuthenticationProvider {
	return inst
}

func (inst *AuthProvider) Name() string {
	return "password.AuthProvider"
}

func (inst *AuthProvider) Supports(token auths.AuthenticationToken) bool {
	return true
}

func (inst *AuthProvider) Authenticate(token auths.AuthenticationToken) (auths.AuthenticationInfo, error) {

	user := token.Principal().(string)
	mechanism := inst.Mechanism

	// find record
	rec, err := inst.DAO.Find(user, mechanism)
	if err != nil {
		return nil, err
	}

	// check
	err = inst.check(token, rec)
	if err != nil {
		return nil, err
	}

	// make info
	accountId := rec.AccountID()
	accountUUID := rec.AccountUUID()
	info := &common.AuthInfo{
		StatusOK:     true,
		StatusText:   "OK",
		AccountID:    accountId,
		AccountRoles: accountUUID,
	}
	return info, nil
}

func (inst *AuthProvider) check(token auths.AuthenticationToken, rec common.Auth) error {

	salt := rec.Salt()
	sum1 := rec.Secret()

	computer := &AuthComputer{}
	computer.Init()
	computer.Salt = salt
	computer.User = token.Principal().(string)
	computer.Password = token.Credentials().(string)
	sum2 := computer.Compute()

	if sum1 != sum2 {
		return errors.New("Bad user or password.")
	}

	return nil
}
