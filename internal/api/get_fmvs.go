package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getFMV(r *http.Request, i handlerInput[seedtypes.FMV, FMV, NamedAPIResource], id int32) (FMV, error) {
	fmv, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return FMV{}, err
	}

	response := FMV{
		ID:                  fmv.ID,
		Name:                fmv.Name,
		Translation:         fmv.Translation,
		CutsceneDescription: fmv.CutsceneDescription,
		Area:                locAreaToAreaAPIResource(srv, E.Areas, fmv.LocationArea),
		Song:                namePtrToNamedAPIResPtr(srv, E.Songs, fmv.SongName, nil),
	}

	return response, nil
}

func (srv *Server) retrieveFMVs(r *http.Request, i handlerInput[seedtypes.FMV, FMV, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationFmvIDs),
	})
}
