package session

import (
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
)

// MainSessionProvider ...
type MainSessionProvider struct {
	markup.Component `id:"keeper-main-session-provider"  class:"keeper-configurer"  initMethod:"Init"`

	Selector  string                           `inject:"${keeper.session-provider.name}"` // the name of regist
	Providers []keeper.SessionProviderRegistry `inject:".keeper-session-provider-registry"`

	provider keeper.SessionProvider
}

func (inst *MainSessionProvider) _Impl() (keeper.SessionProvider, keeper.Configurer) {
	return inst, inst
}

// Configure ...
func (inst *MainSessionProvider) Configure(c *keeper.Context) error {
	c.SessionProvider = inst
	return nil
}

// Init ...
func (inst *MainSessionProvider) Init() error {
	inst.getTargetProvider()
	return nil
}

func (inst *MainSessionProvider) getProviderTable() map[string]keeper.SessionProvider {
	table := make(map[string]keeper.SessionProvider)
	src := inst.Providers
	for _, item := range src {
		list2 := item.GetRegistrationList()
		for _, item2 := range list2 {
			name := item2.Name
			p := item2.Provider
			old := table[name]
			if p == nil {
				continue
			}
			if old != nil {
				panic("the name of keeper seesion provider is override: " + name)
			}
			table[name] = p
		}
	}
	return table
}

func (inst *MainSessionProvider) getTargetProvider() keeper.SessionProvider {
	provider := inst.provider
	if provider != nil {
		return provider
	}
	name := inst.Selector
	table := inst.getProviderTable()
	provider = table[name]
	if provider == nil {
		panic("no keeper session provider with name:" + name)
	}
	inst.provider = provider
	return provider
}

// GetSessionFactory ...
func (inst *MainSessionProvider) GetSessionFactory() keeper.SessionFactory {
	return inst.getTargetProvider().GetSessionFactory()
}

// GetAdapterFactory ...
func (inst *MainSessionProvider) GetAdapterFactory() keeper.SessionAdapterFactory {
	return inst.getTargetProvider().GetAdapterFactory()
}
