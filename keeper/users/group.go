package users

import "strings"

// Group 表示一个用户组
type Group string

func (value Group) String() string {
	return "group:" + string(value)
}

// QName 取用户组的规范名称 (Qualified Name)
func (value Group) QName() string {
	return "group:" + string(value)
}

// Normalize 标准化
func (value Group) Normalize() Group {
	v := string(value)
	v = strings.ToLower(v)
	v = strings.TrimSpace(v)
	return Group(v)
}
