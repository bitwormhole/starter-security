package startersecurity

import (
	"embed"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myName     = "github.com/bitwormhole/starter-security"
	myVersion  = "v0.0.3"
	myRevision = 3
)

//go:embed src/main/resources
var theResFS embed.FS

// Module 导出本模块
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myName).Version(myVersion).Revision(myRevision)
	mb.Resources(collection.LoadEmbedResources(&theResFS, "src/main/resources"))

	return mb.Create()
}
