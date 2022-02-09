package session

import (
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
)

// DefaultSessionProvider ...
type DefaultSessionProvider struct {
	markup.Component `class:"keeper-session-provider-registry" initMethod:"Init"`

	adapterFactory keeper.SessionAdapterFactory
	sessionFactory keeper.SessionFactory
}

func (inst *DefaultSessionProvider) _Impl() (keeper.SessionProvider, keeper.SessionProviderRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *DefaultSessionProvider) GetRegistrationList() []*keeper.SessionProviderRegistration {
	r := &keeper.SessionProviderRegistration{}
	r.Name = "default"
	r.Provider = inst
	return []*keeper.SessionProviderRegistration{r}
}

// Init ...
func (inst *DefaultSessionProvider) Init() error {
	inst.adapterFactory = &DefaultSessionAdapterFactory{}
	inst.sessionFactory = &DefaultSessionFactory{}
	return nil
}

// GetSessionFactory ...
func (inst *DefaultSessionProvider) GetSessionFactory() keeper.SessionFactory {
	return inst.sessionFactory
}

// GetAdapterFactory ...
func (inst *DefaultSessionProvider) GetAdapterFactory() keeper.SessionAdapterFactory {
	return inst.adapterFactory
}
