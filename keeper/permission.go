package keeper

import "context"

// Permission 表示一个许可处理器
type Permission interface {
	Accept(ctx context.Context, a Access) bool

	AcceptAnonymous() bool
}

// PermissionManager 许可管理器
type PermissionManager interface {
	FindPermissions(ctx context.Context, a Access) []Permission
}

// PermissionRegistration 身份验证器注册项
type PermissionRegistration struct {
	Methods  []string // 操作方法表达式
	Paths    []string // 资源路径表达式
	Subjects []string // 操作主体表达式

	Permission Permission // 注册时如果为nil，就使用默认的处理器
}

// PermissionRegistry 许可注册器
// 【inject:".keeper-permission-registry"】
type PermissionRegistry interface {
	GetRegistrationList() []*PermissionRegistration
}
