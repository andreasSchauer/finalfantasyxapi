package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getFMV(r *http.Request, i handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedAPIResourceList], id int32) (FMV, error) {
	fmv, err := verifyParamsAndGet(r, i, id)
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

func (cfg *Config) retrieveFMVs(r *http.Request, i handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		idQuery(r, i, ids, QpnLocation, cfg.L.Locations, cfg.db.GetLocationFmvIDs),
	})
}
