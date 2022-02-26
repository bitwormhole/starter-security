package keeper

import "github.com/bitwormhole/starter-security/keeper/users"

// Identity 身份
type Identity interface {
	Avatar() string
	Email() string
	Nickname() string
	Roles() users.Roles
	UserID() users.UserID
	UserName() users.UserName
	UserUUID() users.UserUUID
}

// UserGroup 用户组
type UserGroup interface {
	GroupName() users.Group
	Contains(user Identity) bool
}

////////////////////////////////////////////////////////////////////////////////

// IdentityBuilder 用来创建一个简单的身份信息
type IdentityBuilder struct {
	Avatar   string
	Email    string
	Nickname string
	Roles    users.Roles
	UserID   users.UserID
	UserName users.UserName
	UserUUID users.UserUUID
}

// Identity 创建一个简单的身份信息
func (inst *IdentityBuilder) Identity() Identity {
	ident := &simpleIdentity{
		avatar:   inst.Avatar,
		email:    inst.Email,
		nickname: inst.Nickname,
		roles:    inst.Roles,
		userid:   inst.UserID,
		username: inst.UserName,
		uuid:     inst.UserUUID,
	}
	return ident
}

////////////////////////////////////////////////////////////////////////////////

type simpleIdentity struct {
	avatar   string
	email    string
	nickname string
	roles    users.Roles
	userid   users.UserID
	username users.UserName
	uuid     users.UserUUID
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

func (inst *simpleIdentity) Email() string {
	return inst.email
}

////////////////////////////////////////////////////////////////////////////////
