package seeding

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getRawDataHash() (string, error) {
	dataDir, err := h.GetAbsoluteFilepath("data")
	if err != nil {
		return "", err
	}

	files, err := os.ReadDir(dataDir)
	if err != nil {
		return "", err
	}

	var combinedBytes bytes.Buffer

	for _, file := range files {
		if file.IsDir() || strings.HasPrefix(file.Name(), ".") {
			continue
		}

		filePath := filepath.Join(dataDir, file.Name())
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return "", err
		}

		combinedBytes.Write(fileContent)
	}

	hashBytes := sha256.Sum256(combinedBytes.Bytes())
	hash := hex.EncodeToString(hashBytes[:])

	return hash, nil
}