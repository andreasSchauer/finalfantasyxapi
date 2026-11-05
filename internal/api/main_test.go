package api

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
)

var testSrv Server

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
	testSrv.InitGlobals()

	code := m.Run()
	_ = config.Cfg.DBConn.Close()
	os.Exit(code)
}
