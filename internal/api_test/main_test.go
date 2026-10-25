package apiTest

import (
	// "filepath"
	"io"
	"log"
	"os"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
)

var testCfg *api.Config

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)

	err := api.LoadEnvFromRoot()
	if err != nil {
		panic(err)
	}

	cfg, err := api.ConfigInit()
	if err != nil {
		panic(err)
	}

	testCfg = cfg

	code := m.Run()
	_ = testCfg.DbConn.Close()
	os.Exit(code)
}
