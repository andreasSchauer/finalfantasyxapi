package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func KeyItems(r *http.Request, i handlerinput.Main[seedtypes.KeyItem, apitypes.KeyItem, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.KeyItemCategory, ids, ts.QpnCategory, config.Cfg.DB.GetKeyItemIDsCategory),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetKeyItemIDsByMethods),
		i.QueryID(r, ids, ts.QpnLocation, apilookups.E.Locations.ObjLookup, config.Cfg.DB.GetKeyItemIDsByLocation),
		i.QueryID(r, ids, ts.QpnSublocation, apilookups.E.Sublocations.ObjLookup, config.Cfg.DB.GetKeyItemIDsBySublocation),
		i.QueryID(r, ids, ts.QpnArea, apilookups.E.Areas.ObjLookup, config.Cfg.DB.GetKeyItemIDsByArea),
	})
}
