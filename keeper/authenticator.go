package keeper

import "context"

// Identity 身份
type Identity interface {
	UserID() string
	UserUUID() string
	Nickname() string
	Avatar() string
	Roles() []string
}

// Authentication 身份验证请求
type Authentication interface {
	Mechanism() string
	UserID() string
	UserSecret() []byte
}

// Authenticator 身份验证器
type Authenticator interface {
	Supports(ctx context.Context, a Authentication) bool
	Verify(ctx context.Context, a Authentication) (Identity, error)
}

// AuthenticationManager 验证管理器
type AuthenticationManager interface {
	Authenticate(ctx context.Context, a Authentication) (Identity, error)
}

// AuthenticatorRegistration 身份验证器注册项
type AuthenticatorRegistration struct {
	Name          string
	Authenticator Authenticator
}

// AuthenticatorRegistry 身份验证器注册器
// 【inject:".keeper-authenticator-registry"】
type AuthenticatorRegistry interface {
	GetRegistrationList() []*AuthenticatorRegistration
}
