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

func RetrieveKeyItems(r *http.Request, i handlerinput.Main[seedtypes.KeyItem, apitypes.KeyItem, apiresource.Named]) ([]int32, error) {
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
