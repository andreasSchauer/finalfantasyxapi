package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSong(r *http.Request, i handlerInput[seeding.Song, Song, NamedAPIResource, NamedApiResourceList], id int32) (Song, error) {
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

func (cfg *Config) retrieveSongs(r *http.Request, i handlerInput[seeding.Song, Song, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idQuery(r, i, ids, qpnLocation, cfg.l.Locations, cfg.db.GetLocationSongIDs)),
		fidl(idQuery(r, i, ids, qpnSublocation, cfg.l.Sublocations, cfg.db.GetSublocationSongIDs)),
		fidl(idQuery(r, i, ids, qpnArea, cfg.l.Areas, cfg.db.GetAreaSongIDs)),
		fidl(boolQuery2(r, i, ids, qpnSpecialUse, cfg.db.GetSongIDsWithSpecialUseCase)),
		fidl(boolQuery2(r, i, ids, qpnFMVs, cfg.db.GetSongIDsWithFMVs)),
		fidl(enumQuery(r, i, cfg.t.Composer, ids, qpnComposer, ToEnumQuery(cfg.t.Composer, cfg.db.GetSongIDsByComposer))),
		fidl(enumQuery(r, i, cfg.t.Arranger, ids, qpnArranger, ToEnumQuery(cfg.t.Arranger, cfg.db.GetSongIDsByArranger))),
	})
}
