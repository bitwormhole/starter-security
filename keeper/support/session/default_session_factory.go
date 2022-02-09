package session

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/vlog"
)

////////////////////////////////////////////////////////////////////////////////

// DefaultSessionFactory ...
type DefaultSessionFactory struct {
}

func (inst *DefaultSessionFactory) _Impl() keeper.SessionFactory {
	return inst
}

// Create ...
func (inst *DefaultSessionFactory) Create(adapter keeper.SessionAdapter) (keeper.Session, error) {

	holder, err := keeper.GetHolder(adapter.GetContext())
	if err != nil {
		return nil, err
	}

	sc := holder.GetSessionContext()

	session := &DefaultSession{}
	session.adapter = adapter
	session.context = sc
	session.props = collection.CreateProperties()

	err = adapter.Load(session)
	if err != nil {
		vlog.Warn(err)
	}

	return session, nil
}

////////////////////////////////////////////////////////////////////////////////

// DefaultSession ...
type DefaultSession struct {
	context    *keeper.SessionContext
	adapter    keeper.SessionAdapter
	props      collection.Properties
	wantCommit bool
	countTran  int
}

func (inst *DefaultSession) _Impl() keeper.Session {
	return inst
}

// GetContext ...
func (inst *DefaultSession) GetContext() context.Context {
	return inst.context.Context
}

// func (inst *DefaultSession) Attributes() collection.Attributes {
// 	return inst.atts
// }

// Properties ...
func (inst *DefaultSession) Properties() collection.Properties {
	return inst.props
}

// BeginTransaction ...
func (inst *DefaultSession) BeginTransaction() keeper.SessionTransaction {
	inst.countTran++
	t := &DefaultSessionTran{session: inst}
	return t
}

func (inst *DefaultSession) flush() error {
	count := inst.countTran
	want := inst.wantCommit
	if count == 0 && want {
		return inst.adapter.Store(inst)
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// DefaultSessionTran ...
type DefaultSessionTran struct {
	session *DefaultSession
}

func (inst *DefaultSessionTran) _Impl() keeper.SessionTransaction {
	return inst
}

// Commit ...
func (inst *DefaultSessionTran) Commit() error {
	inst.session.wantCommit = true
	return nil
}

// Close ...
func (inst *DefaultSessionTran) Close() error {
	inst.session.countTran--
	return inst.session.flush()
}

////////////////////////////////////////////////////////////////////////////////
