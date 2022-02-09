package configlib

import "github.com/bitwormhole/starter/application"

func ExportConfigForKeeperLib(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
