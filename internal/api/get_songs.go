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

	bossFights, err := getResourcesDB(cfg, r, cfg.e.monsterFormations, song, cfg.db.GetSongMonsterFormationIDs)
	if err != nil {
		return Song{}, err
	}

	fmvs, err := getResourcesDB(cfg, r, cfg.e.fmvs, song, cfg.db.GetSongFmvIDs)
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
		BossFights:           bossFights,
		FMVs:                 fmvs,
	}

	if song.Credits != nil {
		response.Composer = song.Credits.Composer
		response.Arranger = song.Credits.Arranger
		response.Performer = song.Credits.Performer
		response.Lyricist = song.Credits.Lyricist
	}

	return response, nil
}

func (cfg *Config) retrieveSongs(r *http.Request, i handlerInput[seeding.Song, Song, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(idQueryWrapper(cfg, r, i, resources, "location", len(cfg.l.Locations), getSongsLocation)),
		frl(idQueryWrapper(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), getSongsSublocation)),
		frl(idQueryWrapper(cfg, r, i, resources, "area", len(cfg.l.Areas), getSongsArea)),
		frl(boolQuery2(cfg, r, i, resources, "special_use", cfg.db.GetSongIDsWithSpecialUseCase)),
		frl(boolQuery2(cfg, r, i, resources, "fmvs", cfg.db.GetSongIDsWithFMVs)),
		frl(nullTypeQuery(cfg, r, i, cfg.t.Composer, resources, "composer", cfg.db.GetSongIDsByComposer)),
		frl(nullTypeQuery(cfg, r, i, cfg.t.Arranger, resources, "arranger", cfg.db.GetSongIDsByArranger)),
	})
}

func getSongsLocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	location, _ := seeding.GetResourceByID(id, cfg.l.LocationsID)

	queries := LocBasedMusicQueries{
		CueSongs:  cfg.db.GetLocationCueSongIDs,
		BmSongs:   cfg.db.GetLocationBackgroundMusicSongIDs,
		FMVSongs:  cfg.db.GetLocationFMVSongIDs,
		BossMusic: cfg.db.GetLocationBossSongIDs,
	}

	return getLocBasedSongs(cfg, r, location, queries)
}

func getSongsSublocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	sublocation, _ := seeding.GetResourceByID(id, cfg.l.SublocationsID)

	queries := LocBasedMusicQueries{
		CueSongs:  cfg.db.GetSublocationCueSongIDs,
		BmSongs:   cfg.db.GetSublocationBackgroundMusicSongIDs,
		FMVSongs:  cfg.db.GetSublocationFMVSongIDs,
		BossMusic: cfg.db.GetSublocationBossSongIDs,
	}

	return getLocBasedSongs(cfg, r, sublocation, queries)
}

func getSongsArea(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	area, _ := seeding.GetResourceByID(id, cfg.l.AreasID)

	queries := LocBasedMusicQueries{
		CueSongs:  cfg.db.GetAreaCueSongIDs,
		BmSongs:   cfg.db.GetAreaBackgroundMusicSongIDs,
		FMVSongs:  cfg.db.GetAreaFMVSongIDs,
		BossMusic: cfg.db.GetAreaBossSongIDs,
	}

	return getLocBasedSongs(cfg, r, area, queries)
}

func getLocBasedSongs(cfg *Config, r *http.Request, item seeding.LookupableID, queries LocBasedMusicQueries) ([]NamedAPIResource, error) {
	i := cfg.e.songs
	resources := []NamedAPIResource{}

	filteredLists := []filteredResList[NamedAPIResource]{
		frl(getResourcesDB(cfg, r, i, item, queries.CueSongs)),
		frl(getResourcesDB(cfg, r, i, item, queries.BmSongs)),
		frl(getResourcesDB(cfg, r, i, item, queries.FMVSongs)),
		frl(getResourcesDB(cfg, r, i, item, queries.BossMusic)),
	}

	for _, filtered := range filteredLists {
		if filtered.err != nil {
			return nil, filtered.err
		}
		resources = combineResources(resources, filtered.resources)
	}

	return resources, nil
}
