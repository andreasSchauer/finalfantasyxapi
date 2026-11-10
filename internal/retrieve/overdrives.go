package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func RetrieveOverdrives(r *http.Request, i handlerinput.Main[seedtypes.Overdrive, apitypes.Overdrive, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryIntList(r, ids, ts.QpnRank, config.Cfg.DB.GetOverdriveIDsByRank),
		i.QueryNameID(r, ids, ts.QpnUser, apilookups.E.CharacterClasses.ResTypeSingle, lookup.L.CharClasses, handlerinput.ToIntManyNull(config.Cfg.DB.GetOverdriveIDsByUser)),
	})
}
