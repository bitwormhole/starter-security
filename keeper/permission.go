package keeper

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper/users"
)

// Permission 表示一个许可实例 (路径中不带参数)
type Permission interface {
	Method() string

	Path() string

	Owner() Identity

	Friends() UserGroup

	Template() PermissionTemplate

	IsOwner(user Identity) bool

	IsFriend(user Identity) bool

	AcceptUser(user Identity) bool

	AcceptRole(role users.Role) bool

	AcceptRoles(roles users.Roles) bool
}

// PermissionTemplate 表示一个许可模板 (路径中带参数)
type PermissionTemplate interface {
	Method() string

	PathPattern() string

	AcceptRole(role users.Role) bool

	AcceptRoles(roles users.Roles) bool

	LoadPermission(params map[string]string) (Permission, error)
}

// PermissionTemplateFactory 表示一个许可模板工厂
type PermissionTemplateFactory interface {
	CreateTemplate(spr *SimplePermissionRegistration) (PermissionTemplate, error)
}

// PermissionLoader 表示一个许可加载器
type PermissionLoader interface {
	Load(template PermissionTemplate, params map[string]string) (Permission, error)
}

// PermissionLoaderFactory 表示一个许可加载器工厂
type PermissionLoaderFactory interface {
	CreateLoader(spr *SimplePermissionRegistration) (PermissionLoader, error)
}

// PermissionManager 许可管理器
type PermissionManager interface {
	FindTemplate(ctx context.Context, a Access) (PermissionTemplate, error)
}

// ComplexPermissionRegistration 复合的身份验证器注册项
type ComplexPermissionRegistration struct {
	Methods []string     // 操作方法表达式
	Paths   []string     // 路径模板表达式
	Roles   []users.Role // 操作角色表达式

	Enabled                 bool
	LoaderFactorySelector   string
	TemplateFactorySelector string

	Loader   PermissionLoader   // 注册时如果为nil，就使用默认的处理器
	Template PermissionTemplate // 注册时如果为nil，就使用默认的处理器
}

// SimplePermissionRegistration 简单的身份验证器注册项
type SimplePermissionRegistration struct {
	Method      string       // 操作方法表达式
	PathPattern string       // 路径模板表达式
	Roles       []users.Role // 操作角色表达式

	Enabled                 bool
	LoaderFactorySelector   string
	TemplateFactorySelector string

	Loader   PermissionLoader   // 注册时如果为nil，就使用默认的处理器
	Template PermissionTemplate // 注册时如果为nil，就使用默认的处理器
}

// PermissionRegistry 许可注册器
// 【inject:".keeper-permission-registry"】
type PermissionRegistry interface {
	GetRegistrationList() []*ComplexPermissionRegistration
}
