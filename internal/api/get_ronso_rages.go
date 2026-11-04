package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (s *Server) getRonsoRage(r *http.Request, i handlerInput[seedtypes.RonsoRage, RonsoRage, NamedAPIResource], id int32) (RonsoRage, error) {
	ronsoRage, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return RonsoRage{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(s, r.Context(), ronsoRage, s.Cfg.DB.GetRonsoRageMonsterIDs)
	if err != nil {
		return RonsoRage{}, err
	}

	response := RonsoRage{
		ID:          ronsoRage.ID,
		Name:        ronsoRage.Name,
		Description: ronsoRage.Description,
		Effect:      ronsoRage.Effect,
		Overdrive:   nameToNamedAPIResource(s, E.Overdrives, ronsoRage.Name, nil),
		Monsters:    monsters,
	}

	return response, nil
}

func (s *Server) retrieveRonsoRages(r *http.Request, i handlerInput[seedtypes.RonsoRage, RonsoRage, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
