package common

type Auth interface {
	AccountID() string
	AccountUUID() string
	Secret() string
	Salt() string
	Mechanism() string
}

type AuthDAO interface {
	Find(user string, mechanism string) (Auth, error)
}
