package keeper

import (
	"context"
)

// Subject 代表操作的主体
type Subject interface {

	// getters

	GetSession(create bool) (Session, error)

	IsAuthenticated() bool

	// setter

	SetSession(s Session)

	SetAuthenticated(authenticated bool)

	// operation

	Login(ctx context.Context, a Authentication) (Identity, error)

	Logout(ctx context.Context) error

	Authorize(ctx context.Context) error

	HasPermission(ctx context.Context) bool
}

// SubjectManager 主体管理器
// 【inject:"#keeper-subject-manager"】
type SubjectManager interface {
	GetSubject(ctx context.Context) (Subject, error)
}
