package subject

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
	"github.com/bitwormhole/starter/vlog"
)

// DefaultSubject ...
type DefaultSubject struct {
	ac *keeper.AccessContext
}

func (inst *DefaultSubject) _Impl() keeper.Subject {
	return inst
}

// GetSession ...
func (inst *DefaultSubject) GetSession(create bool) (keeper.Session, error) {
	session := inst.ac.Session
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

	inst.ac.Session = session
	return session, nil
}

func (inst *DefaultSubject) loadSession() (keeper.Session, error) {
	ctx := inst.ac.Context
	sc := inst.ac.SecurityContext
	adapter := inst.ac.Adapter
	factory := sc.GetSessionProvider().GetSessionFactory()
	if adapter == nil {
		return nil, errors.New("the use of keeper.SessionAdapter without setup")
	}
	session, err := factory.Create(ctx, adapter)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// GetAccess ...
func (inst *DefaultSubject) GetAccess() keeper.Access {
	return inst.ac.Access
}

// GetContext ...
func (inst *DefaultSubject) GetContext() context.Context {
	return inst.ac.Context
}

// GetPermission ...
func (inst *DefaultSubject) GetPermission() keeper.Permission {
	return inst.ac.Permission
}

// SetPermission ...
func (inst *DefaultSubject) SetPermission(p keeper.Permission) {
	if p == nil {
		return
	}
	inst.ac.Permission = p
}

// IsAuthenticated ...
func (inst *DefaultSubject) IsAuthenticated() bool {
	const name = keeper.SessionFieldAuthenticated
	session, err := inst.GetSession(false)
	if err != nil {
		return false
	}
	getter := session.Properties().Getter()
	return getter.GetBool(name, false)
}

// SetSession ...
func (inst *DefaultSubject) SetSession(s keeper.Session) {
	if s == nil {
		return
	}
	inst.ac.Session = s
}

// SetAccess ...
func (inst *DefaultSubject) SetAccess(a keeper.Access) {
	if a == nil {
		return
	}
	inst.ac.Access = a
}

// SetAuthenticated ...
func (inst *DefaultSubject) SetAuthenticated(authenticated bool) {
	const name = keeper.SessionFieldAuthenticated
	session, err := inst.GetSession(true)
	if err != nil {
		panic(err)
	}
	setter := session.Properties().Setter()
	setter.SetBool(name, authenticated)
}

// Login 登录
func (inst *DefaultSubject) Login(ctx context.Context, a keeper.Authentication) (keeper.Identity, error) {

	sc := inst.ac.SecurityContext

	// Authenticate
	identity, err := sc.GetAuthentications().Authenticate(ctx, a)
	if err != nil {
		return nil, err
	}

	// init session
	err = inst.initSessionWithIdentity(identity)
	if err != nil {
		return nil, err
	}

	return identity, nil
}

func (inst *DefaultSubject) initSessionWithIdentity(identity keeper.Identity) error {

	session, err := inst.GetSession(true)
	if err != nil {
		return err
	}

	tr := session.BeginTransaction()
	defer func() {
		err = tr.Close()
		if err != nil {
			vlog.Warn(err)
		}
	}()

	session.SetIdentity(identity)
	setter := session.Properties().Setter()
	setter.SetBool(keeper.SessionFieldAuthenticated, true)

	return tr.Commit()
}

// Logout ...
func (inst *DefaultSubject) Logout(ctx context.Context) error {
	return errors.New("no impl")
}

// Authorize  授权
func (inst *DefaultSubject) Authorize(ctx context.Context) error {

	perm := inst.ac.Permission
	if perm == nil {
		access := inst.ac.Access
		pm := inst.ac.SecurityContext.GetPermissions()
		pt, err := pm.FindTemplate(ctx, access)
		if err != nil {
			return err
		}
		perm2, err := pt.LoadPermission(access.Params())
		if err != nil {
			return err
		}
		perm = perm2
		inst.ac.Permission = perm2
	}

	am := inst.ac.SecurityContext.GetAuthorizations()
	return am.Authorize(ctx)
}

// HasPermission 鉴权
func (inst *DefaultSubject) HasPermission(ctx context.Context) bool {

	const ok = true
	const forbidden = false

	perm := inst.ac.Permission
	if perm == nil {
		return forbidden
	}

	if perm.AcceptRole(users.RoleAnonymous) || perm.AcceptRole(users.RoleAnyone) {
		return ok
	}

	access := inst.ac.SecurityAccess
	roles := access.GetRoles()

	if perm.AcceptRoles(roles) {
		return ok
	}

	// 获取身份
	subject := access.GetSubject()
	session, err := subject.GetSession(true)
	if err != nil {
		return forbidden
	}

	roles = session.GetRoles()
	ident := session.GetIdentity()

	if perm.AcceptRoles(roles) {
		return ok
	}

	// check owner
	if perm.AcceptRole(users.RoleOwner) {
		if perm.IsOwner(ident) {
			return ok
		}
	}

	// check firend
	if perm.AcceptRole(users.RoleFriend) {
		if perm.IsFriend(ident) {
			return ok
		}
	}

	return forbidden
}
