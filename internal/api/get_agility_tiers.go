package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getAgilityTier(r *http.Request, i handlerInput[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedApiResourceList], id int32) (AgilityTier, error) {
	agilityTier, err := verifyParamsAndGet(r, i, id)
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

func (cfg *Config) retrieveAgilityTiers(r *http.Request, i handlerInput[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		intQuery(r, i, ids, qpnAgility, cfg.db.GetAgilityTierIDsByAgility),
	})
}
