package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getAgilityTier(r *http.Request, i handlerInput[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedApiResourceList], id int32) (AgilityTier, error) {
	agilityTier, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return AgilityTier{}, err
	}

	response := AgilityTier{
		ID:       		agilityTier.ID,
		FromAgility: 	agilityTier.MinAgility,
		ToAgility: 		agilityTier.MaxAgility,
		TickSpeed: 		agilityTier.TickSpeed,
		MonMinICV: 		agilityTier.MonsterMinICV,
		MonMaxICV: 		agilityTier.MonsterMaxICV,
		CharMaxICV: 	agilityTier.CharacterMaxICV,
		CharMinICVs: 	convertObjSlice(cfg, agilityTier.CharacterMinICVs, convertAgilitySubtier),
	}

	return response, nil
}

func (cfg *Config) retrieveAgilityTiers(r *http.Request, i handlerInput[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedApiResourceList]) (UnnamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[UnnamedAPIResource]{
		frl(intQuery(cfg, r, i, resources, "agility", cfg.db.GetAgilityTierIDsByAgility)),
	})
}
