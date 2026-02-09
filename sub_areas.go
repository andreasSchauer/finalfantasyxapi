package main

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AreaSub struct {
	ID                int32            `json:"id"`
	URL               string           `json:"url"`
	ParentLocation    SubRef           `json:"parent_location"`
	ParentSublocation SubRef           `json:"parent_sublocation"`
	Name              string           `json:"name"`
	Version           *int32           `json:"version,omitempty"`
	Specification     *string          `json:"specification,omitempty"`
	HasSaveSphere     bool             `json:"has_save_sphere"`
	StoryOnly         bool             `json:"story_only"`
	Shops             []ShopLocSub     `json:"shops"`
	Treasures         *TreasuresLocSub `json:"treasures"`
	Monsters          []string         `json:"monsters"`
}

func (a AreaSub) GetURL() string {
	return a.URL
}

func handleAreasSection(cfg *Config, r *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.areas
	areas := []AreaSub{}

	for _, areaID := range dbIDs {
		area, _ := seeding.GetResourceByID(areaID, i.objLookupID)
		monsters, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.monsters.resourceType, areaID, cfg.db.GetAreaMonsterIDs, idToMonsterSubString)
		if err != nil {
			return nil, err
		}

		shops, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.shops.resourceType, areaID, cfg.db.GetAreaShopIDs, idToShopLocSub)
		if err != nil {
			return nil, err
		}

		treasures, err := getTreasuresLocSub(cfg, r, i.resourceType, areaID, cfg.db.GetAreaTreasureIDs)
		if err != nil {
			return nil, err
		}

		areaSub := AreaSub{
			ID:                area.ID,
			URL:               createResourceURL(cfg, i.endpoint, areaID),
			ParentLocation:    createSubReference(area.Sublocation.Location.ID, area.Sublocation.Location.Name, nil, nil),
			ParentSublocation: createSubReference(area.Sublocation.ID, area.Sublocation.Name, nil, nil),
			Name:              area.Name,
			Version:           area.Version,
			Specification:     area.Specification,
			HasSaveSphere:     area.HasSaveSphere,
			StoryOnly:         area.StoryOnly,
			Shops:             shops,
			Treasures:         treasures,
			Monsters:          monsters,
		}

		areas = append(areas, areaSub)
	}

	return toSubResourceSlice(areas), nil
}
