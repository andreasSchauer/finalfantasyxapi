package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getProperty(r *http.Request, i handlerInput[seedtypes.Property, Property, NamedAPIResource], id int32) (Property, error) {
	property, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Property{}, err
	}

	rel, err := getPropertyRelationships(r, property)
	if err != nil {
		return Property{}, err
	}

	response := Property{
		ID:             property.ID,
		Name:           property.Name,
		Effect:         property.Effect,
		NullifyArmored: property.NullifyArmored,
		RelatedStats:   namesToNamedAPIResources(E.Stats, property.RelatedStats),
		ModifierChange: convertObjPtr(property.ModifierChange, convertModifierChange),
		AutoAbilities:  rel.AutoAbilities,
		Monsters:       rel.Monsters,
	}

	return response, nil
}

func retrieveProperties(r *http.Request, i handlerInput[seedtypes.Property, Property, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
