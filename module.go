package startersecurity

import (
	"embed"

	"github.com/bitwormhole/starter-security/gen/configdemo"
	"github.com/bitwormhole/starter-security/gen/configlib"
	"github.com/bitwormhole/starter-security/gen/configtest"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myName     = "github.com/bitwormhole/starter-security"
	myVersion  = "v0.1.5"
	myRevision = 12
)

////////////////////////////////////////////////////////////////////////////////

//go:embed src/main/resources
var theResFS embed.FS

// Module 导出本模块
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myName).Version(myVersion).Revision(myRevision)
	mb.Resources(collection.LoadEmbedResources(&theResFS, "src/main/resources"))
	mb.OnMount(configlib.ExportConfigForKeeperLib)

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

//go:embed src/demo/resources
var theSrcDemoRes embed.FS

// ModuleForDemo  导出本模块
func ModuleForDemo() application.Module {

	parent := Module()
	mb := &application.ModuleBuilder{}

	mb.Name(parent.GetName() + "#demo")
	mb.Version(parent.GetVersion())
	mb.Revision(parent.GetRevision())

	mb.Resources(collection.LoadEmbedResources(&theSrcDemoRes, "src/demo/resources"))
	mb.OnMount(configdemo.ExportConfigForKeeperDemo)
	mb.Dependency(parent)

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

//go:embed src/test/resources
var theSrcTestRes embed.FS

// ModuleForTest 导出本模块
func ModuleForTest() application.Module {

	parent := Module()
	mb := &application.ModuleBuilder{}

	mb.Name(parent.GetName() + "#test")
	mb.Version(parent.GetVersion())
	mb.Revision(parent.GetRevision())

	mb.Resources(collection.LoadEmbedResources(&theSrcTestRes, "src/test/resources"))
	mb.OnMount(configtest.ExportConfigForKeeperTest)
	mb.Dependency(parent)

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////
