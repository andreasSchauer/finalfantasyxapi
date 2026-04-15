package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getProperty(r *http.Request, i handlerInput[seeding.Property, Property, NamedAPIResource, NamedApiResourceList], id int32) (Property, error) {
	property, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Property{}, err
	}

	autoAbilities, err := getResourcesDbItem(cfg, r, cfg.e.autoAbilities, property, cfg.db.GetPropertyAutoAbilityIDs)
	if err != nil {
		return Property{}, err
	}

	monsters, err := getResourcesDbItem(cfg, r, cfg.e.monsters, property,cfg.db.GetPropertyMonsterIDs)
	if err != nil {
		return Property{}, err
	}

	response := Property{
		ID:                	property.ID,
		Name:               property.Name,
		Effect: 			property.Effect,
		NullifyArmored: 	property.NullifyArmored,
		RelatedStats: 		namesToNamedAPIResources(cfg, cfg.e.stats, property.RelatedStats),
		StatChanges: 		convertObjSlice(cfg, property.StatChanges, convertStatChange),
		ModifierChanges: 	convertObjSlice(cfg, property.ModifierChanges, convertModifierChange),
		AutoAbilities: 		autoAbilities,
		Monsters: 			monsters,
	}

	return response, nil
}

func (cfg *Config) retrieveProperties(r *http.Request, i handlerInput[seeding.Property, Property, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return i.resToListFunc(cfg, r, resources)
}
