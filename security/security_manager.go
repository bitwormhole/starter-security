package security

// SecurityManager 安全管理器
type SecurityManager interface {

	// 同时实现以下三个接口
	Authenticator
	Authorizer
	SessionManager

	Login(subject Subject, token AuthenticationToken) (Subject, error)

	Logout(subject Subject)

	CreateSubject(context SubjectContext) Subject
}
