package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getShop(r *http.Request, i handlerInput[seedtypes.Shop, Shop, UnnamedAPIResource], id int32) (Shop, error) {
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

func (cfg *Config) retrieveShops(r *http.Request, i handlerInput[seedtypes.Shop, Shop, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryJoined(cfg, r, ids, []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnEmptySlots, ts.QpnCharacter}, filterShopsEquipment),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, cfg.db.GetShopIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, cfg.db.GetShopIDsBySublocation),
		i.queryBool2(r, ids, ts.QpnItems, cfg.db.GetShopIDsWithItems),
		i.queryBool2(r, ids, ts.QpnEquipment, cfg.db.GetShopIDsWithEquipment),
		i.queryEnumList(cfg, r, enums.T.ShopCategory, ids, ts.QpnCategory, cfg.db.GetShopIDsByCategory),
	})
}
