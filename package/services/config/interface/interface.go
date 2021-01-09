package config_interface

type Config interface {
	GetSections(secType string) ([]string, bool, error)
	Get(section, option string) ([]string, bool, error)
	GetLast(section, option string) (string, bool, error)
	GetBool(section, option string) (bool, bool, error)
	GetBoolDefault(section, option string, def bool) (bool, error)
	GetIntList(section, option string) ([]int, bool, error)
	GetInt(section, option string) (int, bool, error)
	GetIntDefault(section, option string, def int) (int, error)
	GetStringDefault(section, option string, def string) (string, error)
	Set(section, option string, values ...string) (bool, error)
	Del(section, option string) error
	AddSection(section, typ string) error
	DelSection(section string) error
}
