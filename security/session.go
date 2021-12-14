package security

// Session 会话
type Session interface {
	GetID() string // the session id

	GetProperty(name string) string
	SetProperty(name, value string)

	GetIdentity() Identity
	SetIdentity(ident Identity)
}

// SessionManager 会话管理器
type SessionManager interface {
	InsertSession(session Session) (Session, error)
	RemoveSession(sid string) error
	UpdateSession(sid string, session Session) (Session, error)
	FindSession(sid string) (Session, error)
}

// SessionFactory 会话工厂
type SessionFactory interface {
	CreateSession() Session
}

// SessionIDGenerator 会话ID生成器
type SessionIDGenerator interface {
	GenerateID() string
}
