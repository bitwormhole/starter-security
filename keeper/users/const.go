package users

// 定义一些常用的的角色
const (
	RoleAdmin     Role = "admin"
	RoleAnonymous Role = "anonymous"
	RoleAnyone    Role = "anyone"
	RoleAuthe     Role = "authe" // 已认证的
	RoleFriend    Role = "friend"
	RoleManager   Role = "manager"
	RoleNormal    Role = "normal"
	RoleOwner     Role = "owner"
	RoleRoot      Role = "root"
)
