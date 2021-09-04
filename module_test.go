package startersecurity

import (
	"testing"

	"github.com/bitwormhole/starter/tests"
)

func TestSecurity(t *testing.T) {

	myMod := Module()

	rt, _ := tests.TestingStarter(t).UsePanic().Use(myMod).RunEx()

	rt.Loop()

}
