package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (cfg *Config) getAgilityTier(r *http.Request, i handlerInput[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource], id int32) (AgilityTier, error) {
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
		CharMinICVs: convertObjSlice(cfg, agilityTier.CharacterMinICVs, convertAgilitySubtier),
	}

	return response, nil
}

func (cfg *Config) retrieveAgilityTiers(r *http.Request, i handlerInput[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryInt(r, ids, QpnAgility, cfg.db.GetAgilityTierIDsByAgility),
	})
}
