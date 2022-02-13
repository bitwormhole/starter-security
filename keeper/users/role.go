package users

import (
	"sort"
	"strings"
)

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

// Equal 判断两个角色是否相等
func (value Role) Equal(v2 Role) bool {
	str1 := strings.TrimSpace(value.String())
	str2 := strings.TrimSpace(v2.String())
	return str1 == str2
}

// Normalize 标准化
func (value Role) Normalize() Role {
	v := string(value)
	v = strings.ToLower(v)
	v = strings.TrimSpace(v)
	return Role(v)
}

////////////////////////////////////////////////////////////////////////////////

// Roles 角色列表 (以“,”为分隔符)
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

// Normalize 标准化角色列表（排重+排序）
func (value Roles) Normalize() Roles {
	table := make(map[string]bool)
	list := make([]string, 0)
	src := strings.Split(value.String(), ",")
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		older := table[item]
		if older {
			continue
		}
		table[item] = true
		list = append(list, item)
	}
	sort.Strings(list)
	builder := strings.Builder{}
	for i, item := range list {
		if i > 0 {
			builder.WriteRune(',')
		}
		builder.WriteString(item)
	}
	return Roles(builder.String())
}

// Contains 判断列表中是否包含某个元素
func (value Roles) Contains(role Role) bool {
	list := value.List()
	for _, item := range list {
		if item.Equal(role) {
			return true
		}
	}
	return false
}

// Add 向列表添加一个元素
func (value Roles) Add(role Role) Roles {
	str1 := value.String()
	str2 := role.String()
	result := Roles(str1 + "," + str2)
	return result.Normalize()
}

// Remove 将指定的元素从列表移除
func (value Roles) Remove(role Role) Roles {
	list := value.List()
	builder := strings.Builder{}
	for _, item := range list {
		if item.Equal(role) {
			continue
		}
		builder.WriteString(item.String())
		builder.WriteRune(',')
	}
	result := Roles(builder.String())
	return result.Normalize()
}

////////////////////////////////////////////////////////////////////////////////
