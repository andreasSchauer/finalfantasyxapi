package apitest

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
)

var testSrv api.Server

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)

	err := api.LoadEnvFromRoot()
	if err != nil {
		panic(err)
	}

	err = testSrv.Init()
	if err != nil {
		log.Fatal(err)
	}
	testSrv.InitGlobals()

	code := m.Run()
	_ = testSrv.Cfg.DBConn.Close()
	os.Exit(code)
}
