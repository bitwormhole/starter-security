package security

import "github.com/bitwormhole/starter/lang"

// AuthenticationToken 验证输入
type AuthenticationToken interface {

	// GetPrincipal 获取身份
	GetPrincipal() lang.Object

	// GetCredentials 获取凭证
	GetCredentials() lang.Object
}

// AuthenticationInfo 验证输出
type AuthenticationInfo interface {

	// GetPrincipal 获取身份
	GetPrincipals() PrincipalCollection

	// GetCredentials 获取凭证
	GetCredentials() lang.Object
}
