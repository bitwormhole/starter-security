package keeper

import (
	"context"

	"github.com/bitwormhole/starter/contexts"
)

// Holder 持有会话相关的对象
type Holder struct {
	target *SessionContext
}

// GetSessionContext 获取会话上下文，如果没有就新建一个
func (inst *Holder) GetSessionContext() *SessionContext {
	t := inst.target
	if t == nil {
		t = &SessionContext{}
		inst.target = t
	}
	return t
}

// GetHolder 获取会话持有者对象
func GetHolder(ctx context.Context) (*Holder, error) {
	const key = "github.com/bitwormhole/starter-security/keeper/Holder#binding"
	o1 := ctx.Value(key)
	o2, ok := o1.(*Holder)
	if ok {
		return o2, nil
	}
	setter, err := contexts.GetContextSetter(ctx)
	if err != nil {
		return nil, err
	}
	o2 = &Holder{}
	setter.SetValue(key, o2)
	return o2, nil
}
