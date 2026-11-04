package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (s *Server) getElement(r *http.Request, i handlerInput[seedtypes.Element, Element, NamedAPIResource], id int32) (Element, error) {
	element, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Element{}, err
	}

	rel, err := getElementRelationships(s, r, element)
	if err != nil {
		return Element{}, err
	}

	response := Element{
		ID:                 element.ID,
		Name:               element.Name,
		OppositeElement:    namePtrToNamedAPIResPtr(s, i, element.OppositeElement, nil),
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

func (s *Server) retrieveElements(r *http.Request, i handlerInput[seedtypes.Element, Element, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
