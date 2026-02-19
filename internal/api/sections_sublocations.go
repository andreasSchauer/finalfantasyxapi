package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type SublocationSimple struct {
	ID             int32               `json:"id"`
	URL            string              `json:"url"`
	ParentLocation SimpleRef           `json:"parent_location"`
	Name           string              `json:"name"`
	Shops          []ShopLocSimple     `json:"shops"`
	Treasures      *TreasuresLocSimple `json:"treasures"`
	Monsters       []string            `json:"monsters"`
}

func (s SublocationSimple) GetURL() string {
	return s.URL
}

func createSublocationSimple(cfg *Config, r *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.sublocations
	sublocation, _ := seeding.GetResourceByID(id, i.objLookupID)

	monsters, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.monsters.resourceType, id, cfg.db.GetSublocationMonsterIDs, idToMonsterSimpleString)
	if err != nil {
		return SublocationSimple{}, err
	}

	shops, err := dbQueryToSlice(cfg, r, i.resourceType, cfg.e.shops.resourceType, id, cfg.db.GetSublocationShopIDs, idToShopLocSimple)
	if err != nil {
		return SublocationSimple{}, err
	}

	treasures, err := getTreasuresLocSimple(cfg, r, i.resourceType, id, cfg.db.GetSublocationTreasureIDs)
	if err != nil {
		return SublocationSimple{}, err
	}

	sublocationSimple := SublocationSimple{
		ID:             sublocation.ID,
		URL:            createResourceURL(cfg, i.endpoint, id),
		ParentLocation: createSimpleRef(sublocation.Location.ID, sublocation.Location.Name, nil, nil),
		Name:           sublocation.Name,
		Shops:          shops,
		Treasures:      treasures,
		Monsters:       monsters,
	}

	return sublocationSimple, nil
}
