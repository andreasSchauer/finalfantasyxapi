package get

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

func GetFMV(r *http.Request, i handlerinput.Main[seedtypes.FMV, apitypes.FMV, apiresource.Named], id int32) (apitypes.FMV, error) {
	fmv, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.FMV{}, err
	}

	response := apitypes.FMV{
		ID:                  fmv.ID,
		Name:                fmv.Name,
		Translation:         fmv.Translation,
		CutsceneDescription: fmv.CutsceneDescription,
		Area:                handlerinput.LocAreaToAreaAPIResource(apilookups.E.Areas, fmv.LocationArea),
		Song:                handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Songs, fmv.SongName, nil),
	}

	return response, nil
}

func RetrieveFMVs(r *http.Request, i handlerinput.Main[seedtypes.FMV, apitypes.FMV, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationFmvIDs),
	})
}
