package main

import (
	"os"
	"testing"
)

var testCfg *Config

func TestMain(m *testing.M) {
	cfg, err := ConfigInit()
	if err != nil {
		panic(err)
	}

	testCfg = &cfg

	code := m.Run()
	_ = testCfg.dbConn.Close()
	os.Exit(code)
}