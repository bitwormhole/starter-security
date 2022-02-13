package keeper

import (
	"context"
)

// Authentication 身份验证请求
type Authentication interface {
	Mechanism() string
	User() string
	Secret() []byte
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
	Name          string // 名称是全局唯一的
	Mechanism     string // 多个 Authenticator 可以支持同一种机制
	Authenticator Authenticator
}

// AuthenticatorRegistry 身份验证器注册器
// 【inject:".keeper-authenticator-registry"】
type AuthenticatorRegistry interface {
	GetRegistrationList() []*AuthenticatorRegistration
}
