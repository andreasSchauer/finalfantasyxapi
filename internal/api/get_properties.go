package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getProperty(r *http.Request, i handlerInput[seedtypes.Property, apitypes.Property, apiresource.NamedAPIResource], id int32) (apitypes.Property, error) {
	property, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.Property{}, err
	}

	rel, err := getPropertyRelationships(r, property)
	if err != nil {
		return apitypes.Property{}, err
	}

	response := apitypes.Property{
		ID:             property.ID,
		Name:           property.Name,
		Effect:         property.Effect,
		NullifyArmored: property.NullifyArmored,
		RelatedStats:   namesToNamedAPIResources(E.Stats, property.RelatedStats),
		ModifierChange: ah.ConvertObjPtr(property.ModifierChange, convertModifierChange),
		AutoAbilities:  rel.AutoAbilities,
		Monsters:       rel.Monsters,
	}

	return response, nil
}

func retrieveProperties(r *http.Request, i handlerInput[seedtypes.Property, apitypes.Property, apiresource.NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
