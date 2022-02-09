package core

import (
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// KeeperCore the core of keeper
type KeeperCore struct {
	markup.Component `id:"keeper-core"  initMethod:"Init"`

	Authentications keeper.AuthenticationManager `inject:"#keeper-authentication-manager"`
	Authorizations  keeper.AuthorizationManager  `inject:"#keeper-authorization-manager"`
	Subjects        keeper.SubjectManager        `inject:"#keeper-subject-manager"`
	SessionProvider keeper.SessionProvider       `inject:"#keeper-main-session-provider"`

	Authenticators   []keeper.AuthenticatorRegistry   `inject:".keeper-authenticator-registry"`
	Authorizers      []keeper.AuthorizerRegistry      `inject:".keeper-authorizer-registry"`
	Configurers      []keeper.Configurer              `inject:".keeper-configurer"`
	SessionProviders []keeper.SessionProviderRegistry `inject:".keeper-session-provider-registry"`

	context *keeper.Context
}

func (inst *KeeperCore) _Impl() application.Looper {
	return inst
}

// Init ...
func (inst *KeeperCore) Init() error {
	return inst.config()
}

// Loop NOP
func (inst *KeeperCore) Loop() error {
	return nil
}

func (inst *KeeperCore) config() error {
	ctx := &keeper.Context{}
	all := inst.Configurers
	for _, item := range all {
		err := item.Configure(ctx)
		if err != nil {
			return err
		}
	}
	inst.context = ctx
	return nil
}
