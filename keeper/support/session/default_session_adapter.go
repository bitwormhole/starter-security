package session

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/vlog"
)

////////////////////////////////////////////////////////////////////////////////

// NewDefaultSessionAdapter 新建一个默认的会话适配器
func NewDefaultSessionAdapter(c context.Context) keeper.SessionAdapter {
	return &DefaultSessionAdapter{context: c}
}

// DefaultSessionAdapter ...
type DefaultSessionAdapter struct {
	context context.Context
}

func (inst *DefaultSessionAdapter) _Impl() keeper.SessionAdapter {
	return inst
}

// GetContext ...
func (inst *DefaultSessionAdapter) GetContext() context.Context {
	return inst.context
}

// Load ...
func (inst *DefaultSessionAdapter) Load(s keeper.Session) error {
	vlog.Warn("NOP for keeper.SessionAdapter.Load()")
	return nil
}

// Store ...
func (inst *DefaultSessionAdapter) Store(s keeper.Session) error {
	vlog.Warn("NOP for keeper.SessionAdapter.Store()")
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// DefaultSessionAdapterFactory ...
type DefaultSessionAdapterFactory struct {
}

func (inst *DefaultSessionAdapterFactory) _Impl() keeper.SessionAdapterFactory {
	return inst
}

// Create ...
func (inst *DefaultSessionAdapterFactory) Create(ctx context.Context) (keeper.SessionAdapter, error) {
	adapter := &DefaultSessionAdapter{
		context: ctx,
	}
	return adapter, nil
}

////////////////////////////////////////////////////////////////////////////////
