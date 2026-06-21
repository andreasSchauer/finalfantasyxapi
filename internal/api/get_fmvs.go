package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getFMV(r *http.Request, i handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList], id int32) (FMV, error) {
	fmv, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return FMV{}, err
	}

	response := FMV{
		ID:                  fmv.ID,
		Name:                fmv.Name,
		Translation:         fmv.Translation,
		CutsceneDescription: fmv.CutsceneDescription,
		Area:                locAreaToAreaAPIResource(cfg, cfg.e.areas, fmv.LocationArea),
		Song:                namePtrToNamedAPIResPtr(cfg, cfg.e.songs, fmv.SongName, nil),
	}

	return response, nil
}

func (cfg *Config) retrieveFMVs(r *http.Request, i handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idQuery(r, i, ids, "location", cfg.l.Locations, cfg.db.GetLocationFmvIDs)),
	})
}
