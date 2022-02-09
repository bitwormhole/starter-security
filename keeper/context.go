package keeper

// SecurityContext 安全上下文
type SecurityContext interface {
	GetAuthentications() AuthenticationManager

	GetAuthorizations() AuthorizationManager

	GetSubjects() SubjectManager

	GetSessionProvider() SessionProvider
}

// Context 默认的安全上下文
type Context struct {
	Authentications AuthenticationManager

	Authorizations AuthorizationManager

	Subjects SubjectManager

	SessionProvider SessionProvider
}

func (inst *Context) _Impl() SecurityContext {
	return inst
}

func (inst *Context) GetAuthentications() AuthenticationManager {
	return inst.Authentications
}

func (inst *Context) GetAuthorizations() AuthorizationManager {
	return inst.Authorizations
}

func (inst *Context) GetSubjects() SubjectManager {
	return inst.Subjects
}

func (inst *Context) GetSessionProvider() SessionProvider {
	return inst.SessionProvider
}
