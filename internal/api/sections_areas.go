package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AreaSimple struct {
	ID                int32               `json:"id"`
	URL               string              `json:"url"`
	ParentLocation    SimpleRef           `json:"parent_location"`
	ParentSublocation SimpleRef           `json:"parent_sublocation"`
	Name              string              `json:"name"`
	Version           *int32              `json:"version,omitempty"`
	Specification     *string             `json:"specification,omitempty"`
	HasSaveSphere     bool                `json:"has_save_sphere"`
	StoryOnly         bool                `json:"story_only"`
	Shops             []ShopLocSimple     `json:"shops"`
	Treasures         *TreasuresLocSimple `json:"treasures"`
	Monsters          []string            `json:"monsters"`
}

func (a AreaSimple) GetURL() string {
	return a.URL
}

func createAreaSimple(cfg *Config, r *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.areas
	area, _ := seeding.GetResourceByID(id, i.objLookupID)

	monsters, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.monsters.resourceType, id, cfg.db.GetAreaMonsterIDs, idToMonsterSimpleString)
	if err != nil {
		return AreaSimple{}, err
	}

	shops, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.shops.resourceType, id, cfg.db.GetAreaShopIDs, idToShopLocSimple)
	if err != nil {
		return AreaSimple{}, err
	}

	treasures, err := getTreasuresLocSimple(cfg, r, i.resourceType, id, cfg.db.GetAreaTreasureIDs)
	if err != nil {
		return AreaSimple{}, err
	}

	areaSimple := AreaSimple{
		ID:                area.ID,
		URL:               createResourceURL(cfg, i.endpoint, id),
		ParentLocation:    createSimpleRef(area.Sublocation.Location.ID, area.Sublocation.Location.Name, nil, nil),
		ParentSublocation: createSimpleRef(area.Sublocation.ID, area.Sublocation.Name, nil, nil),
		Name:              area.Name,
		Version:           area.Version,
		Specification:     area.Specification,
		HasSaveSphere:     area.HasSaveSphere,
		StoryOnly:         area.StoryOnly,
		Shops:             shops,
		Treasures:         treasures,
		Monsters:          monsters,
	}

	return areaSimple, nil
}
