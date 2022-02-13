package session

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
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
func (inst *DefaultSessionFactory) Create(ctx context.Context, adapter keeper.SessionAdapter) (keeper.Session, error) {

	holder, err := keeper.GetHolder(ctx)
	if err != nil {
		return nil, err
	}

	sc := holder.GetAccessContext()

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
	context    *keeper.AccessContext
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

// IsAuthenticated ...
func (inst *DefaultSession) IsAuthenticated() bool {
	const key = keeper.SessionFieldAuthenticated
	return inst.Properties().Getter().GetBool(key, false)
}

// GetRoles ...
func (inst *DefaultSession) GetRoles() users.Roles {
	const key = keeper.SessionFieldRoles
	value := inst.Properties().GetProperty(key, "")
	return users.Roles(value)
}

// GetIdentity ...
func (inst *DefaultSession) GetIdentity() keeper.Identity {

	getter := inst.Properties().Getter()
	ib := keeper.IdentityBuilder{}

	userid := getter.GetInt64(keeper.SessionFieldUserID, 0)
	roles := getter.GetString(keeper.SessionFieldRoles, "")
	username := getter.GetString(keeper.SessionFieldUserName, "")
	uuid := getter.GetString(keeper.SessionFieldUserUUID, "")

	ib.Avatar = getter.GetString(keeper.SessionFieldAvatar, "")
	ib.Nickname = getter.GetString(keeper.SessionFieldDisplayName, "")
	ib.Roles = users.Roles(roles)

	ib.UserID = users.UserID(userid)
	ib.UserName = users.UserName(username)
	ib.UserUUID = users.UserUUID(uuid)

	return ib.Identity()
}

// SetIdentity ...
func (inst *DefaultSession) SetIdentity(ident keeper.Identity) {
	if ident == nil {
		return
	}
	setter := inst.Properties().Setter()

	setter.SetString(keeper.SessionFieldAvatar, ident.Avatar())
	setter.SetString(keeper.SessionFieldDisplayName, ident.Nickname())
	setter.SetString(keeper.SessionFieldRoles, ident.Roles().String())

	setter.SetString(keeper.SessionFieldUserID, ident.UserID().String())
	setter.SetString(keeper.SessionFieldUserName, ident.UserName().String())
	setter.SetString(keeper.SessionFieldUserUUID, ident.UserUUID().String())
}

// SetRoles ...
func (inst *DefaultSession) SetRoles(r users.Roles) {
	const key = keeper.SessionFieldRoles
	value := r.String()
	inst.Properties().SetProperty(key, value)
}

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
