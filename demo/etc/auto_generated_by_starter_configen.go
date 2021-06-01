// 这个文件是由 starter-configen 工具生成的配置代码，千万不要手工修改里面的任何内容。
package etc

import(
	auths_51a30f66 "github.com/bitwormhole/starter-security/auths"
	common_2ec7a97f "github.com/bitwormhole/starter-security/auths/common"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func Config(cb application.ConfigBuilder) error {

    // authByPassword
    cb.AddComponent(&config.ComInfo{
		ID: "authByPassword",
		Class: "auth-mechanism",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &common_2ec7a97f.SimpleAuthWrapper{}
		},
		OnInject: func(obj lang.Object,context application.Context) error {
		    target := obj.(*common_2ec7a97f.SimpleAuthWrapper)
		    return authByPassword(target,context)
		},
    })

    // authService
    cb.AddComponent(&config.ComInfo{
		ID: "auths",
		Class: "",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &auths_51a30f66.DefaultAuthenticationManager{}
		},
		OnInject: func(obj lang.Object,context application.Context) error {
		    target := obj.(*auths_51a30f66.DefaultAuthenticationManager)
		    return authService(target,context)
		},
    })

    return nil
}

