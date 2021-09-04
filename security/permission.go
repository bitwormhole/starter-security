package security

// Permission 许可
type Permission interface {
	Implies(p Permission) bool
}
