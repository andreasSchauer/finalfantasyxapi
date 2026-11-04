package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (cfg *Config) getFMV(r *http.Request, i handlerInput[seedtypes.FMV, FMV, NamedAPIResource], id int32) (FMV, error) {
	fmv, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return FMV{}, err
	}

	response := FMV{
		ID:                  fmv.ID,
		Name:                fmv.Name,
		Translation:         fmv.Translation,
		CutsceneDescription: fmv.CutsceneDescription,
		Area:                locAreaToAreaAPIResource(cfg, cfg.E.Areas, fmv.LocationArea),
		Song:                namePtrToNamedAPIResPtr(cfg, cfg.E.Songs, fmv.SongName, nil),
	}

	return response, nil
}

func (cfg *Config) retrieveFMVs(r *http.Request, i handlerInput[seedtypes.FMV, FMV, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryID(r, ids, QpnLocation, lookup.L.Locations, cfg.db.GetLocationFmvIDs),
	})
}
