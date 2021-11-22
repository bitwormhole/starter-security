package security

import (
	"time"

	"github.com/bitwormhole/starter/lang"
)

// SessionKey 跟session绑定的键
type SessionKey interface {
	GetSessionId() string
}

const SessionKeyName = "github.com/bitwormhole/starter-gin/security/Session#binding"

// Session 表示一个具体的会话
type Session interface {
	GetID() string

	GetStartTimestamp() time.Time

	GetLastAccessTime() time.Time

	GetTimeout() (int64, error)

	SetTimeout(maxIdleTimeInMillis int64) error

	GetHost() string

	Touch() error

	Stop() error

	GetAttributeKeys() ([]string, error)

	GetAttribute(key string) (lang.Object, error)

	SetAttribute(key string, value lang.Object) error

	RemoveAttribute(key string) error
}
