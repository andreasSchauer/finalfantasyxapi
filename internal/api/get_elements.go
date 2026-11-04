package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (srv *Server) getElement(r *http.Request, i handlerInput[seedtypes.Element, Element, NamedAPIResource], id int32) (Element, error) {
	element, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Element{}, err
	}

	rel, err := getElementRelationships(srv, r, element)
	if err != nil {
		return Element{}, err
	}

	response := Element{
		ID:                 element.ID,
		Name:               element.Name,
		OppositeElement:    namePtrToNamedAPIResPtr(srv, i, element.OppositeElement, nil),
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

func (srv *Server) retrieveElements(r *http.Request, i handlerInput[seedtypes.Element, Element, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
