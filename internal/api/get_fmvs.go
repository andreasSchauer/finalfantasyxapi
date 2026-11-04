package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getFMV(r *http.Request, i handlerInput[seedtypes.FMV, FMV, NamedAPIResource], id int32) (FMV, error) {
	fmv, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return FMV{}, err
	}

	response := FMV{
		ID:                  fmv.ID,
		Name:                fmv.Name,
		Translation:         fmv.Translation,
		CutsceneDescription: fmv.CutsceneDescription,
		Area:                locAreaToAreaAPIResource(s, E.Areas, fmv.LocationArea),
		Song:                namePtrToNamedAPIResPtr(s, E.Songs, fmv.SongName, nil),
	}

	return response, nil
}

func (s *Server) retrieveFMVs(r *http.Request, i handlerInput[seedtypes.FMV, FMV, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, s.Cfg.DB.GetLocationFmvIDs),
	})
}
