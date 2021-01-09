package config

import (
	"github.com/servicemngr/core/package/instance"
	smerrors "github.com/servicemngr/servicemngr-aux/package/errors"
	"strconv"
)

type DependencyConfig struct {
	config *Config
	id     instance.ID
}

func (d DependencyConfig) ID() instance.ID {
	return d.config.id
}

func (d *DependencyConfig) GetSections(secType string) ([]string, bool, error) {
	if d.config.uci == nil {
		return nil, false, smerrors.DataNotReadyError
	}
	a, b := d.config.uci.GetSections(d.id.String(), secType)
	return a, b, nil
}

func (d *DependencyConfig) Get(section, option string) ([]string, bool, error) {
	if d.config.uci == nil {
		return nil, false, smerrors.DataNotReadyError
	}
	a, b := d.config.uci.Get(d.id.String(), section, option)
	return a, b, nil
}

func (d *DependencyConfig) GetLast(section, option string) (string, bool, error) {
	if d.config.uci == nil {
		return "", false, smerrors.DataNotReadyError
	}
	a, b := d.config.uci.GetLast(d.id.String(), section, option)
	return a, b, nil
}

func (d *DependencyConfig) GetBool(section, option string) (bool, bool, error) {
	if d.config.uci == nil {
		return false, false, smerrors.DataNotReadyError
	}
	a, b := d.config.uci.GetBool(d.id.String(), section, option)
	return a, b, nil
}

func (d *DependencyConfig) GetBoolDefault(section, option string, def bool) (bool, error) {
	if d.config.uci == nil {
		return false, smerrors.DataNotReadyError
	}
	if v, ok := d.config.uci.GetBool(d.id.String(), section, option); ok {
		return v, nil
	} else {
		return def, nil
	}
}

func (d *DependencyConfig) GetIntList(section, option string) ([]int, bool, error) {
	if d.config.uci == nil {
		return nil, false, smerrors.DataNotReadyError
	}
	if raw, ok, err := d.Get(section, option); ok && err != nil {
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

func (d *DependencyConfig) GetInt(section, option string) (int, bool, error) {
	if raw, ok, err := d.GetLast(section, option); ok && err != nil {
		if v, err := strconv.Atoi(raw); err == nil {
			return v, true, nil
		} else {
			return 0, true, err
		}
	} else {
		return 0, ok, err
	}
}

func (d *DependencyConfig) GetIntDefault(section, option string, def int) (int, error) {
	if v, ok, err := d.GetInt(section, option); err == nil {
		if ok {
			return v, nil
		} else {
			return def, nil
		}
	} else {
		return 0, err
	}
}

func (d *DependencyConfig) GetStringDefault(section, option string, def string) (string, error) {
	if v, ok, err := d.GetLast(section, option); err == nil {
		if ok {
			return v, nil
		} else {
			return def, nil
		}
	} else {
		return "", err
	}
}

func (d *DependencyConfig) Set(section, option string, values ...string) (bool, error) {
	if d.config.uci == nil {
		return false, smerrors.DataNotReadyError
	}
	return d.config.uci.Set(d.id.String(), section, option, values...), nil
}

func (d *DependencyConfig) Del(section, option string) error {
	if d.config.uci == nil {
		return smerrors.DataNotReadyError
	}
	d.config.uci.Del(d.id.String(), section, option)
	return nil
}

func (d *DependencyConfig) AddSection(section, typ string) error {
	if d.config.uci == nil {
		return smerrors.DataNotReadyError
	}
	return d.config.uci.AddSection(d.id.String(), section, typ)
}

func (d *DependencyConfig) DelSection(section string) error {
	if d.config.uci == nil {
		return smerrors.DataNotReadyError
	}
	d.config.uci.DelSection(d.id.String(), section)
	return nil
}
