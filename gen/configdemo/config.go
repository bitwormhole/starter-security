package configdemo

import "github.com/bitwormhole/starter/application"

// ExportConfigForKeeperDemo ...
func ExportConfigForKeeperDemo(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
	// return nil
}
