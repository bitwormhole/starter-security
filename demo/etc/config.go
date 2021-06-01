package etc

import (
	"github.com/bitwormhole/starter-security/auths"
	"github.com/bitwormhole/starter-security/auths/common"
	"github.com/bitwormhole/starter-security/auths/password"
	"github.com/bitwormhole/starter-security/demo/element"
	"github.com/bitwormhole/starter/application"
)

func authService(inst *auths.DefaultAuthenticationManager, ctx application.Context) error {

	// [component]
	//    id=auths

	return inst.Inject(ctx)
}

func authByPassword(inst *common.SimpleAuthWrapper, ctx application.Context) error {

	// [component]
	//    class=auth-mechanism

	const mechanism = "PASSWORD"

	provider := &password.AuthProvider{}
	provider.Mechanism = mechanism
	provider.DAO = &element.DemoPasswordAuthDAO{}

	inst.Mechanism = mechanism
	inst.Inner = provider

	return nil
}
