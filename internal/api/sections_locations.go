package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type LocationSub struct {
	ID        int32            `json:"id"`
	URL       string           `json:"url"`
	Name      string           `json:"name"`
	Shops     []ShopLocSub     `json:"shops"`
	Treasures *TreasuresLocSub `json:"treasures"`
	Monsters  []string         `json:"monsters"`
}

func (l LocationSub) GetURL() string {
	return l.URL
}


func createLocationSub(cfg *Config, r *http.Request, id int32) (SubResource, error) {
	i := cfg.e.locations
	location, _ := seeding.GetResourceByID(id, i.objLookupID)

	monsters, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.monsters.resourceType, id, cfg.db.GetLocationMonsterIDs, idToMonsterSubString)
	if err != nil {
		return LocationSub{}, err
	}

	shops, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.shops.resourceType, id, cfg.db.GetLocationShopIDs, idToShopLocSub)
	if err != nil {
		return LocationSub{}, err
	}

	treasures, err := getTreasuresLocSub(cfg, r, i.resourceType, id, cfg.db.GetLocationTreasureIDs)
	if err != nil {
		return LocationSub{}, err
	}

	locationSub := LocationSub{
		ID:        location.ID,
		URL:       createResourceURL(cfg, i.endpoint, id),
		Name:      location.Name,
		Shops:     shops,
		Treasures: treasures,
		Monsters:  monsters,
	}

	return locationSub, nil
}