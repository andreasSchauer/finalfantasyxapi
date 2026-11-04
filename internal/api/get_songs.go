package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getSong(r *http.Request, i handlerInput[seedtypes.Song, Song, NamedAPIResource], id int32) (Song, error) {
	song, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Song{}, err
	}

	rel, err := getSongRelationships(cfg, r, song)
	if err != nil {
		return Song{}, err
	}

	response := Song{
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
		BackgroundMusic:      convertObjSlice(cfg, song.BackgroundMusic, convertBackgroundMusic),
		Cues:                 convertObjSlice(cfg, song.Cues, convertCue),
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

func (cfg *Config) retrieveSongs(r *http.Request, i handlerInput[seedtypes.Song, Song, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, cfg.db.GetLocationSongIDs),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, cfg.db.GetSublocationSongIDs),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, cfg.db.GetAreaSongIDs),
		i.queryBool2(r, ids, ts.QpnSpecialUse, cfg.db.GetSongIDsWithSpecialUseCase),
		i.queryBool2(r, ids, ts.QpnFMVs, cfg.db.GetSongIDsWithFMVs),
		i.queryEnum(r, enums.T.Composer, ids, ts.QpnComposer, ToEnumQuery(enums.T.Composer, cfg.db.GetSongIDsByComposer)),
		i.queryEnum(r, enums.T.Arranger, ids, ts.QpnArranger, ToEnumQuery(enums.T.Arranger, cfg.db.GetSongIDsByArranger)),
	})
}
