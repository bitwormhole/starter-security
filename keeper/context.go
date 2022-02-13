package keeper

// SecurityContext 安全上下文接口
type SecurityContext interface {
	GetAuthentications() AuthenticationManager

	GetAuthorizations() AuthorizationManager

	GetSubjects() SubjectManager

	GetPermissions() PermissionManager

	GetSessionProvider() SessionProvider
}

////////////////////////////////////////////////////////////////////////////////

// Context 安全上下文
type Context struct {
	Authentications AuthenticationManager

	Authorizations AuthorizationManager

	Subjects SubjectManager

	Permissions PermissionManager

	SessionProvider SessionProvider
}

func (inst *Context) _Impl() SecurityContext {
	return inst
}

func (inst *Context) GetPermissions() PermissionManager {
	return inst.Permissions
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

////////////////////////////////////////////////////////////////////////////////
