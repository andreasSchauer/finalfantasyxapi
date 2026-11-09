package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetProperty(r *http.Request, i handlerinput.Main[seedtypes.Property, apitypes.Property, apiresource.Named], id int32) (apitypes.Property, error) {
	property, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Property{}, err
	}

	rel, err := assemble.GetPropertyRelationships(r, property)
	if err != nil {
		return apitypes.Property{}, err
	}

	response := apitypes.Property{
		ID:             property.ID,
		Name:           property.Name,
		Effect:         property.Effect,
		NullifyArmored: property.NullifyArmored,
		RelatedStats:   handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, property.RelatedStats),
		ModifierChange: ah.ConvertObjPtr(property.ModifierChange, apiconv.ConvertModifierChange),
		AutoAbilities:  rel.AutoAbilities,
		Monsters:       rel.Monsters,
	}

	return response, nil
}

func RetrieveProperties(r *http.Request, i handlerinput.Main[seedtypes.Property, apitypes.Property, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
