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
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(joinedQuery(cfg, r, i, ids, []string{"auto_ability", "empty_slots", "character"}, filterShopsEquipment)),
		fidl(idQuery(r, i, ids, "location", cfg.l.Locations, cfg.db.GetShopIDsByLocation)),
		fidl(idQuery(r, i, ids, "sublocation", cfg.l.Sublocations, cfg.db.GetShopIDsBySublocation)),
		fidl(boolQuery2(r, i, ids, "items", cfg.db.GetShopIDsWithItems)),
		fidl(boolQuery2(r, i, ids, "equipment", cfg.db.GetShopIDsWithEquipment)),
		fidl(enumListQuery(cfg, r, i, cfg.t.ShopCategory, ids, "category", cfg.db.GetShopIDsByCategory)),
	})
}
