package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getShop(r *http.Request, i handlerInput[seeding.Shop, Shop, UnnamedAPIResource], id int32) (Shop, error) {
	shop, err := i.queryParamsVerifyGet(r, id)
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

func (cfg *Config) retrieveShops(r *http.Request, i handlerInput[seeding.Shop, Shop, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryJoined(cfg, r, ids, []QueryParamName{QpnAutoAbility, QpnEmptySlots, QpnCharacter}, filterShopsEquipment),
		i.queryID(r, ids, QpnLocation, cfg.L.Locations, cfg.db.GetShopIDsByLocation),
		i.queryID(r, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetShopIDsBySublocation),
		i.queryBool2(r, ids, QpnItems, cfg.db.GetShopIDsWithItems),
		i.queryBool2(r, ids, QpnEquipment, cfg.db.GetShopIDsWithEquipment),
		i.queryEnumList(cfg, r, cfg.T.ShopCategory, ids, QpnCategory, cfg.db.GetShopIDsByCategory),
	})
}
