package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getRonsoRage(r *http.Request, i handlerInput[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedApiResourceList], id int32) (RonsoRage, error) {
	ronsoRage, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return RonsoRage{}, err
	}

	monsters, err := getResourcesDbItem(cfg, r, cfg.e.monsters, ronsoRage, cfg.db.GetRonsoRageMonsterIDs)
	if err != nil {
		return RonsoRage{}, err
	}

	response := RonsoRage{
		ID:          ronsoRage.ID,
		Name:        ronsoRage.Name,
		Description: ronsoRage.Description,
		Effect:      ronsoRage.Effect,
		Overdrive:   nameToNamedAPIResource(cfg, cfg.e.overdrives, ronsoRage.Name, nil),
		Monsters:    monsters,
	}

	return response, nil
}

func (cfg *Config) retrieveRonsoRages(r *http.Request, i handlerInput[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return i.resToListFunc(cfg, r, resources)
}
