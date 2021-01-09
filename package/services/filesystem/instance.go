package filesystem

import (
	"github.com/servicemngr/core/package/instance"
	"github.com/spf13/afero"
)

type FileSystem struct {
	afero.Fs
	id       instance.ID
	logger   instance.LoggerFunc
	error    instance.ErrorFunc
	selfkill instance.SelfKillFunc
}

func (f *FileSystem) Start() {
	f.Fs = afero.NewOsFs()
}

func (f *FileSystem) Stop() {
	f.Fs = nil
}

func (f *FileSystem) ID() instance.ID {
	return f.id
}

/*
In theory an instance could call this method with an suspicious argument.
Currently this is not dangerous because all instances that have access to an filesystem instance have the same rights
*/
func (f *FileSystem) GetDependentInstance(_ instance.ID) instance.DependencyInstance {
	return f
}
