package security

// Access 访问参数
type Access interface {
	SessionID() string
	Path() string
	Method() string

	SetSessionID(sid string)
}
