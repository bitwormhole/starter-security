// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package configdemo

import (
	keeper0x6d39ef "github.com/bitwormhole/starter-security/keeper"
	demo10xf489b5 "github.com/bitwormhole/starter-security/src/demo/golang/demo/demo1"
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

	// component: com0-demo10xf489b5.Demo1
	cominfobuilder.Next()
	cominfobuilder.ID("com0-demo10xf489b5.Demo1").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDemo1{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-demo10xf489b5.Demo1authorizer
	cominfobuilder.Next()
	cominfobuilder.ID("com1-demo10xf489b5.Demo1authorizer").Class("keeper-authorizer-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDemo1authorizer{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-demo10xf489b5.Demo1passwordAuth
	cominfobuilder.Next()
	cominfobuilder.ID("com2-demo10xf489b5.Demo1passwordAuth").Class("keeper-authenticator-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDemo1passwordAuth{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDemo1 : the factory of component: com0-demo10xf489b5.Demo1
type comFactory4pComDemo1 struct {

    mPrototype * demo10xf489b5.Demo1

	
	mACSelector config.InjectionSelector
	mSubjectsSelector config.InjectionSelector

}

func (inst * comFactory4pComDemo1) init() application.ComponentFactory {

	
	inst.mACSelector = config.NewInjectionSelector("context",nil)
	inst.mSubjectsSelector = config.NewInjectionSelector("#keeper-subject-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDemo1) newObject() * demo10xf489b5.Demo1 {
	return & demo10xf489b5.Demo1 {}
}

func (inst * comFactory4pComDemo1) castObject(instance application.ComponentInstance) * demo10xf489b5.Demo1 {
	return instance.Get().(*demo10xf489b5.Demo1)
}

func (inst * comFactory4pComDemo1) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDemo1) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDemo1) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDemo1) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AC = inst.getterForFieldACSelector(context)
	obj.Subjects = inst.getterForFieldSubjectsSelector(context)
	return context.LastError()
}

//getterForFieldACSelector
func (inst * comFactory4pComDemo1) getterForFieldACSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldSubjectsSelector
func (inst * comFactory4pComDemo1) getterForFieldSubjectsSelector (context application.InstanceContext) keeper0x6d39ef.SubjectManager {

	o1 := inst.mSubjectsSelector.GetOne(context)
	o2, ok := o1.(keeper0x6d39ef.SubjectManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-demo10xf489b5.Demo1")
		eb.Set("field", "Subjects")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.SubjectManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDemo1authorizer : the factory of component: com1-demo10xf489b5.Demo1authorizer
type comFactory4pComDemo1authorizer struct {

    mPrototype * demo10xf489b5.Demo1authorizer

	

}

func (inst * comFactory4pComDemo1authorizer) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDemo1authorizer) newObject() * demo10xf489b5.Demo1authorizer {
	return & demo10xf489b5.Demo1authorizer {}
}

func (inst * comFactory4pComDemo1authorizer) castObject(instance application.ComponentInstance) * demo10xf489b5.Demo1authorizer {
	return instance.Get().(*demo10xf489b5.Demo1authorizer)
}

func (inst * comFactory4pComDemo1authorizer) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDemo1authorizer) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDemo1authorizer) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDemo1authorizer) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1authorizer) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1authorizer) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDemo1passwordAuth : the factory of component: com2-demo10xf489b5.Demo1passwordAuth
type comFactory4pComDemo1passwordAuth struct {

    mPrototype * demo10xf489b5.Demo1passwordAuth

	

}

func (inst * comFactory4pComDemo1passwordAuth) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDemo1passwordAuth) newObject() * demo10xf489b5.Demo1passwordAuth {
	return & demo10xf489b5.Demo1passwordAuth {}
}

func (inst * comFactory4pComDemo1passwordAuth) castObject(instance application.ComponentInstance) * demo10xf489b5.Demo1passwordAuth {
	return instance.Get().(*demo10xf489b5.Demo1passwordAuth)
}

func (inst * comFactory4pComDemo1passwordAuth) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDemo1passwordAuth) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDemo1passwordAuth) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDemo1passwordAuth) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1passwordAuth) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1passwordAuth) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




