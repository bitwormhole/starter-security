package security

import "context"

// SessionManager 会话管理器
type SessionManager interface {
	Start(ctx context.Context) Session

	GetSession(key SessionKey) (Session, error)
}
