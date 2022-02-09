package authentication

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
)

type DefaultAuthenticationManager struct {
	markup.Component `id:"keeper-authentication-manager" class:"keeper-configurer"`
}

func (inst *DefaultAuthenticationManager) _Impl() (keeper.AuthenticationManager, keeper.Configurer) {
	return inst, inst
}

func (inst *DefaultAuthenticationManager) Configure(c *keeper.Context) error {
	c.Authentications = inst
	return nil
}

func (inst *DefaultAuthenticationManager) Authenticate(ctx context.Context, a keeper.Authentication) (keeper.Identity, error) {

	return nil, nil
}
