package security

import "github.com/bitwormhole/starter/lang"

// type Principal struct {
// }

// PrincipalCollection Principal 的集合
type PrincipalCollection interface {
	GetPrimaryPrincipal() lang.Object

	AsList() []lang.Object

	FromRealm(realmName string) []lang.Object

	GetRealmNames() []string

	IsEmpty() bool
}
