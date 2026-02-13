package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Song struct {
	ID                   int32                `json:"id"`
	Name                 string               `json:"name"`
	StreamingName        *string              `json:"streaming_name"`
	InGameName           *string              `json:"in_game_name"`
	OstName              *string              `json:"ost_name"`
	Translation          *string              `json:"translation,omitempty"`
	Composer             *string              `json:"composer,omitempty"`
	Arranger             *string              `json:"arranger,omitempty"`
	Performer            *string              `json:"performer,omitempty"`
	Lyricist             *string              `json:"lyricist,omitempty"`
	DurationInSeconds    int32                `json:"duration_in_seconds"`
	CanLoop              bool                 `json:"can_loop"`
	SpecialUseCase       *string              `json:"special_use_case"`
	StreamingTrackNumber *int32               `json:"streaming_track_number,omitempty"`
	MusicSphereID        *int32               `json:"music_sphere_id,omitempty"`
	OstDisc              *int32               `json:"ost_disc,omitempty"`
	OstTrackNumber       *int32               `json:"ost_track_number,omitempty"`
	BackgroundMusic      []BackgroundMusic    `json:"background_music"`
	Cues                 []Cue                `json:"cues"`
	BossFights           []UnnamedAPIResource `json:"boss_fights"`
	FMVs                 []NamedAPIResource   `json:"fmvs"`
}

type BackgroundMusic struct {
	Condition              *string           `json:"condition,omitempty"`
	ReplacesEncounterMusic bool              `json:"replaces_encounter_music"`
	Areas                  []AreaAPIResource `json:"areas"`
}

func convertBackgroundMusic(cfg *Config, bm seeding.BackgroundMusic) BackgroundMusic {
	return BackgroundMusic{
		Condition:              bm.Condition,
		ReplacesEncounterMusic: bm.ReplacesEncounterMusic,
		Areas:                  locAreasToAreaAPIResources(cfg, cfg.e.areas, bm.LocationAreas),
	}
}

type Cue struct {
	SceneDescription       string            `json:"scene_description"`
	TriggerArea            *AreaAPIResource  `json:"trigger_area"`
	IncludedAreas          []AreaAPIResource `json:"included_areas"`
	ReplacesEncounterMusic bool              `json:"replaces_encounter_music"`
	ReplacesBGMusic        *string           `json:"replaces_bg_music"`
	EndTrigger             *string           `json:"end_trigger,omitempty"`
}

func convertCue(cfg *Config, cue seeding.Cue) Cue {
	var triggerAreaPtr *AreaAPIResource

	if cue.TriggerLocationArea != nil {
		triggerArea := locAreaToAreaAPIResource(cfg, cfg.e.areas, *cue.TriggerLocationArea)
		triggerAreaPtr = &triggerArea
	}

	return Cue{
		SceneDescription:       cue.SceneDescription,
		TriggerArea:            triggerAreaPtr,
		IncludedAreas:          locAreasToAreaAPIResources(cfg, cfg.e.areas, cue.IncludedAreas),
		ReplacesEncounterMusic: cue.ReplacesEncounterMusic,
		ReplacesBGMusic:        cue.ReplacesBGMusic,
		EndTrigger:             cue.EndTrigger,
	}
}

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
