package api

import (


    "github.com/joho/godotenv"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func LoadEnvFromRoot() error {
    envPath, err := h.GetAbsoluteFilepath(".env")
    if err != nil {
        return err
    }

    return godotenv.Load(envPath)
}