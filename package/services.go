package svcs_base

import (
	"github.com/servicemngr/core/package/manifest"
	config "github.com/servicemngr/servicemngr-svcs-base/package/services/config/manifest"
	filesystem "github.com/servicemngr/servicemngr-svcs-base/package/services/filesystem/manifest"
	web "github.com/servicemngr/servicemngr-svcs-base/package/services/web/manifest"
)

var Services = []manifest.ServiceManifest{
	config.Manifest,
	filesystem.Manifest,
	web.Manifest,
}
