package permissions

type Permission struct {
	Method   string
	Resource string
	Accept   bool
}

type Permissions struct {
	All   []*Permission
	cache map[string]*Permission
}
