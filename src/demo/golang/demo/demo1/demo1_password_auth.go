package demo1

import (
	"context"
	"errors"
	"strings"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
	"github.com/bitwormhole/starter/markup"
)

// Demo1passwordAuth ...
type Demo1passwordAuth struct {
	markup.Component `class:"keeper-authenticator-registry"`
}

func (inst *Demo1passwordAuth) _Impl() (keeper.Authenticator, keeper.AuthenticatorRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *Demo1passwordAuth) GetRegistrationList() []*keeper.AuthenticatorRegistration {

	ar := &keeper.AuthenticatorRegistration{}
	ar.Name = "demo-password-auth"
	ar.Mechanism = "PASSWORD"
	ar.Authenticator = inst

	return []*keeper.AuthenticatorRegistration{ar}
}

// Supports ...
func (inst *Demo1passwordAuth) Supports(ctx context.Context, a keeper.Authentication) bool {
	mech := a.Mechanism()
	return strings.ToLower(mech) == "password"
}

// Verify ...
func (inst *Demo1passwordAuth) Verify(ctx context.Context, a keeper.Authentication) (keeper.Identity, error) {

	user := a.User()
	pass := string(a.Secret())

	ib := keeper.IdentityBuilder{}
	ib.UserID = users.UserID(0)
	ib.UserName = users.UserName(user)
	ib.UserUUID = users.UserUUID(user)
	ib.Roles = users.Roles(users.RoleAdmin)
	ident := ib.Identity()

	if user == theUserName && pass == theUserPass {
		return ident, nil
	}

	return nil, errors.New("bad pass+user")
}

////////////////////////////////////////////////////////////////////////////////

// type Demo1passwordIdent struct {
// 	user string
// }

// func (inst *Demo1passwordIdent) _Impl() keeper.Identity {
// 	return inst
// }

// func (inst *Demo1passwordIdent) UserID() string {
// 	return inst.user
// }

// func (inst *Demo1passwordIdent) UserName() users.Name {
// 	return inst.user
// }

// func (inst *Demo1passwordIdent) UserUUID() string {
// 	return "000000000000000"
// }

// func (inst *Demo1passwordIdent) Nickname() string {
// 	return inst.user
// }

// func (inst *Demo1passwordIdent) Avatar() string {
// 	return ""
// }

// func (inst *Demo1passwordIdent) Roles() users.Roles {
// 	return "normal"
// }

////////////////////////////////////////////////////////////////////////////////
