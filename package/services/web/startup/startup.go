package web_startup

import (
	"github.com/servicemngr/core/package/instance"
	"github.com/servicemngr/core/package/manifest"
)

func Startup(configurator manifest.Configurator) []instance.ID {
	/*
		Proposal:
		Via configuration: Set web hosting directories. If set, start instance automatically.
	*/
	return []instance.ID{}
}
