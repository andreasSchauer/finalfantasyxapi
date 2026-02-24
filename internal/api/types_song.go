package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"


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