package keeper

import "context"

// Subject 代表操作的主体
type Subject interface {
	GetSession() Session

	GetAccess() Access

	GetContext() context.Context

	IsAuthenticated() bool

	SetSession(s Session)

	SetAccess(a Access)

	SetAuthenticated(authenticated bool)

	Login(ctx context.Context, a Authentication) (Identity, error)

	Logout() error

	Authorize(ctx context.Context, a Access) (bool, error)
}

// SubjectManager 主体管理器
type SubjectManager interface {
	GetSubject(ctx context.Context) (Subject, error)
}
