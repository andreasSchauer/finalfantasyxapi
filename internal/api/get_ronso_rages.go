package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (srv *Server) getRonsoRage(r *http.Request, i handlerInput[seedtypes.RonsoRage, RonsoRage, NamedAPIResource], id int32) (RonsoRage, error) {
	ronsoRage, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return RonsoRage{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(srv, r.Context(), ronsoRage, config.Cfg.DB.GetRonsoRageMonsterIDs)
	if err != nil {
		return RonsoRage{}, err
	}

	response := RonsoRage{
		ID:          ronsoRage.ID,
		Name:        ronsoRage.Name,
		Description: ronsoRage.Description,
		Effect:      ronsoRage.Effect,
		Overdrive:   nameToNamedAPIResource(srv, E.Overdrives, ronsoRage.Name, nil),
		Monsters:    monsters,
	}

	return response, nil
}

func (srv *Server) retrieveRonsoRages(r *http.Request, i handlerInput[seedtypes.RonsoRage, RonsoRage, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
