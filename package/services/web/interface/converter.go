package web_interface

func Convert(v interface{}) (Web, bool) {
	web, ok := v.(Web)
	return web, ok
}
