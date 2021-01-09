package web_startup

import (
	"github.com/servicemngr/core/package/instance"
	"github.com/servicemngr/core/package/manifest"
	consts "github.com/servicemngr/servicemngr-svcs-base/package/services/web/consts"
)

func Startup(configurator manifest.Configurator) []instance.ID {
	return []instance.ID{{
		Name:     consts.SERVICE_NAME,
		Instance: instance.NON_INSTANCE_NAME,
	}}
}
