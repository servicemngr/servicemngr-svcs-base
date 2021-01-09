package web_manifest

import (
	"github.com/servicemngr/core/package/manifest"
	"github.com/servicemngr/servicemngr-svcs-base/package/services/web"
	consts "github.com/servicemngr/servicemngr-svcs-base/package/services/web/consts"
	startup "github.com/servicemngr/servicemngr-svcs-base/package/services/web/startup"
)

var Manifest = manifest.ServiceManifest{
	Name:                    consts.SERVICE_NAME,
	Instantiable:            false,
	GetStartupInstancesFunc: startup.Startup,
	InstanceInitFunc:        web.Init,
}
