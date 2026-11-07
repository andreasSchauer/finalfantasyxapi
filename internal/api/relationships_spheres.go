package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getSphereRelationships(r *http.Request, sphere seedtypes.Sphere) (apitypes.Sphere, error) {
	item, _ := lookup.GetResourceByID(sphere.ItemID, lookup.L.ItemsID)

	availabilityParams, err := E.Spheres.QueryAvlRelGetParams(r, item.ID)
	if err != nil {
		return apitypes.Sphere{}, err
	}

	itemRel, err := runItemRelQueries(r, item, availabilityParams)
	if err != nil {
		return apitypes.Sphere{}, err
	}

	rel := apitypes.Sphere{
		Item:            nameToNamedAPIResource(E.Items, item.Name, nil),
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
