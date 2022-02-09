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

func (inst *DefaultSubject) GetSession() keeper.Session {
	session := inst.sc.Session
	if session == nil {
		session = inst.loadSession()
		inst.sc.Session = session
	}
	return session
}

func (inst *DefaultSubject) loadSession() keeper.Session {
	sc := inst.sc.SecurityContext
	adapter := inst.sc.Adapter
	factory := sc.GetSessionProvider().GetSessionFactory()
	if adapter == nil {
		panic("the use of keeper.SessionAdapter without setup")
	}
	session, err := factory.Create(adapter)
	if err != nil {
		panic(err)
	}
	return session
}

func (inst *DefaultSubject) GetAccess() keeper.Access {
	return inst.sc.Access
}

func (inst *DefaultSubject) GetContext() context.Context {
	return inst.sc.Context
}

func (inst *DefaultSubject) IsAuthenticated() bool {
	const name = keeper.SessionFieldAuthenticated
	getter := inst.GetSession().Properties().Getter()
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
	setter := inst.GetSession().Properties().Setter()
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
