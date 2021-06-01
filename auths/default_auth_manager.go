package auths

import (
	"errors"
	"fmt"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

type DefaultAuthenticationManager struct {
	context   application.Context
	providers []AuthenticationProvider
}

func (inst *DefaultAuthenticationManager) _impl_() AuthenticationManager {
	return inst
}

func (inst *DefaultAuthenticationManager) Inject(context application.Context) error {

	in := context.Injector()
	list := []AuthenticationProvider{}

	in.Inject("*").AsList().Accept(func(name string, holder application.ComponentHolder) bool {
		o := holder.GetPrototype()
		_, ok := o.(AuthenticationProvider)
		return ok
	}).To(func(o lang.Object) bool {
		item, ok := o.(AuthenticationProvider)
		if ok {
			list = append(list, item)
		}
		return ok
	})

	inst.providers = list
	inst.context = context
	return in.Done()
}

func (inst *DefaultAuthenticationManager) Authenticate(token AuthenticationToken) (AuthenticationInfo, error) {
	list := inst.providers
	if list != nil {
		for index := range list {
			item := list[index]
			if item == nil {
				continue
			}
			if item.Supports(token) {
				return item.Authenticate(token)
			}
		}
	}
	msg := fmt.Sprint("no auth provider support the token:", token)
	return nil, errors.New(msg)
}

func (inst *DefaultAuthenticationManager) Providers() []AuthenticationProvider {
	dst := []AuthenticationProvider{}
	src := inst.providers
	if src == nil {
		return dst
	}
	for index := range src {
		item := src[index]
		dst = append(dst, item)
	}
	return dst
}
