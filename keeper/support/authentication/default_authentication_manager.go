package authentication

import (
	"context"
	"errors"
	"strings"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// DefaultAuthenticationManager ...
type DefaultAuthenticationManager struct {
	markup.Component `id:"keeper-authentication-manager" class:"keeper-configurer"`

	RegistryList []keeper.AuthenticatorRegistry `inject:".keeper-authenticator-registry"`

	registrations []*keeper.AuthenticatorRegistration // cache for RegistryList
}

func (inst *DefaultAuthenticationManager) _Impl() (keeper.AuthenticationManager, keeper.Configurer) {
	return inst, inst
}

// Configure ...
func (inst *DefaultAuthenticationManager) Configure(c *keeper.Context) error {
	c.Authentications = inst
	return nil
}

func (inst *DefaultAuthenticationManager) normalizeMechanismName(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	return s
}

func (inst *DefaultAuthenticationManager) loadRegistrationList() []*keeper.AuthenticatorRegistration {
	src := inst.RegistryList
	dst := make([]*keeper.AuthenticatorRegistration, 0)
	table := make(map[string]keeper.Authenticator)
	for _, item1 := range src {
		mid := item1.GetRegistrationList()
		for _, item2 := range mid {
			target := item2.Authenticator
			if target == nil {
				continue
			}
			name := inst.normalizeMechanismName(item2.Mechanism)
			old := table[name]
			if old != nil {
				panic("the name of Authenticator-Mechanism is override: " + name)
			}
			table[name] = target
			dst = append(dst, item2)
		}
	}
	return dst
}

func (inst *DefaultAuthenticationManager) getRegistrationList() []*keeper.AuthenticatorRegistration {
	list := inst.registrations
	if list == nil {
		list = inst.loadRegistrationList()
		inst.registrations = list
	}
	return list
}

// Authenticate ...
func (inst *DefaultAuthenticationManager) Authenticate(ctx context.Context, a keeper.Authentication) (keeper.Identity, error) {
	all := inst.getRegistrationList()
	for _, item := range all {
		if inst.isSupported(ctx, a, item) {
			ident, err := item.Authenticator.Verify(ctx, a)
			if err == nil {
				return ident, nil
			}
			vlog.Warn(err)
		}
	}
	mechanism := inst.normalizeMechanismName(a.Mechanism())
	vlog.Error("bad auth, with mechanism:" + mechanism)
	return nil, errors.New("bad auth")
}

func (inst *DefaultAuthenticationManager) isSupported(ctx context.Context, a1 keeper.Authentication, ar *keeper.AuthenticatorRegistration) bool {
	if ar == nil {
		return false
	}
	a2 := ar.Authenticator
	if a1 == nil || a2 == nil {
		return false
	}
	m1 := inst.normalizeMechanismName(a1.Mechanism())
	m2 := inst.normalizeMechanismName(ar.Mechanism)
	if m1 != m2 {
		return false
	}
	return a2.Supports(ctx, a1)
}
