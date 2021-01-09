package web_interface

import (
	"github.com/servicemngr/servicemngr-svcs-base/package/services/web"
)

type Web interface {
	RegisterInstanceAPI(instance web.WebInstance)
	UnregisterInstanceAPI()
}
