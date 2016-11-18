// Package ioutilh provides some helper functions and types to simplify working with ioutil package.
package ioutilh

import (
	"io"
	"os"
	"sort"
)

// ReadDirNames reads the directory named by dirName and returns list of directory entries.
// Result slice will be sorted if arrange is set to true.
func ReadDirNames(dirName string, arrange bool) ([]string, error) {
	f, err := os.Open(dirName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	names, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	if arrange {
		sort.Strings(names)
	}
	return names, nil
}

// IsDirEmpty checks if directory named by dirName is empty (has no child).
func IsDirEmpty(dirName string) (bool, error) {
	f, err := os.Open(dirName)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
