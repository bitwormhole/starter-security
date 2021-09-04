// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etc

// <import>

import (
	errors "errors"

	auths_85a1da "github.com/bitwormhole/starter-security/auths"
	common_422e0c "github.com/bitwormhole/starter-security/auths/common"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

// </import>

func Config(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

	// authByPassword
	cominfobuilder.Reset()
	cominfobuilder.ID("authByPassword").Class("auth-provider").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &common_422e0c.SimpleAuthWrapper{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &authByPassword{}
		adapter.instance = o.(*common_422e0c.SimpleAuthWrapper)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
	if err != nil {
		return err
	}

	// authService
	cominfobuilder.Reset()
	cominfobuilder.ID("auths").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &auths_85a1da.DefaultAuthenticationManager{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &authService{}
		adapter.instance = o.(*auths_85a1da.DefaultAuthenticationManager)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// type authByPassword struct

func (inst *authByPassword) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters

	// to instance

	// invoke custom inject method
	err = inst.inject(injection)
	if err != nil {
		return err
	}

	return injection.Close()
}

////////////////////////////////////////////////////////////////////////////////
// type authService struct

func (inst *authService) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.ProviderList = inst.__get_ProviderList__(injection, ".auth-provider")

	// to instance
	instance.ProviderList = inst.ProviderList

	// invoke custom inject method

	return injection.Close()
}

func (inst *authService) __get_ProviderList__(injection application.Injection, selector string) []auths_85a1da.AuthenticationProvider {
	list := make([]auths_85a1da.AuthenticationProvider, 0)
	reader := injection.Select(selector)
	defer reader.Close()
	for reader.HasMore() {
		o1, err := reader.Read()
		if err != nil {
			injection.OnError(err)
			return list
		}
		o2, ok := o1.(auths_85a1da.AuthenticationProvider)
		if !ok {
			err = errors.New("bad cast, selector:" + selector)
			injection.OnError(err)
			return list
		}
		list = append(list, o2)
	}
	return list

}
