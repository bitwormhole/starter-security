package permission

import (
	"strings"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
)

// DefaultPermissionLoaderFactory ...
type DefaultPermissionLoaderFactory struct {
	markup.Component `id:"keeper-default-permission-loader-factory"`
}

func (inst *DefaultPermissionLoaderFactory) _Impl() keeper.PermissionLoaderFactory {
	return inst
}

// CreateLoader ...
func (inst *DefaultPermissionLoaderFactory) CreateLoader(spr *keeper.SimplePermissionRegistration) (keeper.PermissionLoader, error) {
	l := &facadePermissionLoader{spr: spr}
	return l, nil
}

////////////////////////////////////////////////////////////////////////////////

type facadePermissionLoader struct {
	spr *keeper.SimplePermissionRegistration
}

func (inst *facadePermissionLoader) _Impl() keeper.PermissionLoader {
	return inst
}

func (inst *facadePermissionLoader) Load(template keeper.PermissionTemplate, params map[string]string) (keeper.Permission, error) {
	path, err := inst.computeFullPath(template, params)
	if err != nil {
		return nil, err
	}
	perm := &defaultPermission{}
	perm.template = template
	perm.path = path
	return perm, nil
}

func (inst *facadePermissionLoader) computeFullPath(template keeper.PermissionTemplate, params map[string]string) (string, error) {
	params2 := make(map[string]string)
	for k, v := range params {
		params2[":"+k] = v
	}
	builder := strings.Builder{}
	pattern := template.PathPattern()
	parts := strings.Split(pattern, "/")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		value := params2[part]
		if value == "" {
			value = part
		}
		builder.WriteRune('/')
		builder.WriteString(value)
	}
	return builder.String(), nil
}

////////////////////////////////////////////////////////////////////////////////
