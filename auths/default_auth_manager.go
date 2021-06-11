package auths

import (
	"errors"
	"fmt"
)

type DefaultAuthenticationManager struct {
	// context      application.Context
	ProviderList []AuthenticationProvider
}

func (inst *DefaultAuthenticationManager) _impl_() AuthenticationManager {
	return inst
}

func (inst *DefaultAuthenticationManager) Authenticate(token AuthenticationToken) (AuthenticationInfo, error) {
	list := inst.ProviderList
	if list == nil {
		return nil, errors.New("DefaultAuthenticationManager.ProviderList == nil , type: []AuthenticationProvider")
	}
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
	src := inst.ProviderList
	if src == nil {
		return dst
	}
	for index := range src {
		item := src[index]
		dst = append(dst, item)
	}
	return dst
}
