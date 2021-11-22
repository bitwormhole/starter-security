package security

import (
	"context"
)

// Gate   安全闸门
type Gate interface {
	Control(ctx context.Context) GateController
}

// GateFactory  安全闸门工厂
type GateFactory interface {
	Create() Gate
}

// GateController 安全闸门控制器
type GateController interface {
	Context() context.Context
	Check() (context.Context, error)

	UsePanic() GateController
	DisusePanic() GateController
}
