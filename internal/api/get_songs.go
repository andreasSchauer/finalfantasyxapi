package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getSong(r *http.Request, i handlerInput[seedtypes.Song, apitypes.Song, apiresource.NamedAPIResource], id int32) (apitypes.Song, error) {
	song, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.Song{}, err
	}

	rel, err := getSongRelationships(r, song)
	if err != nil {
		return apitypes.Song{}, err
	}

	response := apitypes.Song{
		ID:                   song.ID,
		Name:                 song.Name,
		StreamingName:        song.StreamingName,
		InGameName:           song.InGameName,
		OstName:              song.OSTName,
		Translation:          song.Translation,
		DurationInSeconds:    song.DurationInSeconds,
		CanLoop:              song.CanLoop,
		SpecialUseCase:       song.SpecialUseCase,
		StreamingTrackNumber: song.StreamingTrackNumber,
		MusicSphereID:        song.MusicSphereID,
		OstDisc:              song.OSTDisc,
		OstTrackNumber:       song.OSTTrackNumber,
		BackgroundMusic:      ah.ConvertObjSlice(song.BackgroundMusic, convertBackgroundMusic),
		Cues:                 ah.ConvertObjSlice(song.Cues, convertCue),
		BossFights:           rel.BossFights,
		FMVs:                 rel.FMVs,
	}

	if song.Credits != nil {
		response.Composer = song.Credits.Composer
		response.Arranger = song.Credits.Arranger
		response.Performer = song.Credits.Performer
		response.Lyricist = song.Credits.Lyricist
	}

	return response, nil
}

func retrieveSongs(r *http.Request, i handlerInput[seedtypes.Song, apitypes.Song, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationSongIDs),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetSublocationSongIDs),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, config.Cfg.DB.GetAreaSongIDs),
		i.queryBool2(r, ids, ts.QpnSpecialUse, config.Cfg.DB.GetSongIDsWithSpecialUseCase),
		i.queryBool2(r, ids, ts.QpnFMVs, config.Cfg.DB.GetSongIDsWithFMVs),
		i.queryEnum(r, enums.T.Composer, ids, ts.QpnComposer, ToEnumQuery(enums.T.Composer, config.Cfg.DB.GetSongIDsByComposer)),
		i.queryEnum(r, enums.T.Arranger, ids, ts.QpnArranger, ToEnumQuery(enums.T.Arranger, config.Cfg.DB.GetSongIDsByArranger)),
	})
}
