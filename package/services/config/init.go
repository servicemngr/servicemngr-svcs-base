package config

import (
	"github.com/servicemngr/core/package/instance"
	"github.com/servicemngr/core/package/manifest"
)

func Init(id instance.ID, logger instance.LoggerFunc, error instance.ErrorFunc, selfkill instance.SelfKillFunc, _ manifest.Configurator) instance.Instance {
	return &Config{
		id:       id,
		logger:   logger,
		error:    error,
		selfkill: selfkill,
	}
}
