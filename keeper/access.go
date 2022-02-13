package keeper

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper/users"
)

// Access 访问参数
type Access interface {
	Method() string

	Path() string

	PathPattern() string

	Params() map[string]string

	// 这两个方法【已废弃】，用代替 SessionAdapter
	// GetSessionData() []byte
	// SetSessionData(data []byte)
}

// SecurityAccess 安全的访问参数
type SecurityAccess interface {
	Access

	// getters

	GetContext() context.Context

	GetSubject() Subject

	GetPermission() Permission

	// Session.GetRoles() 和 Access.GetRoles() 分别代表两个作用域的角色，
	// Session > Access
	GetRoles() users.Roles

	// setters

	SetRoles(roles users.Roles)

	SetSubject(s Subject)

	SetPermission(p Permission)
}

////////////////////////////////////////////////////////////////////////////////

// AccessContext 会话上下文
type AccessContext struct {
	Access          Access
	Adapter         SessionAdapter
	Context         context.Context
	Permission      Permission
	Roles           users.Roles
	SecurityAccess  SecurityAccess
	SecurityContext SecurityContext
	Session         Session
	Subject         Subject
}

////////////////////////////////////////////////////////////////////////////////

// AccessBuilder ...
type AccessBuilder struct {
	Method      string
	Path        string
	PathPattern string
	Params      map[string]string
}

func (inst *AccessBuilder) Create() Access {
	a := &defaultAccess{
		method:      inst.Method,
		path:        inst.Path,
		pathPattern: inst.PathPattern,
		params:      inst.Params,
	}
	return a
}

// defaultAccess ...
type defaultAccess struct {
	method      string
	path        string
	pathPattern string
	params      map[string]string
}

func (inst *defaultAccess) _Impl() Access {
	return inst
}

func (inst *defaultAccess) Method() string {
	return inst.method
}

func (inst *defaultAccess) Path() string {
	return inst.path
}

func (inst *defaultAccess) PathPattern() string {
	return inst.pathPattern
}

func (inst *defaultAccess) Params() map[string]string {
	return inst.params
}

////////////////////////////////////////////////////////////////////////////////

// DefaultSecurityAccess ...
type DefaultSecurityAccess struct {
	AccessContext *AccessContext
	Access        Access
}

func (inst *DefaultSecurityAccess) _Impl() SecurityAccess {
	return inst
}

func (inst *DefaultSecurityAccess) Init(ac *AccessContext) SecurityAccess {
	inst.Access = ac.Access
	inst.AccessContext = ac
	return inst
}

func (inst *DefaultSecurityAccess) GetContext() context.Context {
	return inst.AccessContext.Context
}

func (inst *DefaultSecurityAccess) GetSubject() Subject {
	return inst.AccessContext.Subject
}

func (inst *DefaultSecurityAccess) GetPermission() Permission {
	return inst.AccessContext.Permission
}

func (inst *DefaultSecurityAccess) GetRoles() users.Roles {
	return inst.AccessContext.Roles
}

func (inst *DefaultSecurityAccess) SetRoles(roles users.Roles) {
	inst.AccessContext.Roles = roles
}

func (inst *DefaultSecurityAccess) SetSubject(s Subject) {
	inst.AccessContext.Subject = s
}

func (inst *DefaultSecurityAccess) SetPermission(p Permission) {
	inst.AccessContext.Permission = p
}

func (inst *DefaultSecurityAccess) Method() string {
	return inst.Access.Method()
}

func (inst *DefaultSecurityAccess) Path() string {
	return inst.Access.Path()
}

func (inst *DefaultSecurityAccess) PathPattern() string {
	return inst.Access.PathPattern()
}

func (inst *DefaultSecurityAccess) Params() map[string]string {
	return inst.Access.Params()
}
