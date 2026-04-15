package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getStat(r *http.Request, i handlerInput[seeding.Stat, Stat, NamedAPIResource, NamedApiResourceList], id int32) (Stat, error) {
	stat, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Stat{}, err
	}

	rel, err := getStatRelationships(cfg, r, stat)
	if err != nil {
		return Stat{}, err
	}

	response := Stat{
		ID:                	stat.ID,
		Name:               stat.Name,
		Effect: 			stat.Effect,
		MinVal: 			stat.MinVal,
		MaxVal: 			stat.MaxVal,
		MaxVal2: 			stat.MaxVal2,
		ActivationSphere: 	nameToNamedAPIResource(cfg, cfg.e.spheres, stat.ActivationSphere, nil),
		Spheres: 			rel.Spheres,
		AutoAbilities: 		rel.AutoAbilities,
		PlayerAbilities: 	rel.PlayerAbilities,
		OverdriveAbilities: rel.OverdriveAbilities,
		ItemAbilities: 		rel.ItemAbilities,
		TriggerCommands: 	rel.TriggerCommands,
		StatusConditions: 	rel.StatusConditions,
		Properties: 		rel.Properties,
	}

	return response, nil
}


func (cfg *Config) retrieveStats(r *http.Request, i handlerInput[seeding.Stat, Stat, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return i.resToListFunc(cfg, r, resources)
}
