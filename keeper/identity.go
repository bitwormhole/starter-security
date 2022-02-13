package keeper

import "github.com/bitwormhole/starter-security/keeper/users"

// Identity 身份
type Identity interface {
	UserID() users.UserID
	UserName() users.UserName
	UserUUID() users.UserUUID
	Nickname() string
	Avatar() string
	Roles() users.Roles
}

// UserGroup 用户组
type UserGroup interface {
	GroupName() users.Group
	Contains(user Identity) bool
}

////////////////////////////////////////////////////////////////////////////////

// IdentityBuilder 用来创建一个简单的身份信息
type IdentityBuilder struct {
	UserID   users.UserID
	UserName users.UserName
	UserUUID users.UserUUID
	Nickname string
	Avatar   string
	Roles    users.Roles
}

// Identity 创建一个简单的身份信息
func (inst *IdentityBuilder) Identity() Identity {
	ident := &simpleIdentity{
		userid:   inst.UserID,
		username: inst.UserName,
		uuid:     inst.UserUUID,
		nickname: inst.Nickname,
		avatar:   inst.Avatar,
		roles:    inst.Roles,
	}
	return ident
}

////////////////////////////////////////////////////////////////////////////////

type simpleIdentity struct {
	userid   users.UserID
	username users.UserName
	uuid     users.UserUUID
	nickname string
	avatar   string
	roles    users.Roles
}

func (inst *simpleIdentity) UserID() users.UserID {
	return inst.userid
}

func (inst *simpleIdentity) UserName() users.UserName {
	return inst.username
}

func (inst *simpleIdentity) UserUUID() users.UserUUID {
	return inst.uuid
}

func (inst *simpleIdentity) Nickname() string {
	return inst.nickname
}

func (inst *simpleIdentity) Avatar() string {
	return inst.avatar
}

func (inst *simpleIdentity) Roles() users.Roles {
	return inst.roles
}

////////////////////////////////////////////////////////////////////////////////
