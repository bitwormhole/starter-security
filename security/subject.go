package security

import "github.com/bitwormhole/starter/lang"

// Subject 主体，可以是任何可以与应用交互的 “用户”
type Subject interface {

	//todo...

	GetPrincipal() lang.Object

	GetPrincipals() PrincipalCollection

	IsPermittedS(permission string) bool

	IsPermittedP(permission Permission) bool

	IsPermittedSA(permissions ...string) []bool

	IsPermittedPA(permissions []Permission) []bool

	IsPermittedAllSA(permissions ...string) bool

	IsPermittedAllPA(permissions []Permission) bool

	CheckPermissionS(permission string) error

	CheckPermissionP(permission Permission) error

	CheckPermissionsSA(permissions ...string) error

	CheckPermissionsPA(permissions []Permission) error

	HasRole(roleIdentifier string) bool

	HasRoles(roleIdentifiers []string) []bool

	HasAllRoles(roleIdentifiers []string) bool

	CheckRole(roleIdentifier string) error

	CheckRoles(roleIdentifiers ...string) error

	Login(token AuthenticationToken) error

	IsAuthenticated() bool

	IsRemembered() bool

	GetSession(create bool) Session

	Logout()

	//  <V> V execute(Callable<V> callable) throws ExecutionException;

	//  void execute(Runnable runnable);

	//  <V> Callable<V> associateWith(Callable<V> callable);

	//  Runnable associateWith(Runnable runnable);

	//  void runAs(PrincipalCollection principals) throws NullPointerException, IllegalStateException;

	IsRunAs() bool

	GetPreviousPrincipals() PrincipalCollection

	ReleaseRunAs() PrincipalCollection
}
