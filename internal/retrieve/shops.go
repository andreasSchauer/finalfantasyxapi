package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func RetrieveShops(r *http.Request, i handlerinput.Main[seedtypes.Shop, apitypes.Shop, apiresource.Unnamed]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryJoined(r, ids, []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnEmptySlots, ts.QpnCharacter}, FilterShopsEquipment),
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetShopIDsByLocation),
		i.QueryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetShopIDsBySublocation),
		i.QueryBool2(r, ids, ts.QpnItems, config.Cfg.DB.GetShopIDsWithItems),
		i.QueryBool2(r, ids, ts.QpnEquipment, config.Cfg.DB.GetShopIDsWithEquipment),
		i.QueryEnumList(r, enums.T.ShopCategory, ids, ts.QpnCategory, config.Cfg.DB.GetShopIDsByCategory),
	})
}
