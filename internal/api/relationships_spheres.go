package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getSphereRelationships(cfg *Config, r *http.Request, sphere seeding.Sphere) (Sphere, error) {
	item, _ := seeding.GetResourceByID(sphere.ItemID, cfg.l.ItemsID)

	availabilityParams, err := getAvailabilityParams(cfg, r, cfg.e.spheres, item.ID)
	if err != nil {
		return Sphere{}, err
	}

	itemRel, err := runItemRelQueries(cfg, r, item, availabilityParams)
	if err != nil {
		return Sphere{}, err
	}

	rel := Sphere{
		Item: 				nameToNamedAPIResource(cfg, cfg.e.items, item.Name, nil),
		Description: 		item.Description,
		Effect: 			item.Effect,
		Monsters: 			itemRel.Monsters,
		Treasures: 			itemRel.Treasures,
		Shops: 				itemRel.Shops,
		Quests: 			itemRel.Quests,
		BlitzballPrizes: 	itemRel.BlitzballPrizes,
	}

	return rel, nil
}