package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetSphere(r *http.Request, i handlerinput.Main[seedtypes.Sphere, apitypes.Sphere, apiresource.Named], id int32) (apitypes.Sphere, error) {
	sphere, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Sphere{}, err
	}

	rel, err := getSphereRelationships(r, sphere)
	if err != nil {
		return apitypes.Sphere{}, err
	}

	response := apitypes.Sphere{
		ID:                 sphere.ID,
		Name:               sphere.Name,
		Item:               rel.Item,
		Description:        rel.Description,
		Effect:             rel.Effect,
		SgDescription:      sphere.SgDescription,
		SphereColor:        sphere.SphereColor,
		SphereEffect:       sphere.SphereEffect,
		TargetNodePosition: sphere.TargetNodePosition,
		TargetNodeState:    sphere.TargetNodeState,
		TargetableNodes:    sphere.TargetableNodes,
		CreatedNode:        ah.ConvertObjPtr(sphere.CreatedNode, apiconv.ConvertCreatedNode),
		Monsters:           rel.Monsters,
		Treasures:          rel.Treasures,
		Shops:              rel.Shops,
		Quests:             rel.Quests,
		BlitzballPrizes:    rel.BlitzballPrizes,
	}

	return response, nil
}

func getSphereRelationships(r *http.Request, sphere seedtypes.Sphere) (apitypes.Sphere, error) {
	item, _ := lookup.GetResourceByID(sphere.ItemID, lookup.L.ItemsID)

	availabilityParams, err := apilookups.E.Spheres.QueryAvlRelGetParams(r, item.ID)
	if err != nil {
		return apitypes.Sphere{}, err
	}

	itemRel, err := runItemRelQueries(r, item, availabilityParams)
	if err != nil {
		return apitypes.Sphere{}, err
	}

	rel := apitypes.Sphere{
		Item:            handlerinput.NameToNamedAPIResource(apilookups.E.Items, item.Name, nil),
		Description:     item.Description,
		Effect:          item.Effect,
		Monsters:        itemRel.Monsters,
		Treasures:       itemRel.Treasures,
		Shops:           itemRel.Shops,
		Quests:          itemRel.Quests,
		BlitzballPrizes: itemRel.BlitzballPrizes,
	}

	return rel, nil
}
