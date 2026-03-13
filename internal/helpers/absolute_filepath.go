package helpers

import (
	"fmt"
	"path/filepath"
	"os"
)

func GetAbsoluteFilepath(path string) (string, error) {
	root, err := projectRoot()
    if err != nil {
        return "", err
    }

	return filepath.Join(root, path), nil
}


func projectRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		_, err := os.Stat(filepath.Join(wd, "go.mod"))
		if err == nil {
			return wd, nil
		}

		parent := filepath.Dir(wd)
		if wd == parent {
			return "", fmt.Errorf("project root not found from %s", wd)
		}
		
		wd = parent
	}
}