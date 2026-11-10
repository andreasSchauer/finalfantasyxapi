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

func RetrieveSpheres(r *http.Request, i handlerinput.Main[seedtypes.Sphere, apitypes.Sphere, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.SphereColor, ids, ts.QpnColor, config.Cfg.DB.GetSphereIDsByColor),
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetSphereIDsByLocation),
		i.QueryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetSphereIDsBySublocation),
		i.QueryID(r, ids, ts.QpnArea, lookup.L.Areas, config.Cfg.DB.GetSphereIDsByArea),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetSphereIDsByMethods),
	})
}
