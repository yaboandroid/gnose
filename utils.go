package gnose

import (
	"path/filepath"
	"os"
)

func GetCurrentDir() (path string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	path = dir
	return
}
