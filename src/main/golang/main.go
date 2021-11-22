package main

import (
	"github.com/bitwormhole/starter/application"
)

// //go:embed src/main/resources
// var resources embed.FS

func main() {
	// cb := config.NewBuilderFS(&resources, "src/main/resources")

	err := tryMain(nil)
	if err != nil {
		panic(err)
	}
}

func tryMain(cb application.ConfigBuilder) error {

	// err := demo.Config(cb)
	// if err != nil {
	// 	return err
	// }

	// context, err := application.Run(cb.Create(), os.Args)
	// if err != nil {
	// 	return err
	// }

	// err = demo.Run(context)
	// if err != nil {
	// 	return err
	// }

	// code, err := application.Exit(context)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println("exit with code:", code)

	return nil
}
