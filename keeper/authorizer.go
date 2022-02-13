package keeper

import (
	"context"
)

// Authorizer 授权者
type Authorizer interface {
	Authorize(ctx context.Context, a Authorization) error
}

// AuthorizationManager 授权管理器
type AuthorizationManager interface {
	Authorize(ctx context.Context) error
	ListAuthorizers() []Authorizer

	//////// ListAuthorizerRegistrations() []*AuthorizerRegistration
}

// AuthorizerRegistration 授权者注册项
type AuthorizerRegistration struct {
	Name       string
	Scope      string // ["session","access","all"]
	Enabled    bool
	Authorizer Authorizer
}

// AuthorizerRegistry 授权者注册器
// 【inject:".keeper-authorizer-registry"】
type AuthorizerRegistry interface {
	GetRegistrationList() []*AuthorizerRegistration
}
