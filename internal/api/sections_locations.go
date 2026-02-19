package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type LocationSimple struct {
	ID        int32               `json:"id"`
	URL       string              `json:"url"`
	Name      string              `json:"name"`
	Shops     []ShopLocSimple     `json:"shops"`
	Treasures *TreasuresLocSimple `json:"treasures"`
	Monsters  []string            `json:"monsters"`
}

func (l LocationSimple) GetURL() string {
	return l.URL
}

func createLocationSimple(cfg *Config, r *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.locations
	location, _ := seeding.GetResourceByID(id, i.objLookupID)

	monsters, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.monsters.resourceType, id, cfg.db.GetLocationMonsterIDs, idToMonsterSimpleString)
	if err != nil {
		return LocationSimple{}, err
	}

	shops, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.shops.resourceType, id, cfg.db.GetLocationShopIDs, idToShopLocSimple)
	if err != nil {
		return LocationSimple{}, err
	}

	treasures, err := getTreasuresLocSimple(cfg, r, i.resourceType, id, cfg.db.GetLocationTreasureIDs)
	if err != nil {
		return LocationSimple{}, err
	}

	locationSimple := LocationSimple{
		ID:        location.ID,
		URL:       createResourceURL(cfg, i.endpoint, id),
		Name:      location.Name,
		Shops:     shops,
		Treasures: treasures,
		Monsters:  monsters,
	}

	return locationSimple, nil
}
