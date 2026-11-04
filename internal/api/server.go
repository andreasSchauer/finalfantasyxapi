// Package api includes all the api routers, handlers, and their logic
package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
)

type Server struct {
	Cfg *config.Config
}

func (srv *Server) Init() error {
	srv.Cfg = &config.Config{}
	err := srv.Cfg.Init()
	if err != nil {
		return err
	}

	srv.InitGlobals()

	return nil
}

func (srv *Server) InitGlobals() {
	enums.T.Init()
	E.Init(srv)
	Q.Init()
}
