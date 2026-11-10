package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func RetrieveMasterItems(r *http.Request, i handlerinput.Main[seedtypes.MasterItem, apitypes.MasterItem, apiresource.Typed]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.ItemType, ids, ts.QpnType, config.Cfg.DB.GetMasterItemIDsByType),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetMasterItemIDsByMethods),
		i.QueryIDWrapper(r, ids, ts.QpnLocation, apilookups.E.Locations.ObjLookup, GetMasterItemsByLocation),
		i.QueryIDWrapper(r, ids, ts.QpnSublocation, apilookups.E.Sublocations.ObjLookup, GetMasterItemsBySublocation),
		i.QueryIDWrapper(r, ids, ts.QpnArea, apilookups.E.Areas.ObjLookup, GetMasterItemsByArea),
	})
}
