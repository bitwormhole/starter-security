package permissions

type RoleManager struct {
	simpleRoles map[string]*Role
	cachedRoles map[string]*RoleList
}
