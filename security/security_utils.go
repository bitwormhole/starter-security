package security

import "context"

// GetSubject 从给定的上下文取 Subject 对象
func GetSubject(ctx context.Context) Subject {

	//todo...
	return nil
}

// GetSecurityManager 取安全管理器
func GetSecurityManager(ctx context.Context) (SecurityManager, error) {
	// SecurityUtils.securityManager = securityManager;
	return nil, nil
}

// SetSecurityManager 设置安全管理器
func SetSecurityManager(securityManager SecurityManager) {
	// SecurityUtils.securityManager = securityManager;
}
