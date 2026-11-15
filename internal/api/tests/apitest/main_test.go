package apitest

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
)

var testSrv api.Server

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)

	err := config.LoadEnvFromRoot()
	if err != nil {
		panic(err)
	}

	err = testSrv.Init()
	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()
	_ = config.Cfg.DBConn.Close()
	os.Exit(code)
}
