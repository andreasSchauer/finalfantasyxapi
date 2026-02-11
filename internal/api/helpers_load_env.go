package api

import (
    "path/filepath"

    "github.com/joho/godotenv"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func LoadEnvFromRoot() error {
    root, err := h.ProjectRoot()
    if err != nil {
        return err
    }
    envPath := filepath.Join(root, ".env")
    return godotenv.Load(envPath)
}