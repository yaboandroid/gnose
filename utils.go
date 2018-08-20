package gnose

import (
	"path/filepath"
	"os"
	"strings"
)

func GetCurrentDir() (path string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	path = dir
	return
}

func CheckFileExist(filename string) (isExist bool) {
	isExist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		isExist = false
	}
	return
}

func CreateFolder(base, folder string) (err error) {
	separator := GetSystemSeparator()
	s1 := make([]string, 3)
	s1 = append(s1, base, separator, folder)
	fp := strings.Join(s1, "")
	if !CheckFileExist(fp) {
		err = os.MkdirAll(fp, os.ModePerm)
	}
	return
}

func GetSystemSeparator() (separator string) {
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}
	return
}
