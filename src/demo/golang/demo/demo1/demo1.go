package demo1

import (
	"errors"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/contexts"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// Demo1 ...
type Demo1 struct {
	markup.Component `class:"life"`

	AC       application.Context   `inject:"context"`
	Subjects keeper.SubjectManager `inject:"#keeper-subject-manager"`
}

func (inst *Demo1) _Impl() application.LifeRegistry {
	return inst
}

// GetLifeRegistration ...
func (inst *Demo1) GetLifeRegistration() *application.LifeRegistration {
	lr := &application.LifeRegistration{}
	lr.Looper = inst
	lr.OnStart = inst.initContext
	return lr
}

func (inst *Demo1) initContext() error {

	ctx := inst.AC
	err := contexts.SetupApplicationContext(inst.AC)
	if err != nil {
		return err
	}

	h, err := keeper.GetHolder(ctx)
	if err != nil {
		return err
	}

	ac := h.GetAccessContext()
	adapter := &Demo1sessionAdapter{context: ctx}
	ac.Adapter = adapter
	ac.Access = configAccess()
	ac.SecurityAccess = (&keeper.DefaultSecurityAccess{}).Init(ac)

	return err
}

// Loop ...
func (inst *Demo1) Loop() error {

	err := inst.doLogin()
	if err != nil {
		return err
	}

	return inst.doCheckAuth()
}

func (inst *Demo1) doLogin() error {

	ctx := inst.AC
	subject, err := inst.Subjects.GetSubject(ctx)
	if err != nil {
		return err
	}

	loginParams := &myLoginParams{
		username: "demo", password: "demo",
	}

	ident, err := subject.Login(ctx, loginParams)
	if err != nil {
		return err
	}

	vlog.Debug(ident)
	return nil

}

func (inst *Demo1) doCheckAuth() error {

	ctx := inst.AC
	subject, err := inst.Subjects.GetSubject(ctx)
	if err != nil {
		return err
	}

	err = subject.Authorize(ctx)
	if err != nil {
		return err
	}

	ok := subject.HasPermission(ctx)
	if !ok {
		return errors.New("no permission")
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type myLoginParams struct {
	username string
	password string
}

func (inst *myLoginParams) _Impl() keeper.Authentication {
	return inst
}

func (inst *myLoginParams) Mechanism() string {
	return "password"
}

func (inst *myLoginParams) User() string {
	return inst.username
}

func (inst *myLoginParams) Secret() []byte {
	return []byte(inst.password)
}

////////////////////////////////////////////////////////////////////////////////
