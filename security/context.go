package security

// Context 安全上下文
type Context interface {
	GetAuthenticationManager() AuthenticationManager

	GetAuthorizationManager() AuthorizationManager

	GetSubjectManager() SubjectManager

	GetSessionIDGenerator() SessionIDGenerator
	GetSessionFactory() SessionFactory
	GetSessionManager() SessionManager
}
