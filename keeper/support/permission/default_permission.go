package permission

import (
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
)

type defaultPermission struct {
	template keeper.PermissionTemplate
	path     string
}

func (inst *defaultPermission) _Impl() keeper.Permission {
	return inst
}

func (inst *defaultPermission) Method() string {
	return inst.template.Method()
}

func (inst *defaultPermission) Path() string {
	return inst.path
}

func (inst *defaultPermission) Owner() keeper.Identity {
	return nil
}

func (inst *defaultPermission) Friends() keeper.UserGroup {
	return nil
}

func (inst *defaultPermission) Template() keeper.PermissionTemplate {
	return inst.template
}

func (inst *defaultPermission) IsOwner(user keeper.Identity) bool {
	return false
}

func (inst *defaultPermission) IsFriend(user keeper.Identity) bool {
	return false
}

func (inst *defaultPermission) AcceptUser(user keeper.Identity) bool {
	return false
}

func (inst *defaultPermission) AcceptRole(role users.Role) bool {
	return inst.template.AcceptRole(role)
}

func (inst *defaultPermission) AcceptRoles(roles users.Roles) bool {
	return inst.template.AcceptRoles(roles)
}
