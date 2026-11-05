package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getShop(r *http.Request, i handlerInput[seedtypes.Shop, Shop, UnnamedAPIResource], id int32) (Shop, error) {
	shop, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Shop{}, err
	}

	response := Shop{
		ID:          shop.ID,
		Area:        idToAreaAPIResource(srv, E.Areas, shop.AreaID),
		Notes:       shop.Notes,
		Category:    shop.Category,
		PreAirship:  convertObjPtr(srv, shop.PreAirship, convertSubShop),
		PostAirship: convertObjPtr(srv, shop.PostAirship, convertSubShop),
	}

	return response, nil
}

func (srv *Server) retrieveShops(r *http.Request, i handlerInput[seedtypes.Shop, Shop, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryJoined(srv, r, ids, []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnEmptySlots, ts.QpnCharacter}, filterShopsEquipment),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetShopIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetShopIDsBySublocation),
		i.queryBool2(r, ids, ts.QpnItems, config.Cfg.DB.GetShopIDsWithItems),
		i.queryBool2(r, ids, ts.QpnEquipment, config.Cfg.DB.GetShopIDsWithEquipment),
		i.queryEnumList(srv, r, enums.T.ShopCategory, ids, ts.QpnCategory, config.Cfg.DB.GetShopIDsByCategory),
	})
}
