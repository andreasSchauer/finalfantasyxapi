package api

import (
	// "filepath"
	"io"
	"log"
	"os"
	"testing"
)

var testCfg *Config

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)

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
