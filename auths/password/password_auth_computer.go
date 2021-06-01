package password

import (
	"crypto/sha256"
	"strings"

	"github.com/bitwormhole/starter/util"
)

type AuthComputer struct {
	Times    int
	Salt     string
	User     string
	Password string
}

func (inst *AuthComputer) Init() {
	inst.Times = 5
}

func (inst *AuthComputer) Compute() string {
	value := inst.Password
	for cnt := 0; ; {
		value = inst.computeOnce(value)
		if cnt < inst.Times {
			cnt++
		} else {
			break
		}
	}
	return value
}

func (inst *AuthComputer) computeOnce(value string) string {
	builder := &strings.Builder{}
	builder.WriteString(inst.User)
	builder.WriteString("[")
	builder.WriteString(value)
	builder.WriteString("]")
	builder.WriteString(inst.Salt)
	return inst.sha256sum(builder.String())
}

func (inst *AuthComputer) sha256sum(input string) string {
	sum := sha256.Sum256([]byte(input))
	return util.StringifyBytes(sum[:])
}
