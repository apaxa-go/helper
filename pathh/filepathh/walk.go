package filepathh

import (
	"github.com/apaxa-go/helper/ioh/ioutilh"
	"os"
	"path/filepath"
)

// walk recursively descends path, calling walkFn for each element.
func walk(path string, info os.FileInfo, walkFn filepath.WalkFunc, applyToRoot bool, sort bool) error {
	if info.IsDir() {
		if names, err := ioutilh.ReadDirNames(path, sort); err != nil {
			if err = walkFn(path, info, err); err != nil {
				return err
			}
		} else {
			for _, name := range names {
				filename := filepath.Join(path, name)
				fileInfo, err := os.Lstat(filename)
				if err != nil {
					if err = walkFn(filename, nil, err); err != nil {
						return err
					}
				} else {
					if err := walk(filename, fileInfo, walkFn, true, sort); err != nil {
						return err
					}
				}
			}
		}
	}

	if applyToRoot {
		return walkFn(path, info, nil)
	}
	return nil
}

// Walk walks the file tree rooted at root, calling walkFn for each file or directory in the tree. walkFn apply to root itself only if applyToRoot set to true.
// Walk walks from down to up, so it is possible to remove all files in some folder and than remove folder itself.
// All errors that arise visiting files and directories are filtered by walkFn. walkFn will not be raised for filter if error happens at root itself and applyToRoot is false. In case of error the info parameter for walkFn may be null.
// If sort is true files in directory are walked in lexical order, which makes the output deterministic but means that for very large directories it can be inefficient.
// Walk does not follow symbolic links.
// Warning: there is no special meaning for filepath.SkipDir returning from walkFunc (Walk treats SkipDir as any other errors and stop walking).
func Walk(root string, walkFn filepath.WalkFunc, applyToRoot bool, sort bool) error {
	info, err := os.Lstat(root)
	if err != nil {
		if applyToRoot {
			return walkFn(root, nil, err)
		}
		return err
	}
	return walk(root, info, walkFn, applyToRoot, sort)
}
