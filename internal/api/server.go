package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/enums"

type Server struct {
	Cfg		*Config
}

func (s *Server) Init() error {
	var err error

	s.Cfg, err = ConfigInit()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) InitGlobals() {
	enums.T.EnumLookupInit()
	E.EndpointsInit(s.Cfg)
	Q.Init()
}