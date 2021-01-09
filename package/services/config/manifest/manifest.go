package config_manifest

import (
	"github.com/servicemngr/core/package/manifest"
	"github.com/servicemngr/servicemngr-svcs-base/package/services/config"
	consts "github.com/servicemngr/servicemngr-svcs-base/package/services/config/consts"
)

var Manifest = manifest.ServiceManifest{
	Name:             consts.SERVICE_NAME,
	Instantiable:     false,
	InstanceInitFunc: config.Init,
}
