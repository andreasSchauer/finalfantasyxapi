package api

import (
	"io"
	"log"
	"os"
	"testing"
)

var testSrv Server

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)

	err := LoadEnvFromRoot()
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
