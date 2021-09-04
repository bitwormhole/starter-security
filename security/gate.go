package security

import "github.com/bitwormhole/starter/lang"

// Gate   安全闸门
type Gate interface {
	Control(ctx lang.Context) GateController
}

// GateFactory  安全闸门工厂
type GateFactory interface {
	Create() Gate
}

// GateController 安全闸门控制器
type GateController interface {
	Context() lang.Context
	Check() (lang.Context, error)

	UsePanic() GateController
	DisusePanic() GateController
}
