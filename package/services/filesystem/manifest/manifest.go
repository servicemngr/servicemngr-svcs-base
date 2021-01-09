package filesystem_manifest

import (
	"github.com/servicemngr/core/package/manifest"
	"github.com/servicemngr/servicemngr-svcs-base/package/services/filesystem"
	consts "github.com/servicemngr/servicemngr-svcs-base/package/services/filesystem/consts"
)

var Manifest = manifest.ServiceManifest{
	Name:             consts.SERVICE_NAME,
	Instantiable:     false,
	InstanceInitFunc: filesystem.Init,
}
