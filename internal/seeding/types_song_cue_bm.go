package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type BackgroundMusic struct {
	ID                     int32
	Condition              *string        `json:"condition"`
	ReplacesEncounterMusic bool           `json:"replaces_encounter_music"`
	LocationAreas          []LocationArea `json:"location_areas"`
}

func (bm BackgroundMusic) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", bm),
		h.DerefOrNil(bm.Condition),
		bm.ReplacesEncounterMusic,
	}
}

func (bm BackgroundMusic) GetID() int32 {
	return bm.ID
}

func (bm *BackgroundMusic) SetID(id int32) {
	bm.ID = id
}

func (bm BackgroundMusic) Error() string {
	return fmt.Sprintf("background music replacing encounter music: %t, condition: %v", bm.ReplacesEncounterMusic, h.PtrToString(bm.Condition))
}

type Cue struct {
	ID                     int32
	SongID                 int32
	TriggerAreaID          *int32
	SceneDescription       string         `json:"scene_description"`
	TriggerLocationArea    *LocationArea  `json:"trigger_location_area"`
	IncludedAreas          []LocationArea `json:"included_areas"`
	ReplacesBGMusic        *string        `json:"replaces_bg_music"`
	EndTrigger             *string        `json:"end_trigger"`
	ReplacesEncounterMusic bool           `json:"replaces_encounter_music"`
}

func (c Cue) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", c),
		c.SceneDescription,
		c.SongID,
		h.DerefOrNil(c.TriggerAreaID),
		h.DerefOrNil(c.ReplacesBGMusic),
		h.DerefOrNil(c.EndTrigger),
		c.ReplacesEncounterMusic,
	}
}

func (c Cue) GetID() int32 {
	return c.ID
}

func (c *Cue) SetID(id int32) {
	c.ID = id
}

func (c Cue) Error() string {
	return fmt.Sprintf("cue for scene: %s at %v, with song id: %d", c.SceneDescription, h.PtrToString(c.TriggerLocationArea), c.SongID)
}
