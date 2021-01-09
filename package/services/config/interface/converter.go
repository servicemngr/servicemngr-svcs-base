package config_interface

func Convert(v interface{}) (Config, bool) {
	c, ok := v.(Config)
	return c, ok
}
