package web

import (
	"errors"
	"github.com/servicemngr/core/package/instance"
)

type DependencyWeb struct {
	web *Web
	id  instance.ID
}

func (d DependencyWeb) ID() instance.ID {
	return d.web.id
}

func (d *DependencyWeb) RegisterInstanceAPI(wi WebInstance) {
	if _, ok := d.web.instances[d.id]; ok {
		d.web.error(errors.New("\"" + d.id.Name + "/" + d.id.Instance + "\" was registered twice"))
	}
	d.web.lock.Lock()
	defer d.web.lock.Unlock()
	d.web.instances[d.id] = wi
	wi.WebRegisterRoutes(d.web.getServiceGroup(d.id))
}

func (d *DependencyWeb) UnregisterInstanceAPI() {
	if _, ok := d.web.instances[d.id]; !ok {
		d.web.error(errors.New("\"" + d.id.Name + "/" + d.id.Instance + "\" is not registered"))
	}
	d.web.lock.Lock()
	defer d.web.lock.Unlock()
	delete(d.web.instances, d.id)
	d.web.restart()
}
