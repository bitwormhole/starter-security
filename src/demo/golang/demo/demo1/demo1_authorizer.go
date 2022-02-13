package demo1

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
	"github.com/bitwormhole/starter/markup"
)

type Demo1authorizer struct {
	markup.Component `class:"keeper-authorizer-registry"`
}

func (inst *Demo1authorizer) _Impl() (keeper.Authorizer, keeper.AuthorizerRegistry) {
	return inst, inst
}

func (inst *Demo1authorizer) GetRegistrationList() []*keeper.AuthorizerRegistration {
	ar := &keeper.AuthorizerRegistration{}
	ar.Authorizer = inst
	ar.Enabled = true
	ar.Name = "demo"
	return []*keeper.AuthorizerRegistration{ar}
}

func (inst *Demo1authorizer) Authorize(ctx context.Context, a keeper.Authorization) error {

	perm := a.GetPermission()
	roles := a.GetRoles()
	session, err := a.GetSubject().GetSession(true)
	if err != nil {
		return err
	}

	ident := session.GetIdentity()
	rolelist1 := ident.Roles().List()
	for _, role := range rolelist1 {
		roles = roles.Add(role)
	}

	if perm.IsOwner(ident) {
		roles = roles.Add(users.RoleOwner)
	} else if perm.IsFriend(ident) {
		roles = roles.Add(users.RoleFriend)
	}

	a.SetRoles(roles)
	return nil
}
