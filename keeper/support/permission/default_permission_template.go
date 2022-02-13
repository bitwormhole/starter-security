package permission

import (
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
	"github.com/bitwormhole/starter/markup"
)

// DefaultPermissionTemplateFactory ...
type DefaultPermissionTemplateFactory struct {
	markup.Component `id:"keeper-default-permission-template-factory"`
}

func (inst *DefaultPermissionTemplateFactory) _Impl() keeper.PermissionTemplateFactory {
	return inst
}

// CreateTemplate ...
func (inst *DefaultPermissionTemplateFactory) CreateTemplate(spr *keeper.SimplePermissionRegistration) (keeper.PermissionTemplate, error) {
	t := &facadePermissionTemplate{spr: spr}
	return t, nil
}

////////////////////////////////////////////////////////////////////////////////

// facadePermissionTemplate
type facadePermissionTemplate struct {
	spr *keeper.SimplePermissionRegistration
}

func (inst *facadePermissionTemplate) _Impl() keeper.PermissionTemplate {
	return inst
}

func (inst *facadePermissionTemplate) Method() string {
	return inst.spr.Method
}

func (inst *facadePermissionTemplate) PathPattern() string {
	return inst.spr.PathPattern
}

func (inst *facadePermissionTemplate) AcceptRole(role users.Role) bool {
	accepts := inst.spr.Roles
	for _, a := range accepts {
		if a.Equal(role) {
			return true
		}
	}
	return false
}

func (inst *facadePermissionTemplate) AcceptRoles(roles users.Roles) bool {

	mid := make(map[string]bool)
	list2 := roles.List()
	list1 := inst.spr.Roles

	for _, item := range list1 {
		name := item.String()
		if name == "" {
			continue
		}
		mid[name] = true
	}

	for _, item := range list2 {
		ok := mid[item.String()]
		if ok {
			return true
		}
	}

	return false
}

func (inst *facadePermissionTemplate) LoadPermission(params map[string]string) (keeper.Permission, error) {
	loader := inst.spr.Loader
	return loader.Load(inst, params)
}
