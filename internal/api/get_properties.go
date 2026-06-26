package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getProperty(r *http.Request, i handlerInput[seeding.Property, Property, NamedAPIResource, NamedApiResourceList], id int32) (Property, error) {
	property, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Property{}, err
	}

	rel, err := getPropertyRelationships(cfg, r, property)
	if err != nil {
		return Property{}, err
	}

	response := Property{
		ID:             property.ID,
		Name:           property.Name,
		Effect:         property.Effect,
		NullifyArmored: property.NullifyArmored,
		RelatedStats:   namesToNamedAPIResources(cfg, cfg.e.stats, property.RelatedStats),
		ModifierChange: convertObjPtr(cfg, property.ModifierChange, convertModifierChange),
		AutoAbilities:  rel.AutoAbilities,
		Monsters:       rel.Monsters,
	}

	return response, nil
}



func (cfg *Config) retrieveProperties(r *http.Request, i handlerInput[seeding.Property, Property, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	return verifyParamsAndRetrieve(r, i)
}
