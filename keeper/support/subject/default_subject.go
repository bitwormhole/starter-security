package subject

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter-security/keeper"
)

type DefaultSubject struct {
	sc *keeper.SessionContext
}

func (inst *DefaultSubject) _Impl() keeper.Subject {
	return inst
}

func (inst *DefaultSubject) GetSession(create bool) (keeper.Session, error) {
	session := inst.sc.Session
	if session != nil {
		return session, nil
	}

	if !create {
		return nil, errors.New("session is nil")
	}

	// make new
	session, err := inst.loadSession()
	if err != nil {
		return nil, err
	}

	inst.sc.Session = session
	return session, nil
}

func (inst *DefaultSubject) loadSession() (keeper.Session, error) {
	sc := inst.sc.SecurityContext
	adapter := inst.sc.Adapter
	factory := sc.GetSessionProvider().GetSessionFactory()
	if adapter == nil {
		return nil, errors.New("the use of keeper.SessionAdapter without setup")
	}
	session, err := factory.Create(adapter)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (inst *DefaultSubject) GetAccess() keeper.Access {
	return inst.sc.Access
}

func (inst *DefaultSubject) GetContext() context.Context {
	return inst.sc.Context
}

func (inst *DefaultSubject) IsAuthenticated() bool {
	const name = keeper.SessionFieldAuthenticated
	session, err := inst.GetSession(false)
	if err != nil {
		return false
	}
	getter := session.Properties().Getter()
	return getter.GetBool(name, false)
}

func (inst *DefaultSubject) SetSession(s keeper.Session) {
	if s == nil {
		return
	}
	inst.sc.Session = s
}

func (inst *DefaultSubject) SetAccess(a keeper.Access) {
	if a == nil {
		return
	}
	inst.sc.Access = a
}

func (inst *DefaultSubject) SetAuthenticated(authenticated bool) {
	const name = keeper.SessionFieldAuthenticated
	session, err := inst.GetSession(true)
	if err != nil {
		panic(err)
	}
	setter := session.Properties().Setter()
	setter.SetBool(name, authenticated)
}

func (inst *DefaultSubject) Login(ctx context.Context, a keeper.Authentication) (keeper.Identity, error) {
	return nil, errors.New("no impl")
}

func (inst *DefaultSubject) Logout() error {
	return errors.New("no impl")
}

func (inst *DefaultSubject) Authorize(ctx context.Context, a keeper.Access) (bool, error) {
	return false, errors.New("no impl")
}
