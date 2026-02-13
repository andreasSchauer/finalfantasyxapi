package api

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


func createAreaSub(cfg *Config, r *http.Request, id int32) (SubResource, error) {
	i := cfg.e.areas
	area, _ := seeding.GetResourceByID(id, i.objLookupID)
	
	monsters, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.monsters.resourceType, id, cfg.db.GetAreaMonsterIDs, idToMonsterSubString)
	if err != nil {
		return AreaSub{}, err
	}

	shops, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.shops.resourceType, id, cfg.db.GetAreaShopIDs, idToShopLocSub)
	if err != nil {
		return AreaSub{}, err
	}

	treasures, err := getTreasuresLocSub(cfg, r, i.resourceType, id, cfg.db.GetAreaTreasureIDs)
	if err != nil {
		return AreaSub{}, err
	}

	areaSub := AreaSub{
		ID:                area.ID,
		URL:               createResourceURL(cfg, i.endpoint, id),
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

	return areaSub, nil
}