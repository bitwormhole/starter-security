package security

import (
	"errors"
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const SessionKeyName = "github.com/bitwormhole/starter-gin/security/Session#binding"

type Session interface {
	Load() error
	Reload() error
	Save() error
	Commit() error
	Properties() collection.Properties
	GetInfo() *SessionInfo
	SetInfo(*SessionInfo)
}

type SessionInfo struct {
	DisplayName string
	AvatarURL   string
	Account     string
	Roles       string
	auth        bool
	OpenTime    time.Time
	CloseTime   time.Time
}

type SessionFactory interface {
	Create(context application.SimpleContext) Session
}

func GetSession(context application.SimpleContext) (Session, error) {
	obj := context.GetAttribute(SessionKeyName)
	if obj == nil {
		return nil, errors.New("No session is been loaded.")
	}
	session, ok := obj.(Session)
	if !ok {
		return nil, errors.New("No session is been loaded. (convert.ok==false)")
	}
	return session, nil
}

func OpenSession(context application.SimpleContext, factory SessionFactory) (Session, error) {
	session, err := GetSession(context)
	if err == nil {
		return session, err
	}
	if factory == nil {
		return nil, errors.New("SessionFactory==nil")
	}
	session = factory.Create(context)
	session.Load()
	context.SetAttribute(SessionKeyName, session)
	return session, nil
}

func ToSC(context application.Context) application.SimpleContext {
	return context.GetAttributes()
}