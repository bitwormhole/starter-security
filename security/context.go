package security

import (
	"github.com/bitwormhole/starter/application"
)

type Context interface {
	application.SimpleContext

	Session() Session
}
