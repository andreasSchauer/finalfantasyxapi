package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getAgilityTier(r *http.Request, i handlerInput[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource], id int32) (AgilityTier, error) {
	agilityTier, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return AgilityTier{}, err
	}

	response := AgilityTier{
		ID:          agilityTier.ID,
		FromAgility: agilityTier.MinAgility,
		ToAgility:   agilityTier.MaxAgility,
		TickSpeed:   agilityTier.TickSpeed,
		MonMinICV:   agilityTier.MonsterMinICV,
		MonMaxICV:   agilityTier.MonsterMaxICV,
		CharMaxICV:  agilityTier.CharacterMaxICV,
		CharMinICVs: convertObjSlice(srv, agilityTier.CharacterMinICVs, convertAgilitySubtier),
	}

	return response, nil
}

func (srv *Server) retrieveAgilityTiers(r *http.Request, i handlerInput[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryInt(r, ids, ts.QpnAgility, config.Cfg.DB.GetAgilityTierIDsByAgility),
	})
}
