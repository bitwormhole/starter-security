package security

import "github.com/bitwormhole/starter/lang"

// SessionDAO 是Session的数据访问对象
type SessionDAO interface {
	Create(session Session) lang.Serializable

	ReadSession(sessionID lang.Serializable) (Session, error)

	Update(session Session) error

	Delete(session Session)

	GetActiveSessions() []Session
}
