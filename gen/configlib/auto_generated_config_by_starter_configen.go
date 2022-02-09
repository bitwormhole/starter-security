// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package configlib

import (
	keeper0x6d39ef "github.com/bitwormhole/starter-security/keeper"
	authentication0x38bbbc "github.com/bitwormhole/starter-security/keeper/support/authentication"
	authorization0xbd4cc5 "github.com/bitwormhole/starter-security/keeper/support/authorization"
	core0xb21891 "github.com/bitwormhole/starter-security/keeper/support/core"
	session0x14c51e "github.com/bitwormhole/starter-security/keeper/support/session"
	subject0xae5c7a "github.com/bitwormhole/starter-security/keeper/support/subject"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: keeper-authentication-manager
	cominfobuilder.Next()
	cominfobuilder.ID("keeper-authentication-manager").Class("keeper-configurer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultAuthenticationManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: keeper-authorization-manager
	cominfobuilder.Next()
	cominfobuilder.ID("keeper-authorization-manager").Class("keeper-configurer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultAuthorizationManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: keeper-core
	cominfobuilder.Next()
	cominfobuilder.ID("keeper-core").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComKeeperCore{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-session0x14c51e.DefaultSessionProvider
	cominfobuilder.Next()
	cominfobuilder.ID("com3-session0x14c51e.DefaultSessionProvider").Class("keeper-session-provider-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultSessionProvider{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: keeper-main-session-provider
	cominfobuilder.Next()
	cominfobuilder.ID("keeper-main-session-provider").Class("keeper-configurer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMainSessionProvider{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: keeper-subject-manager
	cominfobuilder.Next()
	cominfobuilder.ID("keeper-subject-manager").Class("keeper-configurer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultSubjectManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultAuthenticationManager : the factory of component: keeper-authentication-manager
type comFactory4pComDefaultAuthenticationManager struct {

    mPrototype * authentication0x38bbbc.DefaultAuthenticationManager

	

}

func (inst * comFactory4pComDefaultAuthenticationManager) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultAuthenticationManager) newObject() * authentication0x38bbbc.DefaultAuthenticationManager {
	return & authentication0x38bbbc.DefaultAuthenticationManager {}
}

func (inst * comFactory4pComDefaultAuthenticationManager) castObject(instance application.ComponentInstance) * authentication0x38bbbc.DefaultAuthenticationManager {
	return instance.Get().(*authentication0x38bbbc.DefaultAuthenticationManager)
}

func (inst * comFactory4pComDefaultAuthenticationManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultAuthenticationManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultAuthenticationManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultAuthenticationManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultAuthenticationManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultAuthenticationManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultAuthorizationManager : the factory of component: keeper-authorization-manager
type comFactory4pComDefaultAuthorizationManager struct {

    mPrototype * authorization0xbd4cc5.DefaultAuthorizationManager

	

}

func (inst * comFactory4pComDefaultAuthorizationManager) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultAuthorizationManager) newObject() * authorization0xbd4cc5.DefaultAuthorizationManager {
	return & authorization0xbd4cc5.DefaultAuthorizationManager {}
}

func (inst * comFactory4pComDefaultAuthorizationManager) castObject(instance application.ComponentInstance) * authorization0xbd4cc5.DefaultAuthorizationManager {
	return instance.Get().(*authorization0xbd4cc5.DefaultAuthorizationManager)
}

func (inst * comFactory4pComDefaultAuthorizationManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultAuthorizationManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultAuthorizationManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultAuthorizationManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultAuthorizationManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultAuthorizationManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComKeeperCore : the factory of component: keeper-core
type comFactory4pComKeeperCore struct {

    mPrototype * core0xb21891.KeeperCore

	
	mAuthenticationsSelector config.InjectionSelector
	mAuthorizationsSelector config.InjectionSelector
	mSubjectsSelector config.InjectionSelector
	mSessionProviderSelector config.InjectionSelector
	mAuthenticatorsSelector config.InjectionSelector
	mAuthorizersSelector config.InjectionSelector
	mConfigurersSelector config.InjectionSelector
	mSessionProvidersSelector config.InjectionSelector

}

func (inst * comFactory4pComKeeperCore) init() application.ComponentFactory {

	
	inst.mAuthenticationsSelector = config.NewInjectionSelector("#keeper-authentication-manager",nil)
	inst.mAuthorizationsSelector = config.NewInjectionSelector("#keeper-authorization-manager",nil)
	inst.mSubjectsSelector = config.NewInjectionSelector("#keeper-subject-manager",nil)
	inst.mSessionProviderSelector = config.NewInjectionSelector("#keeper-main-session-provider",nil)
	inst.mAuthenticatorsSelector = config.NewInjectionSelector(".keeper-authenticator-registry",nil)
	inst.mAuthorizersSelector = config.NewInjectionSelector(".keeper-authorizer-registry",nil)
	inst.mConfigurersSelector = config.NewInjectionSelector(".keeper-configurer",nil)
	inst.mSessionProvidersSelector = config.NewInjectionSelector(".keeper-session-provider-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComKeeperCore) newObject() * core0xb21891.KeeperCore {
	return & core0xb21891.KeeperCore {}
}

func (inst * comFactory4pComKeeperCore) castObject(instance application.ComponentInstance) * core0xb21891.KeeperCore {
	return instance.Get().(*core0xb21891.KeeperCore)
}

func (inst * comFactory4pComKeeperCore) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComKeeperCore) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComKeeperCore) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComKeeperCore) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComKeeperCore) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComKeeperCore) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Authentications = inst.getterForFieldAuthenticationsSelector(context)
	obj.Authorizations = inst.getterForFieldAuthorizationsSelector(context)
	obj.Subjects = inst.getterForFieldSubjectsSelector(context)
	obj.SessionProvider = inst.getterForFieldSessionProviderSelector(context)
	obj.Authenticators = inst.getterForFieldAuthenticatorsSelector(context)
	obj.Authorizers = inst.getterForFieldAuthorizersSelector(context)
	obj.Configurers = inst.getterForFieldConfigurersSelector(context)
	obj.SessionProviders = inst.getterForFieldSessionProvidersSelector(context)
	return context.LastError()
}

//getterForFieldAuthenticationsSelector
func (inst * comFactory4pComKeeperCore) getterForFieldAuthenticationsSelector (context application.InstanceContext) keeper0x6d39ef.AuthenticationManager {

	o1 := inst.mAuthenticationsSelector.GetOne(context)
	o2, ok := o1.(keeper0x6d39ef.AuthenticationManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "keeper-core")
		eb.Set("field", "Authentications")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.AuthenticationManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldAuthorizationsSelector
func (inst * comFactory4pComKeeperCore) getterForFieldAuthorizationsSelector (context application.InstanceContext) keeper0x6d39ef.AuthorizationManager {

	o1 := inst.mAuthorizationsSelector.GetOne(context)
	o2, ok := o1.(keeper0x6d39ef.AuthorizationManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "keeper-core")
		eb.Set("field", "Authorizations")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.AuthorizationManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSubjectsSelector
func (inst * comFactory4pComKeeperCore) getterForFieldSubjectsSelector (context application.InstanceContext) keeper0x6d39ef.SubjectManager {

	o1 := inst.mSubjectsSelector.GetOne(context)
	o2, ok := o1.(keeper0x6d39ef.SubjectManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "keeper-core")
		eb.Set("field", "Subjects")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.SubjectManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSessionProviderSelector
func (inst * comFactory4pComKeeperCore) getterForFieldSessionProviderSelector (context application.InstanceContext) keeper0x6d39ef.SessionProvider {

	o1 := inst.mSessionProviderSelector.GetOne(context)
	o2, ok := o1.(keeper0x6d39ef.SessionProvider)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "keeper-core")
		eb.Set("field", "SessionProvider")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.SessionProvider")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldAuthenticatorsSelector
func (inst * comFactory4pComKeeperCore) getterForFieldAuthenticatorsSelector (context application.InstanceContext) []keeper0x6d39ef.AuthenticatorRegistry {
	list1 := inst.mAuthenticatorsSelector.GetList(context)
	list2 := make([]keeper0x6d39ef.AuthenticatorRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(keeper0x6d39ef.AuthenticatorRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldAuthorizersSelector
func (inst * comFactory4pComKeeperCore) getterForFieldAuthorizersSelector (context application.InstanceContext) []keeper0x6d39ef.AuthorizerRegistry {
	list1 := inst.mAuthorizersSelector.GetList(context)
	list2 := make([]keeper0x6d39ef.AuthorizerRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(keeper0x6d39ef.AuthorizerRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldConfigurersSelector
func (inst * comFactory4pComKeeperCore) getterForFieldConfigurersSelector (context application.InstanceContext) []keeper0x6d39ef.Configurer {
	list1 := inst.mConfigurersSelector.GetList(context)
	list2 := make([]keeper0x6d39ef.Configurer, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(keeper0x6d39ef.Configurer)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldSessionProvidersSelector
func (inst * comFactory4pComKeeperCore) getterForFieldSessionProvidersSelector (context application.InstanceContext) []keeper0x6d39ef.SessionProviderRegistry {
	list1 := inst.mSessionProvidersSelector.GetList(context)
	list2 := make([]keeper0x6d39ef.SessionProviderRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(keeper0x6d39ef.SessionProviderRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultSessionProvider : the factory of component: com3-session0x14c51e.DefaultSessionProvider
type comFactory4pComDefaultSessionProvider struct {

    mPrototype * session0x14c51e.DefaultSessionProvider

	

}

func (inst * comFactory4pComDefaultSessionProvider) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultSessionProvider) newObject() * session0x14c51e.DefaultSessionProvider {
	return & session0x14c51e.DefaultSessionProvider {}
}

func (inst * comFactory4pComDefaultSessionProvider) castObject(instance application.ComponentInstance) * session0x14c51e.DefaultSessionProvider {
	return instance.Get().(*session0x14c51e.DefaultSessionProvider)
}

func (inst * comFactory4pComDefaultSessionProvider) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultSessionProvider) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultSessionProvider) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultSessionProvider) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComDefaultSessionProvider) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSessionProvider) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMainSessionProvider : the factory of component: keeper-main-session-provider
type comFactory4pComMainSessionProvider struct {

    mPrototype * session0x14c51e.MainSessionProvider

	
	mSelectorSelector config.InjectionSelector
	mProvidersSelector config.InjectionSelector

}

func (inst * comFactory4pComMainSessionProvider) init() application.ComponentFactory {

	
	inst.mSelectorSelector = config.NewInjectionSelector("${keeper.session-provider.name}",nil)
	inst.mProvidersSelector = config.NewInjectionSelector(".keeper-session-provider-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMainSessionProvider) newObject() * session0x14c51e.MainSessionProvider {
	return & session0x14c51e.MainSessionProvider {}
}

func (inst * comFactory4pComMainSessionProvider) castObject(instance application.ComponentInstance) * session0x14c51e.MainSessionProvider {
	return instance.Get().(*session0x14c51e.MainSessionProvider)
}

func (inst * comFactory4pComMainSessionProvider) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMainSessionProvider) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMainSessionProvider) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMainSessionProvider) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComMainSessionProvider) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainSessionProvider) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Selector = inst.getterForFieldSelectorSelector(context)
	obj.Providers = inst.getterForFieldProvidersSelector(context)
	return context.LastError()
}

//getterForFieldSelectorSelector
func (inst * comFactory4pComMainSessionProvider) getterForFieldSelectorSelector (context application.InstanceContext) string {
    return inst.mSelectorSelector.GetString(context)
}

//getterForFieldProvidersSelector
func (inst * comFactory4pComMainSessionProvider) getterForFieldProvidersSelector (context application.InstanceContext) []keeper0x6d39ef.SessionProviderRegistry {
	list1 := inst.mProvidersSelector.GetList(context)
	list2 := make([]keeper0x6d39ef.SessionProviderRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(keeper0x6d39ef.SessionProviderRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultSubjectManager : the factory of component: keeper-subject-manager
type comFactory4pComDefaultSubjectManager struct {

    mPrototype * subject0xae5c7a.DefaultSubjectManager

	

}

func (inst * comFactory4pComDefaultSubjectManager) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultSubjectManager) newObject() * subject0xae5c7a.DefaultSubjectManager {
	return & subject0xae5c7a.DefaultSubjectManager {}
}

func (inst * comFactory4pComDefaultSubjectManager) castObject(instance application.ComponentInstance) * subject0xae5c7a.DefaultSubjectManager {
	return instance.Get().(*subject0xae5c7a.DefaultSubjectManager)
}

func (inst * comFactory4pComDefaultSubjectManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultSubjectManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultSubjectManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultSubjectManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSubjectManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSubjectManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




