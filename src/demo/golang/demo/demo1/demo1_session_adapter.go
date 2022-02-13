package demo1

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/vlog"
)

type Demo1sessionAdapter struct {
	context context.Context
}

func (inst *Demo1sessionAdapter) _Impl() keeper.SessionAdapter {
	return inst
}

func (inst *Demo1sessionAdapter) GetContext() context.Context {
	return inst.context
}

func (inst *Demo1sessionAdapter) Store(s keeper.Session) error {
	p := s.Properties()
	str := collection.FormatProperties(p)
	vlog.Warn("Demo1sessionAdapter.Store(session): ", str)
	return nil
}

func (inst *Demo1sessionAdapter) Load(s keeper.Session) error {
	return nil
}
