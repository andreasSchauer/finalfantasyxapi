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

func handleLocationsSection(cfg *Config, r *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.locations
	locations := []LocationSub{}

	for _, locID := range dbIDs {
		location, _ := seeding.GetResourceByID(locID, i.objLookupID)

		monsters, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.monsters.resourceType, locID, cfg.db.GetLocationMonsterIDs, idToMonsterSubString)
		if err != nil {
			return nil, err
		}

		shops, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.shops.resourceType, locID, cfg.db.GetLocationShopIDs, idToShopLocSub)
		if err != nil {
			return nil, err
		}

		treasures, err := getTreasuresLocSub(cfg, r, i.resourceType, locID, cfg.db.GetLocationTreasureIDs)
		if err != nil {
			return nil, err
		}

		locationSub := LocationSub{
			ID:        location.ID,
			URL:       createResourceURL(cfg, i.endpoint, locID),
			Name:      location.Name,
			Shops:     shops,
			Treasures: treasures,
			Monsters:  monsters,
		}

		locations = append(locations, locationSub)
	}

	return toSubResourceSlice(locations), nil
}
