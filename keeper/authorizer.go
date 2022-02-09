package keeper

import "context"

// Authorization 授权请求
type Authorization interface {
	Identity() Identity
	Method() string
	Path() string
}

// Authorizer 授权者
type Authorizer interface {
	Supports(ctx context.Context, a Authorization) bool
	Accept(ctx context.Context, a Authorization) bool
}

// AuthorizationManager 授权管理器
type AuthorizationManager interface {
	Accept(ctx context.Context, a Authorization) bool
}

// AuthorizerRegistration 授权者注册项
type AuthorizerRegistration struct {
	Name       string
	Authorizer Authorizer
}

// AuthorizerRegistry 授权者注册器
// 【inject:".keeper-authorizer-registry"】
type AuthorizerRegistry interface {
	GetRegistrationList() []*AuthorizerRegistration
}
