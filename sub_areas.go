package main

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AreaSub struct {
	ID                int32            `json:"id"`
	URL               string           `json:"url"`
	ParentLocation    SubName          `json:"parent_location"`
	ParentSublocation SubName          `json:"parent_sublocation"`
	Name              string           `json:"name"`
	Version           *int32           `json:"version,omitempty"`
	Specification     *string          `json:"specification,omitempty"`
	HasSaveSphere     bool             `json:"has_save_sphere"`
	Shops             []ShopLocSub     `json:"shops"`
	Treasures         *TreasuresLocSub `json:"treasures"`
	Monsters          []SubName        `json:"monsters"`
}

func (a AreaSub) GetSectionName() string {
	return "areas"
}

func (a AreaSub) GetURL() string {
	return a.URL
}

type SubName struct {
	ID            int32   `json:"id,omitempty"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
}

func createSubName(id int32, name string, version *int32, spec *string) SubName {
	return SubName{
		ID:            id,
		Name:          name,
		Version:       version,
		Specification: spec,
	}
}

func handleAreasSection(cfg *Config, r *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.areas
	areas := []AreaSub{}

	for _, areaID := range dbIDs {
		area, _ := seeding.GetResourceByID(areaID, i.objLookupID)
		monsters, err := getMonstersLocSub(cfg, r, i.resourceType, areaID, cfg.db.GetAreaMonsterIDs)
		if err != nil {
			return nil, err
		}

		shops, err := getShopsLocSub(cfg, r, i.resourceType, areaID, cfg.db.GetAreaShopIDs)
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
			ParentLocation:    createSubName(area.SubLocation.Location.ID, area.SubLocation.Location.Name, nil, nil),
			ParentSublocation: createSubName(area.SubLocation.ID, area.SubLocation.Name, nil, nil),
			Name:              area.Name,
			Version:           area.Version,
			Specification:     area.Specification,
			HasSaveSphere:     area.HasSaveSphere,
			Shops:             shops,
			Treasures:         treasures,
			Monsters:          monsters,
		}

		areas = append(areas, areaSub)
	}

	return toSubResourceSlice(areas), nil
}
