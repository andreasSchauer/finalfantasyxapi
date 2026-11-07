package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAgilityTier(r *http.Request, i handlerinput.HandlerInput[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.UnnamedAPIResource], id int32) (apitypes.AgilityTier, error) {
	agilityTier, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.AgilityTier{}, err
	}

	response := apitypes.AgilityTier{
		ID:          agilityTier.ID,
		FromAgility: agilityTier.MinAgility,
		ToAgility:   agilityTier.MaxAgility,
		TickSpeed:   agilityTier.TickSpeed,
		MonMinICV:   agilityTier.MonsterMinICV,
		MonMaxICV:   agilityTier.MonsterMaxICV,
		CharMaxICV:  agilityTier.CharacterMaxICV,
		CharMinICVs: ah.ConvertObjSlice(agilityTier.CharacterMinICVs, convertAgilitySubtier),
	}

	return response, nil
}

func retrieveAgilityTiers(r *http.Request, i handlerinput.HandlerInput[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryInt(r, ids, ts.QpnAgility, config.Cfg.DB.GetAgilityTierIDsByAgility),
	})
}
