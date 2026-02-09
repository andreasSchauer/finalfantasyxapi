package seeding

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func loadJSONFile[T any](path string, target *T) error {
	root, err := projectRoot()
	if err != nil {
		return err
	}
	fullPath := filepath.Join(root, path)
	
	file, err := os.Open(fullPath)
	if err != nil {
		return fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("couldn't read file: %v", err)
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("couldn't parse JSON: %v", err)
	}

	return nil
}


func projectRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		fp := filepath.Join(wd, "go.mod")
		if _, err := os.Stat(fp); err == nil {
			return wd, nil
		}
		parent := filepath.Dir(wd)
		if wd == parent {
			return "", fmt.Errorf("project root not found from %s", wd)
		}
		wd = parent
	}
}