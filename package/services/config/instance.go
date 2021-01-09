package config

import (
	"github.com/digineo/go-uci"
	"github.com/servicemngr/core/package/instance"
	smerrors "github.com/servicemngr/servicemngr-aux/package/errors"
	consts "github.com/servicemngr/servicemngr-svcs-base/package/services/config/consts"
	filesystem "github.com/servicemngr/servicemngr-svcs-base/package/services/filesystem/interface"
	"github.com/spf13/afero"
)

type Config struct {
	id       instance.ID
	logger   instance.LoggerFunc
	error    instance.ErrorFunc
	selfkill instance.SelfKillFunc
	uci      uci.Tree
}

func (c *Config) Start() {}

func (c *Config) Stop() {}

func (c *Config) ID() instance.ID {
	return c.id
}

func (c *Config) Dependencies() instance.Dependencies {
	return []instance.ID{{
		Name: "filesystem",
	}}
}

func (c *Config) SetDependency(i instance.DependencyInstance) {
	if i.ID().Name == "filesystem" && i.ID().Instance == instance.NON_INSTANCE_NAME {
		fs, ok := filesystem.Convert(i)
		if ok {
			c.uci = uci.NewTreeFromFs(afero.NewBasePathFs(fs, consts.CONFIG_PATH))
		} else {
			c.error(smerrors.NewInstanceAssertionError(i))
		}
	}
}

func (c *Config) UnsetDependency(id instance.ID) {
	if id.Name == "filesystem" && id.Instance == instance.NON_INSTANCE_NAME {
		c.uci = nil
	}
}

func (c *Config) OnServiceListChanged() {}

func (c *Config) SetServiceListGetter(_ func() []string) {}

func (c *Config) SetDependenciesChangedHandler(_ func()) {}

func (c *Config) GetDependentInstance(id instance.ID) instance.DependencyInstance {
	return &DependencyConfig{
		config: c,
		id:     id,
	}
}
