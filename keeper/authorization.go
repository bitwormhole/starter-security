package keeper

// Authorization 授权请求
type Authorization interface {

	// GetSecurityAccess() SecurityAccess

	SecurityAccess

	// GetAccess() Access
	// GetIdentity() Identity

	// GetRoles() users.Roles
	// HasRole(role users.Role) bool
	// AddRole(role users.Role)
	// RemoveRole(role users.Role)
	// CleanRoles()
}

////////////////////////////////////////////////////////////////////////////////

// type SimpleAuthorization struct {
// 	Access SecurityAccess
// }

// func (inst *SimpleAuthorization) _Impl() Authorization {
// 	return inst
// }

// func (inst *SimpleAuthorization) GetSecurityAccess() SecurityAccess {
// 	return inst.Access
// }
