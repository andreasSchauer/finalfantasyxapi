package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getShop(r *http.Request, i handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList], id int32) (Shop, error) {
	shop, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Shop{}, err
	}

	response := Shop{
		ID:          shop.ID,
		Area:        idToAreaAPIResource(cfg, cfg.e.areas, shop.AreaID),
		Notes:       shop.Notes,
		Category:    shop.Category,
		PreAirship:  convertObjPtr(cfg, shop.PreAirship, convertSubShop),
		PostAirship: convertObjPtr(cfg, shop.PostAirship, convertSubShop),
	}

	return response, nil
}

func (cfg *Config) retrieveShops(r *http.Request, i handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]) (UnnamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[UnnamedAPIResource]{
		frl(idQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationShopIDs)),
		frl(idQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationShopIDs)),
		frl(idQuery(cfg, r, i, resources, "auto_ability", len(cfg.l.AutoAbilities), cfg.db.GetShopIDsByAutoAbility)),
		frl(boolQuery2(cfg, r, i, resources, "items", cfg.db.GetShopIDsWithItems)),
		frl(boolQuery2(cfg, r, i, resources, "equipment", cfg.db.GetShopIDsWithEquipment)),
		frl(boolQuery2(cfg, r, i, resources, "pre_airship", cfg.db.GetShopIDsPreAirship)),
		frl(boolQuery2(cfg, r, i, resources, "post_airship", cfg.db.GetShopIDsPostAirship)),
		frl(typeQuery(cfg, r, i, cfg.t.ShopCategory, resources, "category", cfg.db.GetShopIDsByCategory)),
	})
}
