package security

import "context"

// Subject 代表操作的主体
type Subject interface {
	GetSession() Session

	GetAccess() Access

	SetSession(s Session)

	SetAccess(a Access)

	Login(ctx context.Context, a Authentication) (Identity, error)

	Logout() error

	Authorize(ctx context.Context, a Access) (bool, error)
}

// SubjectManager 主体管理器
type SubjectManager interface {
	GetSubject(ctx context.Context) Subject
}
