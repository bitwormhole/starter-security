package security

// Authenticator 身份验证者
type Authenticator interface {

	// Authenticate 进行身份验证
	Authenticate(token AuthenticationToken) (AuthenticationInfo, error)
}
