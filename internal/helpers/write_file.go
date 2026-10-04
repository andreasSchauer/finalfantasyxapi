package helpers

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
)

func ConvertJsonData[O, N any](fileName string, old []O, convFn func(O) N) error {
	return WriteFile("new_files", fileName, convertStructs(old, convFn))
}

func convertStructs[O, N any](old []O, convFn func(O) N) []N {
	new := make([]N, len(old))

	for i := range old {
		new[i] = convFn(old[i])
	}

	return new
}

func WriteFile[T any](dirName, fileName string, data []T) error {
	fileDir, err := GetAbsoluteFilepath(dirName)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fileDir, 0755)
	if err != nil {
		return err
	}

	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	jsonBytes = bytes.ReplaceAll(jsonBytes, []byte(`\u0026`), []byte("&"))
	jsonBytes = bytes.ReplaceAll(jsonBytes, []byte(`\u003c`), []byte("<"))
	jsonBytes = bytes.ReplaceAll(jsonBytes, []byte(`\u003e`), []byte(">"))

	destPath := filepath.Join(fileDir, fileName)

	err = os.WriteFile(destPath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
