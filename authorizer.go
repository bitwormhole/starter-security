package security

// Authorizer 授权者
type Authorizer interface {
	CheckPermissionS(subjectPrincipal PrincipalCollection, permission string) error

	CheckPermissionP(subjectPrincipal PrincipalCollection, permission Permission) error

	CheckPermissionsSA(subjectPrincipal PrincipalCollection, permissions ...string) error

	CheckPermissionsPA(subjectPrincipal PrincipalCollection, permissions []Permission) error

	CheckRole(subjectPrincipal PrincipalCollection, roleIdentifier string) error

	CheckRoles(subjectPrincipal PrincipalCollection, roleIdentifiers ...string) error

	HasRole(subjectPrincipal PrincipalCollection, roleIdentifier string) bool

	HasRoles(subjectPrincipal PrincipalCollection, roleIdentifiers []string) []bool

	HasAllRoles(subjectPrincipal PrincipalCollection, roleIdentifiers []string) bool

	IsPermittedS(principals PrincipalCollection, permission string) bool

	IsPermittedP(subjectPrincipal PrincipalCollection, permission Permission) bool

	IsPermittedSA(subjectPrincipal PrincipalCollection, permissions ...string) bool

	IsPermittedPA(subjectPrincipal PrincipalCollection, permissions []Permission) []bool

	IsPermittedAllSA(subjectPrincipal PrincipalCollection, permissions ...string) bool

	IsPermittedAllPA(subjectPrincipal PrincipalCollection, permissions []Permission) bool
}
