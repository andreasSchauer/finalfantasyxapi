package api

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

func (s SublocationSub) GetURL() string {
	return s.URL
}


func createSublocationSub(cfg *Config, r *http.Request, id int32) (SubResource, error) {
	i := cfg.e.sublocations
	sublocation, _ := seeding.GetResourceByID(id, i.objLookupID)

	monsters, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.monsters.resourceType, id, cfg.db.GetSublocationMonsterIDs, idToMonsterSubString)
	if err != nil {
		return SublocationSub{}, err
	}

	shops, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.shops.resourceType, id, cfg.db.GetSublocationShopIDs, idToShopLocSub)
	if err != nil {
		return SublocationSub{}, err
	}

	treasures, err := getTreasuresLocSub(cfg, r, i.resourceType, id, cfg.db.GetSublocationTreasureIDs)
	if err != nil {
		return SublocationSub{}, err
	}

	sublocationSub := SublocationSub{
		ID:             sublocation.ID,
		URL:            createResourceURL(cfg, i.endpoint, id),
		ParentLocation: createSubReference(sublocation.Location.ID, sublocation.Location.Name, nil, nil),
		Name:           sublocation.Name,
		Shops:          shops,
		Treasures:      treasures,
		Monsters:       monsters,
	}

	return sublocationSub, nil
}