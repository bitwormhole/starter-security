package authorization

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
)

type DefaultAuthorizationManager struct {
	markup.Component `id:"keeper-authorization-manager" class:"keeper-configurer"`
}

func (inst *DefaultAuthorizationManager) _Impl() (keeper.AuthorizationManager, keeper.Configurer) {
	return inst, inst
}

func (inst *DefaultAuthorizationManager) Configure(c *keeper.Context) error {
	c.Authorizations = inst
	return nil
}

func (inst *DefaultAuthorizationManager) Accept(ctx context.Context, a keeper.Authorization) bool {
	return false
}
