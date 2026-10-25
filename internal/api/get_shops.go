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
		Area:        idToAreaAPIResource(cfg, cfg.E.Areas, shop.AreaID),
		Notes:       shop.Notes,
		Category:    shop.Category,
		PreAirship:  convertObjPtr(cfg, shop.PreAirship, convertSubShop),
		PostAirship: convertObjPtr(cfg, shop.PostAirship, convertSubShop),
	}

	return response, nil
}

func (cfg *Config) retrieveShops(r *http.Request, i handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		joinedQuery(cfg, r, i, ids, []QueryParamName{QpnAutoAbility, QpnEmptySlots, QpnCharacter}, filterShopsEquipment),
		idQuery(r, i, ids, QpnLocation, cfg.L.Locations, cfg.db.GetShopIDsByLocation),
		idQuery(r, i, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetShopIDsBySublocation),
		boolQuery2(r, i, ids, QpnItems, cfg.db.GetShopIDsWithItems),
		boolQuery2(r, i, ids, QpnEquipment, cfg.db.GetShopIDsWithEquipment),
		enumListQuery(cfg, r, i, cfg.T.ShopCategory, ids, QpnCategory, cfg.db.GetShopIDsByCategory),
	})
}
