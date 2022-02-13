package permission

import (
	"context"
	"errors"
	"sync"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// DefaultPermissionManager ...
type DefaultPermissionManager struct {
	markup.Component `id:"keeper-permission-manager" class:"keeper-configurer"`

	AppCtx                            application.Context         `inject:"context"`
	PermissionRegistryList            []keeper.PermissionRegistry `inject:".keeper-permission-registry"`
	PermissionsPropertiesResourceName string                      `inject:"${security.keeper.permissions.properties}"`

	cache map[string]*keeper.SimplePermissionRegistration
	mutex sync.RWMutex
}

func (inst *DefaultPermissionManager) _Impl() (keeper.PermissionManager, keeper.Configurer) {
	return inst, inst
}

// Configure ...
func (inst *DefaultPermissionManager) Configure(c *keeper.Context) error {
	err := inst.loadPermissions()
	if err != nil {
		return err
	}
	c.Permissions = inst
	return nil
}

func (inst *DefaultPermissionManager) loadPermissions() error {
	ctx := inst.AppCtx
	name := inst.PermissionsPropertiesResourceName
	loader := permissionTemplateTableBuilder{context: inst.AppCtx}
	err := loader.LoadFromRegistryList(inst.PermissionRegistryList)
	if err != nil {
		return err
	}
	err = loader.LoadFromResource(ctx, name)
	if err != nil {
		return err
	}
	table, err := loader.Create()
	if err != nil {
		return err
	}
	inst.cache = table
	return nil
}

// FindTemplate ...
func (inst *DefaultPermissionManager) FindTemplate(ctx context.Context, a keeper.Access) (keeper.PermissionTemplate, error) {

	inst.mutex.RLock()
	defer inst.mutex.RUnlock()

	table := inst.cache
	akey := inst.getAccessKey(a)
	holder := table[akey]
	if holder == nil {
		return nil, errors.New("no permission for " + akey)
	}

	template := holder.Template
	if template == nil {
		return nil, errors.New("no permission for " + akey)
	}

	return template, nil
}

func (inst *DefaultPermissionManager) getAccessKey(a keeper.Access) string {
	method := a.Method()
	path := a.PathPattern()
	return makeAccessKey(method, path)
}
