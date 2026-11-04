package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadJSONFile[T any](path string, target *[]T) error {
	fullPath, err := GetAbsoluteFilepath(path)
	if err != nil {
		return err
	}

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