package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getSphereRelationships(srv *Server, r *http.Request, sphere seedtypes.Sphere) (Sphere, error) {
	item, _ := lookup.GetResourceByID(sphere.ItemID, lookup.L.ItemsID)

	availabilityParams, err := E.Spheres.queryAvlRelGetParams(srv, r, item.ID)
	if err != nil {
		return Sphere{}, err
	}

	itemRel, err := runItemRelQueries(srv, r, item, availabilityParams)
	if err != nil {
		return Sphere{}, err
	}

	rel := Sphere{
		Item:            nameToNamedAPIResource(srv, E.Items, item.Name, nil),
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
