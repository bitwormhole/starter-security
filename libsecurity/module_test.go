package libsecurity

import (
	"testing"

	"github.com/bitwormhole/starter/tests"
)

func TestSecurity(t *testing.T) {

	myMod := Module()

	rt, _ := tests.Starter(t).UsePanic().Use(myMod).RunEx()

	rt.Loop()

}
