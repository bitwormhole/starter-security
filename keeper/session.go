package keeper

import (
	"context"
	"io"

	"github.com/bitwormhole/starter/collection"
)

// Session 会话
type Session interface {
	GetContext() context.Context

	// // 不可持久化的属性
	// Attributes() collection.Attributes

	// 可持久化的属性
	Properties() collection.Properties

	BeginTransaction() SessionTransaction
}

// SessionContext 会话上下文
type SessionContext struct {
	Access          Access
	Adapter         SessionAdapter
	Context         context.Context
	SecurityContext SecurityContext
	Session         Session
	Subject         Subject
}

// SessionTransaction 表示一个会话的事务
type SessionTransaction interface {
	io.Closer
	Commit() error
}

// SessionFactory 会话工厂
type SessionFactory interface {
	Create(adapter SessionAdapter) (Session, error)
}

// SessionAdapter 会话适配器
type SessionAdapter interface {
	GetContext() context.Context
	Load(s Session) error
	Store(s Session) error
}

// SessionAdapterFactory 会话适配器工厂
type SessionAdapterFactory interface {
	Create(ctx context.Context) (SessionAdapter, error)
}

// SessionLoader 会话加载器
type SessionLoader interface {
	Load(data []byte) (Session, error)
}

// SessionSerializer 会话存储器
type SessionSerializer interface {
	Serialize(s Session) ([]byte, error)
}

// SessionProvider 会话提供商
type SessionProvider interface {
	GetSessionFactory() SessionFactory
	GetAdapterFactory() SessionAdapterFactory
}

// SessionProviderRegistration 会话提供商注册项
type SessionProviderRegistration struct {
	Name     string
	Provider SessionProvider
}

// SessionProviderRegistry 会话提供商注册器
// 【inject:".keeper-session-provider-registry"】
type SessionProviderRegistry interface {
	GetRegistrationList() []*SessionProviderRegistration
}
