package users

import (
	"strconv"
	"strings"
)

// UserID 表示一个用户ID
type UserID int64

// UserName 表示一个用户名
type UserName string

// UserUUID 表示一个用户的UUID
type UserUUID string

////////////////////////////////////////////////////////////////////////////////

func (value UserUUID) String() string {
	return string(value)
}

////////////////////////////////////////////////////////////////////////////////

func (value UserID) String() string {
	n := int64(value)
	return strconv.FormatInt(n, 10)
}

////////////////////////////////////////////////////////////////////////////////

// QName 取用户的规范名称 (Qualified Name)
func (value UserName) QName() string {
	return "user:" + string(value)
}

func (value UserName) String() string {
	return string(value)
}

// Normalize 标准化
func (value UserName) Normalize() UserName {
	v := string(value)
	v = strings.ToLower(v)
	v = strings.TrimSpace(v)
	return UserName(v)
}

////////////////////////////////////////////////////////////////////////////////
