package api

import (
	// "filepath"
	"os"
	"testing"

)

var testCfg *Config

func TestMain(m *testing.M) {
	// will be used once the file is in ./internal/api
	// _ = godotenv.Load(filepath.Join("..", "..", ".env"))
	err := LoadEnvFromRoot()
	if err != nil {
		panic(err)
	}

	cfg, err := ConfigInit()
	if err != nil {
		panic(err)
	}

	testCfg = cfg

	code := m.Run()
	_ = testCfg.dbConn.Close()
	os.Exit(code)
}