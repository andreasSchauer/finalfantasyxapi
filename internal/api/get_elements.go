package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getElement(r *http.Request, i handlerInput[seeding.Element, Element, NamedAPIResource, NamedApiResourceList], id int32) (Element, error) {
	element, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Element{}, err
	}

	rel, err := getElementRelationships(cfg, r, element)
	if err != nil {
		return Element{}, err
	}

	response := Element{
		ID:                	element.ID,
		Name:               element.Name,
		OppositeElement: 	namePtrToNamedAPIResPtr(cfg, i, element.OppositeElement, nil),
		StatusProtection: 	rel.StatusProtection,
		AutoAbilities: 		rel.AutoAbilities,
		PlayerAbilities: 	rel.PlayerAbilities,
		OverdriveAbilities: rel.OverdriveAbilities,
		ItemAbilities: 		rel.ItemAbilities,
		EnemyAbilities: 	rel.EnemyAbilities,
		MonstersWeak: 		rel.MonstersWeak,
		MonstersHalved: 	rel.MonstersHalved,
		MonstersImmune: 	rel.MonstersImmune,
		MonstersAbsorb: 	rel.MonstersAbsorb,
	}

	return response, nil
}

func (cfg *Config) retrieveElements(r *http.Request, i handlerInput[seeding.Element, Element, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return i.resToListFunc(cfg, r, resources)
}
