// Package filepathh provides function to walk over a directory tree and manipulate file/path names.
// It is based on original go package path/filepath, but have differences.
package filepathh

import "os"

// ExtractExt returns the file name extension used by path and remaining part of path.
// The extension is the suffix beginning at the final dot (including dot itself) in the final element of path; it is empty if there is no dot.
func ExtractExt(path string) (rem, ext string) {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[:i], path[i:]
		}
	}
	return path, ""
}
