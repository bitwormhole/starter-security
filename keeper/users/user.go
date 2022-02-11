package users

import "strings"

// Name 表示一个用户名
type Name string

// QName 取用户的规范名称 (Qualified Name)
func (value Name) QName() string {
	return "user:" + string(value)
}

func (value Name) String() string {
	return string(value)
}

// Normalize 标准化
func (value Name) Normalize() Name {
	v := string(value)
	v = strings.ToLower(v)
	v = strings.TrimSpace(v)
	return Name(v)
}
