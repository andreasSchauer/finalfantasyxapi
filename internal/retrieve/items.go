package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func RetrieveItems(r *http.Request, i handlerinput.Main[seedtypes.Item, apitypes.Item, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.ItemCategory, ids, ts.QpnCategory, config.Cfg.DB.GetItemIDsCategory),
		i.QueryBool2(r, ids, ts.QpnHasAbility, config.Cfg.DB.GetItemIDsWithAbility),
		i.QueryNameID(r, ids, ts.QpnRelatedStat, apilookups.E.Stats.ResTypeSingle, lookup.L.Stats, config.Cfg.DB.GetItemIDsByRelatedStat),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetItemIDsByMethods),
		i.QueryIDWrapper(r, ids, ts.QpnLocation, apilookups.E.Locations.ObjLookup, GetItemsByLocation),
		i.QueryIDWrapper(r, ids, ts.QpnSublocation, apilookups.E.Sublocations.ObjLookup, GetItemsBySublocation),
		i.QueryIDWrapper(r, ids, ts.QpnArea, apilookups.E.Areas.ObjLookup, GetItemsByArea),
	})
}
