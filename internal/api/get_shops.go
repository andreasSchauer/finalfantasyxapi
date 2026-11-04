package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getShop(r *http.Request, i handlerInput[seedtypes.Shop, Shop, UnnamedAPIResource], id int32) (Shop, error) {
	shop, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Shop{}, err
	}

	response := Shop{
		ID:          shop.ID,
		Area:        idToAreaAPIResource(s, E.Areas, shop.AreaID),
		Notes:       shop.Notes,
		Category:    shop.Category,
		PreAirship:  convertObjPtr(s, shop.PreAirship, convertSubShop),
		PostAirship: convertObjPtr(s, shop.PostAirship, convertSubShop),
	}

	return response, nil
}

func (s *Server) retrieveShops(r *http.Request, i handlerInput[seedtypes.Shop, Shop, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryJoined(s, r, ids, []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnEmptySlots, ts.QpnCharacter}, filterShopsEquipment),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, s.Cfg.DB.GetShopIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, s.Cfg.DB.GetShopIDsBySublocation),
		i.queryBool2(r, ids, ts.QpnItems, s.Cfg.DB.GetShopIDsWithItems),
		i.queryBool2(r, ids, ts.QpnEquipment, s.Cfg.DB.GetShopIDsWithEquipment),
		i.queryEnumList(s, r, enums.T.ShopCategory, ids, ts.QpnCategory, s.Cfg.DB.GetShopIDsByCategory),
	})
}
