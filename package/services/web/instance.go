package web

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/servicemngr/core/package/instance"
	"net/http"
	"sync"
)

type Web struct {
	id        instance.ID
	logger    instance.LoggerFunc
	error     instance.ErrorFunc
	selfkill  instance.SelfKillFunc
	echo      *echo.Echo
	instances map[instance.ID]WebInstance
	lock      sync.Mutex
	wg        sync.WaitGroup
}

func (w *Web) Start() {
	w.instances = make(map[instance.ID]WebInstance)
	w.echo = echo.New()
	w.echo.Logger.SetLevel(log.OFF)
	w.echo.HideBanner = true
	w.echo.HidePort = true
	w.start()
}

func (w *Web) Stop() {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.stop()
}

func (w *Web) ID() instance.ID {
	return w.id
}

func (w *Web) GetDependentInstance(id instance.ID) instance.DependencyInstance {
	return &DependencyWeb{
		web: w,
		id:  id,
	}
}

func (w *Web) start() {
	w.wg.Add(1)
	go func() {
		if err := w.echo.Start(":8080"); err != nil {
			w.wg.Done()
			if err != http.ErrServerClosed {
				w.error(err)
			} else {
				return
			}
		}
	}()
}

func (w *Web) stop() {
	if err := w.echo.Shutdown(context.TODO()); err != nil {
		w.error(err)
	}
	w.wg.Wait()
}

func (w *Web) restart() {
	w.stop()
	w.start()
	for sn, svc := range w.instances {
		svc.WebRegisterRoutes(w.getServiceGroup(sn))
	}
}

func (w *Web) getServiceGroup(id instance.ID) *echo.Group {
	name := id.Name
	if id.Instance != instance.NON_INSTANCE_NAME {
		name = name + "/" + id.Instance
	}
	return w.echo.Group("/api/"+name, middleware.CORS())
}
