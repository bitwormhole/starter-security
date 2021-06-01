package auths

import "github.com/bitwormhole/starter/lang"

// Authentication (认证) token
type AuthenticationToken interface {

	// 实现机制
	Mechanism() string

	// 委托人
	Principal() lang.Object

	// 凭证
	Credentials() lang.Object
}

// Authentication (认证) info
type AuthenticationInfo interface {
	// 委托人
	Principal() lang.Object
	Roles() string
	Success() bool
	Message() string
}

// Authentication 服务提供者
type AuthenticationProvider interface {
	Name() string
	Supports(token AuthenticationToken) bool
	Authenticate(token AuthenticationToken) (AuthenticationInfo, error)
}

// Authentication 管理器
type AuthenticationManager interface {
	Providers() []AuthenticationProvider
	Authenticate(token AuthenticationToken) (AuthenticationInfo, error)
}
