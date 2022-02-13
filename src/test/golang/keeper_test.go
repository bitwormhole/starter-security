package main

import (
	"context"
	"testing"

	startersecurity "github.com/bitwormhole/starter-security"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/support/session"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/contexts"
	"github.com/bitwormhole/starter/tests"
	"github.com/bitwormhole/starter/vlog"
)

func TestKeeper(t *testing.T) {
	err := doTestKeeper(t)
	if err != nil {
		t.Error(err)
	}
}

func doTestKeeper(t *testing.T) error {

	i := tests.Starter(t)
	i.Use(startersecurity.ModuleForTest())
	rt, err := i.RunEx()
	if err != nil {
		return err
	}

	ac := rt.Context()
	cc := &contexts.SimpleContext{}
	err = contexts.SetupContextSetter(cc)
	if err != nil {
		return err
	}

	err = doTestKeeper2(ac, cc)
	if err != nil {
		return err
	}

	return rt.Exit()
}

func doTestKeeper2(appCtx application.Context, cc context.Context) error {

	o1, err := appCtx.GetComponent("#keeper-subject-manager")
	if err != nil {
		return err
	}

	holder, err := keeper.GetHolder(cc)
	if err != nil {
		return err
	}
	ac := holder.GetAccessContext()
	ac.Adapter = session.NewDefaultSessionAdapter(cc)

	subjects := o1.(keeper.SubjectManager)
	subject, err := subjects.GetSubject(cc)
	if err != nil {
		return err
	}
	isauth := subject.IsAuthenticated()
	vlog.Debug(isauth)

	return nil
}
