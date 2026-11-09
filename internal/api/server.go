// Package api includes all the api routers, handlers, and their logic
package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
)

type Server struct{}

func (srv *Server) Init() error {
	err := config.Cfg.Init()
	if err != nil {
		return err
	}

	srv.InitGlobals()

	return nil
}

func (srv *Server) InitGlobals() {
	enums.T.Init()
	InitEndpoints(E)
	InitQueryLookup(Q)
}
