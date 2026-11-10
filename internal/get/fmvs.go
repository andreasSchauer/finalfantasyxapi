package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
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
