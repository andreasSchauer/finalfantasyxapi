// Package api includes all the api routers, handlers, and their logic
package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
)

type Server struct {
	Cfg *config.Config
}

func (s *Server) Init() error {
	var err error

	s.Cfg, err = config.ConfigInit()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) InitGlobals() {
	enums.T.EnumLookupInit()
	E.EndpointsInit(s)
	Q.Init()
}
