package security

// Realm 是一个可以通过访问应用程序特定的安全实体（例如：用户，角色，许可），来确认验证和授权操作的安全组件。
type Realm interface {
	GetName() string

	Supports(token AuthenticationToken) bool

	GetAuthenticationInfo(token AuthenticationToken) (AuthenticationInfo, error)
}
