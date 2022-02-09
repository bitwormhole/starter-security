package keeper

// Access 访问参数
type Access interface {
	Path() string

	PathPattern() string

	Method() string

	GetSessionData() []byte

	SetSessionData(data []byte)
}
