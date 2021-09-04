package security

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

// SubjectContext Subject 上下文
type SubjectContext interface {
	collection.Attributes

	GetSecurityManager() SecurityManager

	SetSecurityManager(securityManager SecurityManager)

	ResolveSecurityManager() SecurityManager

	GetSessionID() lang.Serializable

	SetSessionID(id lang.Serializable)

	GetSubject() Subject

	SetSubject(subject Subject)

	GetPrincipals() PrincipalCollection

	ResolvePrincipals() PrincipalCollection

	SetPrincipals(principals PrincipalCollection)

	GetSession() Session

	SetSession(session Session)

	ResolveSession() Session

	IsAuthenticated() bool

	SetAuthenticated(authc bool)

	IsSessionCreationEnabled() bool

	SetSessionCreationEnabled(enabled bool)

	ResolveAuthenticated() bool

	GetAuthenticationInfo() AuthenticationInfo

	SetAuthenticationInfo(info AuthenticationInfo)

	GetAuthenticationToken() AuthenticationToken

	SetAuthenticationToken(token AuthenticationToken)

	GetHost() string

	SetHost(host string)

	ResolveHost() string
}
