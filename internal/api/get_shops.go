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

func (cfg *Config) retrieveShops(r *http.Request, i handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(joinedQuery(cfg, r, i, ids, []QueryParamName{qpnAutoAbility, qpnEmptySlots, qpnCharacter}, filterShopsEquipment)),
		fidl(idQuery(r, i, ids, qpnLocation, cfg.l.Locations, cfg.db.GetShopIDsByLocation)),
		fidl(idQuery(r, i, ids, qpnSublocation, cfg.l.Sublocations, cfg.db.GetShopIDsBySublocation)),
		fidl(boolQuery2(r, i, ids, qpnItems, cfg.db.GetShopIDsWithItems)),
		fidl(boolQuery2(r, i, ids, qpnEquipment, cfg.db.GetShopIDsWithEquipment)),
		fidl(enumListQuery(cfg, r, i, cfg.t.ShopCategory, ids, qpnCategory, cfg.db.GetShopIDsByCategory)),
	})
}
