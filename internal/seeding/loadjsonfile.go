package seeding

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func loadJSONFile[T any](srcPath string, target *T) error {
	file, err := os.Open(srcPath)
	
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
