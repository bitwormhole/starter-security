package authorization

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
)

// DefaultAuthorizationManager ...
type DefaultAuthorizationManager struct {
	markup.Component `id:"keeper-authorization-manager" class:"keeper-configurer"`

	AuthorizerRegistryList []keeper.AuthorizerRegistry `inject:".keeper-authorizer-registry"`

	cache []keeper.Authorizer
}

func (inst *DefaultAuthorizationManager) _Impl() (keeper.AuthorizationManager, keeper.Configurer) {
	return inst, inst
}

// Configure ...
func (inst *DefaultAuthorizationManager) Configure(c *keeper.Context) error {
	c.Authorizations = inst
	return inst.loadAuthorizers(c)
}

func (inst *DefaultAuthorizationManager) loadAuthorizers(c *keeper.Context) error {
	src := inst.AuthorizerRegistryList
	dst := make([]keeper.Authorizer, 0)
	for _, item1 := range src {
		mid := item1.GetRegistrationList()
		for _, item2 := range mid {
			if !item2.Enabled {
				continue
			}
			if item2.Authorizer == nil {
				continue
			}
			dst = append(dst, item2.Authorizer)
		}
	}
	inst.cache = dst
	return nil
}

// Authorize ...
func (inst *DefaultAuthorizationManager) Authorize(ctx context.Context) error {

	h, err := keeper.GetHolder(ctx)
	if err != nil {
		return err
	}

	err = inst.preparePermission(h)
	if err != nil {
		return err
	}

	ac := h.GetAccessContext()
	sa := ac.SecurityAccess
	all := inst.cache

	for _, item := range all {
		err := item.Authorize(ctx, sa)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *DefaultAuthorizationManager) preparePermission(h *keeper.Holder) error {

	ac := h.GetAccessContext()
	ctx := ac.Context
	access := ac.SecurityAccess
	pm := ac.SecurityContext.GetPermissions()
	template, err := pm.FindTemplate(ctx, access)
	if err != nil {
		return err
	}

	params := access.Params()
	perm, err := template.LoadPermission(params)
	if err != nil {
		return err
	}

	ac.Permission = perm
	return nil
}

// ListAuthorizers ...
func (inst *DefaultAuthorizationManager) ListAuthorizers() []keeper.Authorizer {
	src := inst.cache
	dst := make([]keeper.Authorizer, len(src))
	copy(dst, src)
	return dst
}

// // ListAuthorizerRegistrations ...
// func (inst *DefaultAuthorizationManager) ListAuthorizerRegistrations() []*keeper.AuthorizerRegistration {
// 	return []*keeper.AuthorizerRegistration{}
// }
