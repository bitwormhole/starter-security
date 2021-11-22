package common

// type SimpleAuthWrapper struct {
// 	Mechanism string
// 	Inner     security.AuthenticationProvider
// }

// func (inst *SimpleAuthWrapper) _Impl() security.AuthenticationProvider {
// 	return inst
// }

// func (inst *SimpleAuthWrapper) Name() string {
// 	return inst.Mechanism
// }

// func (inst *SimpleAuthWrapper) Supports(t security.AuthenticationToken) bool {
// 	mech1 := strings.ToLower(t.Mechanism())
// 	mech2 := strings.ToLower(inst.Mechanism)
// 	if mech1 != mech2 {
// 		return false
// 	}
// 	return true
// }

// func (inst *SimpleAuthWrapper) Authenticate(t auths.AuthenticationToken) (auths.AuthenticationInfo, error) {
// 	return inst.Inner.Authenticate(t)
// }
