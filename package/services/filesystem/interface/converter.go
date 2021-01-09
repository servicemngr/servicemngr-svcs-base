package filesystem_interface

func Convert(v interface{}) (FileSystem, bool) {
	fs, ok := v.(FileSystem)
	return fs, ok
}
