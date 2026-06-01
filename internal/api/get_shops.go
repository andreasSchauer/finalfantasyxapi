package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getShop(r *http.Request, i handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList], id int32) (Shop, error) {
	shop, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Shop{}, err
	}

	response := Shop{
		ID:          shop.ID,
		Area:        idToAreaAPIResource(cfg, cfg.e.areas, shop.AreaID),
		Notes:       shop.Notes,
		Category:    enumToNamedAPIResource(cfg, cfg.e.shopCategory.endpoint, shop.Category, cfg.t.ShopCategory),
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
		frl(joinedQuery(cfg, r, i, resources, []string{"auto_ability", "empty_slots", "character"}, filterShopsEquipment)),
		frl(idQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetShopIDsByLocation)),
		frl(idQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetShopIDsBySublocation)),
		frl(boolQuery2(cfg, r, i, resources, "items", cfg.db.GetShopIDsWithItems)),
		frl(boolQuery2(cfg, r, i, resources, "equipment", cfg.db.GetShopIDsWithEquipment)),
		frl(enumListQuery(cfg, r, i, cfg.t.ShopCategory, resources, "category", cfg.db.GetShopIDsByCategory)),
	})
}
