package users

import "strings"

////////////////////////////////////////////////////////////////////////////////

// Role 用户角色
type Role string

func (value Role) String() string {
	return string(value)
}

// QName 取角色的规范名称 (Qualified Name)
func (value Role) QName() string {
	return "role:" + string(value)
}

// Normalize 标准化
func (value Role) Normalize() Role {
	v := string(value)
	v = strings.ToLower(v)
	v = strings.TrimSpace(v)
	return Role(v)
}

////////////////////////////////////////////////////////////////////////////////

// Roles 角色列表
type Roles string

func (value Roles) String() string {
	return string(value)
}

// List 取角色列表
func (value Roles) List() []Role {
	str := string(value)
	src := parseItems(str)
	dst := make([]Role, len(src))
	for i, item := range src {
		dst[i] = Role(item)
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////
