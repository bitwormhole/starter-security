// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package configdemo

import (
	keeper0x6d39ef "github.com/bitwormhole/starter-security/keeper"
	demo10xf489b5 "github.com/bitwormhole/starter-security/src/demo/golang/demo/demo1"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComDemo1 struct {
	instance *demo10xf489b5.Demo1
	 markup0x23084a.Component `class:"life"`
	AC application0x67f6c5.Context `inject:"context"`
	Subjects keeper0x6d39ef.SubjectManager `inject:"#keeper-subject-manager"`
}


type pComDemo1authorizer struct {
	instance *demo10xf489b5.Demo1authorizer
	 markup0x23084a.Component `class:"keeper-authorizer-registry"`
}


type pComDemo1passwordAuth struct {
	instance *demo10xf489b5.Demo1passwordAuth
	 markup0x23084a.Component `class:"keeper-authenticator-registry"`
}

