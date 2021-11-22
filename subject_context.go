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

////////////////////////////////////////////////////////////////////////////////

type DefaultSubjectContext struct {
	atts collection.Attributes

	securityManager SecurityManager
	sessionID       lang.Serializable
	subject         Subject
	principals      PrincipalCollection
	authToken       AuthenticationToken
	authInfo        AuthenticationInfo
	session         Session

	sessionCreationEnabled bool
	authenticated          bool
	host                   string
}

func (inst *DefaultSubjectContext) _Impl() SubjectContext {
	return nil
}

func (inst *DefaultSubjectContext) resolve() error {
	return nil
}

func (inst *DefaultSubjectContext) GetSecurityManager() SecurityManager {
	return inst.securityManager
}

func (inst *DefaultSubjectContext) SetSecurityManager(securityManager SecurityManager) {
	inst.securityManager = securityManager
}

func (inst *DefaultSubjectContext) ResolveSecurityManager() SecurityManager {
	return inst.securityManager
}

func (inst *DefaultSubjectContext) GetSessionID() lang.Serializable {
	return inst.sessionID
}

func (inst *DefaultSubjectContext) SetSessionID(id lang.Serializable) {
	inst.sessionID = id
}

func (inst *DefaultSubjectContext) GetSubject() Subject {
	return inst.subject
}

func (inst *DefaultSubjectContext) SetSubject(subject Subject) {
	inst.subject = subject
}

func (inst *DefaultSubjectContext) GetPrincipals() PrincipalCollection {
	return inst.principals
}

func (inst *DefaultSubjectContext) ResolvePrincipals() PrincipalCollection {
	inst.resolve()
	return inst.principals
}

func (inst *DefaultSubjectContext) SetPrincipals(principals PrincipalCollection) {
	inst.principals = principals
}

func (inst *DefaultSubjectContext) GetSession() Session {
	return inst.session
}

func (inst *DefaultSubjectContext) SetSession(session Session) {
	inst.session = session
}

func (inst *DefaultSubjectContext) ResolveSession() Session {
	inst.resolve()
	return inst.session
}

func (inst *DefaultSubjectContext) IsAuthenticated() bool {
	return inst.authenticated
}

func (inst *DefaultSubjectContext) SetAuthenticated(authc bool) {
	inst.authenticated = authc
}

func (inst *DefaultSubjectContext) IsSessionCreationEnabled() bool {
	return inst.sessionCreationEnabled
}

func (inst *DefaultSubjectContext) SetSessionCreationEnabled(enabled bool) {
	inst.sessionCreationEnabled = enabled
}

func (inst *DefaultSubjectContext) ResolveAuthenticated() bool {
	inst.resolve()
	return inst.authenticated
}

func (inst *DefaultSubjectContext) GetAuthenticationInfo() AuthenticationInfo {
	return inst.authInfo
}

func (inst *DefaultSubjectContext) SetAuthenticationInfo(info AuthenticationInfo) {
	inst.authInfo = info
}

func (inst *DefaultSubjectContext) GetAuthenticationToken() AuthenticationToken {
	return inst.authToken
}

func (inst *DefaultSubjectContext) SetAuthenticationToken(token AuthenticationToken) {
	inst.authToken = token
}

func (inst *DefaultSubjectContext) GetHost() string {
	return inst.host
}

func (inst *DefaultSubjectContext) SetHost(host string) {
	inst.host = host
}

func (inst *DefaultSubjectContext) ResolveHost() string {
	inst.resolve()
	return inst.host
}
