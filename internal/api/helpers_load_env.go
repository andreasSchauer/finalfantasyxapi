package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/joho/godotenv"
)

func LoadEnvFromRoot() error {
	envPath, err := h.GetAbsoluteFilepath(".env")
	if err != nil {
		return err
	}

	return godotenv.Load(envPath)
}
