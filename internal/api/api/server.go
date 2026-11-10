// Package api includes all the api routers, handlers, and their logic
package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/initlookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
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
	enums.InitEnums()
	initlookup.Endpoints()
	initlookup.Queries()
}
