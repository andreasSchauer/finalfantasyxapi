package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSong(r *http.Request, i handlerInput[seeding.Song, Song, NamedAPIResource, NamedAPIResourceList], id int32) (Song, error) {
	song, err := verifyParamsAndGet(r, i, id)
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

func (cfg *Config) retrieveSongs(r *http.Request, i handlerInput[seeding.Song, Song, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		idQuery(r, i, ids, QpnLocation, cfg.L.Locations, cfg.db.GetLocationSongIDs),
		idQuery(r, i, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetSublocationSongIDs),
		idQuery(r, i, ids, QpnArea, cfg.L.Areas, cfg.db.GetAreaSongIDs),
		boolQuery2(r, i, ids, QpnSpecialUse, cfg.db.GetSongIDsWithSpecialUseCase),
		boolQuery2(r, i, ids, QpnFMVs, cfg.db.GetSongIDsWithFMVs),
		enumQuery(r, i, cfg.T.Composer, ids, QpnComposer, ToEnumQuery(cfg.T.Composer, cfg.db.GetSongIDsByComposer)),
		enumQuery(r, i, cfg.T.Arranger, ids, QpnArranger, ToEnumQuery(cfg.T.Arranger, cfg.db.GetSongIDsByArranger)),
	})
}
