package configtest

import "github.com/bitwormhole/starter/application"

func ExportConfigForKeeperTest(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
