package main

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type SublocationSub struct {
	ID             int32            `json:"id"`
	URL            string           `json:"url"`
	ParentLocation SubRef           `json:"parent_location"`
	Name           string           `json:"name"`
	Shops          []ShopLocSub     `json:"shops"`
	Treasures      *TreasuresLocSub `json:"treasures"`
	Monsters       []string         `json:"monsters"`
}

func (s SublocationSub) GetSectionName() string {
	return "sublocations"
}

func (s SublocationSub) GetURL() string {
	return s.URL
}

func handleSublocationsSection(cfg *Config, r *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.sublocations
	sublocations := []SublocationSub{}

	for _, subLocID := range dbIDs {
		sublocation, _ := seeding.GetResourceByID(subLocID, i.objLookupID)
		monsters, err := getMonstersLocSub(cfg, r, i.resourceType, subLocID, cfg.db.GetSublocationMonsterIDs)
		if err != nil {
			return nil, err
		}

		shops, err := getShopsLocSub(cfg, r, i.resourceType, subLocID, cfg.db.GetSublocationShopIDs)
		if err != nil {
			return nil, err
		}

		treasures, err := getTreasuresLocSub(cfg, r, i.resourceType, subLocID, cfg.db.GetSublocationTreasureIDs)
		if err != nil {
			return nil, err
		}

		sublocationSub := SublocationSub{
			ID:             sublocation.ID,
			URL:            createResourceURL(cfg, i.endpoint, subLocID),
			ParentLocation: createSubReference(sublocation.Location.ID, sublocation.Location.Name),
			Name:           sublocation.Name,
			Shops:          shops,
			Treasures:      treasures,
			Monsters:       monsters,
		}

		sublocations = append(sublocations, sublocationSub)
	}

	return toSubResourceSlice(sublocations), nil
}
