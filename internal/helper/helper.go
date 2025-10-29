package helper

import (
	"os"
	"path/filepath"
)

// InitPaths: Initializes the paths provided if they dont exist
func InitPaths(paths []string) error {
	for _, p := range paths {
		if filepath.Ext(p) == "" {
			if err := os.MkdirAll(p, 0777); err != nil {
				return err
			}
		} else {
			dir := filepath.Dir(p)
			if err := os.MkdirAll(dir, 0777); err != nil {
				return err
			}
			if _, err := os.Stat(p); os.IsNotExist(err) {
				f, err := os.Create(p)
				if err != nil {
					return err
				}
				f.Close()
			}
		}
	}
	return nil
}
