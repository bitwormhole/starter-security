package golang

import (
	"testing"

	"github.com/bitwormhole/starter/lang"
)

func TestLogin(t *testing.T) {
	err := tryTestLogin(t)
	if err != nil {
		t.Error(err)
	}
}

func tryTestLogin(t *testing.T) error {

	ctx := &lang.SimpleContext{}

	_, err := lang.GetContext(ctx)

	return err
}
