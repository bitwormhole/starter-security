// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package configlib

import (
	keeper0x6d39ef "github.com/bitwormhole/starter-security/keeper"
	authentication0x38bbbc "github.com/bitwormhole/starter-security/keeper/support/authentication"
	authorization0xbd4cc5 "github.com/bitwormhole/starter-security/keeper/support/authorization"
	core0xb21891 "github.com/bitwormhole/starter-security/keeper/support/core"
	permission0xf95196 "github.com/bitwormhole/starter-security/keeper/support/permission"
	session0x14c51e "github.com/bitwormhole/starter-security/keeper/support/session"
	subject0xae5c7a "github.com/bitwormhole/starter-security/keeper/support/subject"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComDefaultAuthenticationManager struct {
	instance *authentication0x38bbbc.DefaultAuthenticationManager
	 markup0x23084a.Component `id:"keeper-authentication-manager" class:"keeper-configurer"`
	RegistryList []keeper0x6d39ef.AuthenticatorRegistry `inject:".keeper-authenticator-registry"`
}


type pComDefaultAuthorizationManager struct {
	instance *authorization0xbd4cc5.DefaultAuthorizationManager
	 markup0x23084a.Component `id:"keeper-authorization-manager" class:"keeper-configurer"`
	AuthorizerRegistryList []keeper0x6d39ef.AuthorizerRegistry `inject:".keeper-authorizer-registry"`
}


type pComKeeperCore struct {
	instance *core0xb21891.KeeperCore
	 markup0x23084a.Component `id:"keeper-core"  initMethod:"Init"`
	Authentications keeper0x6d39ef.AuthenticationManager `inject:"#keeper-authentication-manager"`
	Authorizations keeper0x6d39ef.AuthorizationManager `inject:"#keeper-authorization-manager"`
	Subjects keeper0x6d39ef.SubjectManager `inject:"#keeper-subject-manager"`
	SessionProvider keeper0x6d39ef.SessionProvider `inject:"#keeper-main-session-provider"`
	Authenticators []keeper0x6d39ef.AuthenticatorRegistry `inject:".keeper-authenticator-registry"`
	Authorizers []keeper0x6d39ef.AuthorizerRegistry `inject:".keeper-authorizer-registry"`
	Configurers []keeper0x6d39ef.Configurer `inject:".keeper-configurer"`
	SessionProviders []keeper0x6d39ef.SessionProviderRegistry `inject:".keeper-session-provider-registry"`
}


type pComDefaultPermissionLoaderFactory struct {
	instance *permission0xf95196.DefaultPermissionLoaderFactory
	 markup0x23084a.Component `id:"keeper-default-permission-loader-factory"`
}


type pComDefaultPermissionManager struct {
	instance *permission0xf95196.DefaultPermissionManager
	 markup0x23084a.Component `id:"keeper-permission-manager" class:"keeper-configurer"`
	AppCtx application0x67f6c5.Context `inject:"context"`
	PermissionRegistryList []keeper0x6d39ef.PermissionRegistry `inject:".keeper-permission-registry"`
	PermissionsPropertiesResourceName string `inject:"${security.keeper.permissions.properties}"`
}


type pComDefaultPermissionTemplateFactory struct {
	instance *permission0xf95196.DefaultPermissionTemplateFactory
	 markup0x23084a.Component `id:"keeper-default-permission-template-factory"`
}


type pComDefaultSessionProvider struct {
	instance *session0x14c51e.DefaultSessionProvider
	 markup0x23084a.Component `class:"keeper-session-provider-registry" initMethod:"Init"`
}


type pComMainSessionProvider struct {
	instance *session0x14c51e.MainSessionProvider
	 markup0x23084a.Component `id:"keeper-main-session-provider"  class:"keeper-configurer"  initMethod:"Init"`
	Selector string `inject:"${keeper.session-provider.name}"`
	Providers []keeper0x6d39ef.SessionProviderRegistry `inject:".keeper-session-provider-registry"`
}


type pComDefaultSubjectManager struct {
	instance *subject0xae5c7a.DefaultSubjectManager
	 markup0x23084a.Component `id:"keeper-subject-manager" class:"keeper-configurer" `
}

