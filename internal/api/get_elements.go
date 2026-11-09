package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getElement(r *http.Request, i handlerinput.Main[seedtypes.Element, apitypes.Element, apiresource.Named], id int32) (apitypes.Element, error) {
	element, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Element{}, err
	}

	rel, err := assemble.GetElementRelationships(r, element)
	if err != nil {
		return apitypes.Element{}, err
	}

	response := apitypes.Element{
		ID:                 element.ID,
		Name:               element.Name,
		OppositeElement:    handlerinput.NamePtrToNamedAPIResPtr(i, element.OppositeElement, nil),
		StatusProtection:   rel.StatusProtection,
		AutoAbilities:      rel.AutoAbilities,
		PlayerAbilities:    rel.PlayerAbilities,
		OverdriveAbilities: rel.OverdriveAbilities,
		ItemAbilities:      rel.ItemAbilities,
		EnemyAbilities:     rel.EnemyAbilities,
		MonstersWeak:       rel.MonstersWeak,
		MonstersHalved:     rel.MonstersHalved,
		MonstersImmune:     rel.MonstersImmune,
		MonstersAbsorb:     rel.MonstersAbsorb,
	}

	return response, nil
}

func retrieveElements(r *http.Request, i handlerinput.Main[seedtypes.Element, apitypes.Element, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
