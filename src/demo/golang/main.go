package main

import (
	"github.com/bitwormhole/starter"
	startersecurity "github.com/bitwormhole/starter-security"
)

// main of demo
func main() {
	i := starter.InitApp()
	i.Use(startersecurity.ModuleForDemo())
	i.Run()
}
