package config

import (
	"github.com/digineo/go-uci"
	"github.com/servicemngr/core/package/instance"
	smerrors "github.com/servicemngr/servicemngr-aux/package/errors"
	consts "github.com/servicemngr/servicemngr-svcs-base/package/services/config/consts"
	filesystem "github.com/servicemngr/servicemngr-svcs-base/package/services/filesystem/interface"
	"github.com/spf13/afero"
	"strconv"
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

/*
In theory an instance could call this method with an suspicious argument.
Currently this is not dangerous because all instances that have access to an filesystem instance have the same rights
*/
func (c *Config) GetDependentInstance(_ instance.ID) instance.DependencyInstance {
	return c
}

func (c *Config) GetSections(secType string) ([]string, bool, error) {
	if c.uci == nil {
		return nil, false, smerrors.DataNotReadyError
	}
	a, b := c.uci.GetSections(c.id.Instance, secType)
	return a, b, nil
}

func (c *Config) Get(section, option string) ([]string, bool, error) {
	if c.uci == nil {
		return nil, false, smerrors.DataNotReadyError
	}
	a, b := c.uci.Get(c.id.Instance, section, option)
	return a, b, nil
}

func (c *Config) GetLast(section, option string) (string, bool, error) {
	if c.uci == nil {
		return "", false, smerrors.DataNotReadyError
	}
	a, b := c.uci.GetLast(c.id.Instance, section, option)
	return a, b, nil
}

func (c *Config) GetBool(section, option string) (bool, bool, error) {
	if c.uci == nil {
		return false, false, smerrors.DataNotReadyError
	}
	a, b := c.uci.GetBool(c.id.Instance, section, option)
	return a, b, nil
}

func (c *Config) GetBoolDefault(section, option string, def bool) (bool, error) {
	if c.uci == nil {
		return false, smerrors.DataNotReadyError
	}
	if v, ok := c.uci.GetBool(c.id.Instance, section, option); ok {
		return v, nil
	} else {
		return def, nil
	}
}

func (c *Config) GetIntList(section, option string) ([]int, bool, error) {
	if c.uci == nil {
		return nil, false, smerrors.DataNotReadyError
	}
	if raw, ok, err := c.Get(section, option); ok && err != nil {
		var data []int
		for _, rv := range raw {
			if v, err := strconv.Atoi(rv); err == nil {
				data = append(data, v)
			} else {
				return []int{}, ok, err
			}
		}
		return data, ok, nil
	} else {
		return []int{}, ok, err
	}
}

func (c *Config) GetInt(section, option string) (int, bool, error) {
	if raw, ok, err := c.GetLast(section, option); ok && err != nil {
		if v, err := strconv.Atoi(raw); err == nil {
			return v, true, nil
		} else {
			return 0, true, err
		}
	} else {
		return 0, ok, err
	}
}

func (c *Config) GetIntDefault(section, option string, def int) (int, error) {
	if v, ok, err := c.GetInt(section, option); err == nil {
		if ok {
			return v, nil
		} else {
			return def, nil
		}
	} else {
		return 0, err
	}
}

func (c *Config) GetStringDefault(section, option string, def string) (string, error) {
	if v, ok, err := c.GetLast(section, option); err == nil {
		if ok {
			return v, nil
		} else {
			return def, nil
		}
	} else {
		return "", err
	}
}

func (c *Config) Set(section, option string, values ...string) (bool, error) {
	if c.uci == nil {
		return false, smerrors.DataNotReadyError
	}
	return c.uci.Set(c.id.Instance, section, option, values...), nil
}

func (c *Config) Del(section, option string) error {
	if c.uci == nil {
		return smerrors.DataNotReadyError
	}
	c.uci.Del(c.id.Instance, section, option)
	return nil
}

func (c *Config) AddSection(section, typ string) error {
	if c.uci == nil {
		return smerrors.DataNotReadyError
	}
	return c.uci.AddSection(c.id.Instance, section, typ)
}

func (c *Config) DelSection(section string) error {
	if c.uci == nil {
		return smerrors.DataNotReadyError
	}
	c.uci.DelSection(c.id.Instance, section)
	return nil
}
