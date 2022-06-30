package configs

import (
	"path/filepath"
	"runtime"
)

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/examples/data directory in the
// user's GOPATH.  If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	//nolint:dogsled
	_, currentFile, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(currentFile)

	return filepath.Join(basepath, rel)
}
