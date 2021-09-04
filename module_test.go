package startersecurity

import (
	"testing"

	"github.com/bitwormhole/starter/tests"
)

func TestSecurity(t *testing.T) {

	rt, _ := tests.TestingStarter(t).UsePanic().RunEx()

	rt.Loop()

}
