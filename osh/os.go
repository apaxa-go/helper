// Package osh provides some helpers for system os package.
package osh

import (
	"os"
	"time"
)

// ModTime returns modification timestamp for file named by fileName.
// In case of errors (such as no such file) ModTime returns error.
func ModTime(fileName string) (modTime time.Time, err error) {
	var stat os.FileInfo
	stat, err = os.Stat(fileName)
	if err == nil {
		modTime = stat.ModTime()
	}
	return
}

// Exists check if file exists.
// It does not check for type of file (regular, directory, ...).
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
