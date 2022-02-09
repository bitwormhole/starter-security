package keeper

// Configurer 用来配置keeper上下文
// 【inject:".keeper-configurer"】
type Configurer interface {
	Configure(c *Context) error
}
