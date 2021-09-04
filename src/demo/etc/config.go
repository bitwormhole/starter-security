package etc

import (
	"github.com/bitwormhole/starter-security/auths"
	"github.com/bitwormhole/starter-security/auths/common"
	"github.com/bitwormhole/starter-security/auths/password"
	"github.com/bitwormhole/starter-security/demo/element"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

type authService struct {
	markup.Component
	instance *auths.DefaultAuthenticationManager `id:"auths"`

	ProviderList []auths.AuthenticationProvider `inject:".auth-provider"`
}

type authByPassword struct {
	markup.Component `class:"auth-provider"`
	instance         *common.SimpleAuthWrapper `injectMethod:"inject"`
}

func (inst *authByPassword) inject(injection application.Injection) error {

	instance := inst.instance
	const mechanism = "PASSWORD"

	provider := &password.AuthProvider{}
	provider.Mechanism = mechanism
	provider.DAO = &element.DemoPasswordAuthDAO{}

	instance.Mechanism = mechanism
	instance.Inner = provider

	return nil
}
