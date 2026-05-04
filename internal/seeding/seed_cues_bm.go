package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop1SeedBackgroundMusic(qtx *database.Queries, ctx context.Context) error {
	bm := l.extractBackgroundMusic()

	params := database.CreateBackgroundMusicBulkParams{
		DataHash:               make([]string, len(bm)),
		Condition:              make([]sql.NullString, len(bm)),
		ReplacesEncounterMusic: make([]bool, len(bm)),
	}

	for i, music := range bm {
		params.DataHash[i] = generateDataHash(music)
		params.Condition[i] = h.GetNullString(music.Condition)
		params.ReplacesEncounterMusic[i] = music.ReplacesEncounterMusic
	}

	dbRows, err := qtx.CreateBackgroundMusicBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create background music: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractBackgroundMusic() []BackgroundMusic {
	bgMusic := []BackgroundMusic{}

	for _, song := range l.json.songs {
		bgMusic = append(bgMusic, song.BackgroundMusic...)
	}

	return dedupeRows(bgMusic, l.Hashes)
}

func (l *Lookup) loop4SeedCues(qtx *database.Queries, ctx context.Context) error {
	cues, err := l.extractCues()
	if err != nil {
		return err
	}

	params := database.CreateCueBulkParams{
		DataHash:               make([]string, len(cues)),
		SongID:                 make([]int32, len(cues)),
		SceneDescription:       make([]string, len(cues)),
		TriggerAreaID:          make([]sql.NullInt32, len(cues)),
		ReplacesBgMusic:        make([]database.NullBgReplacementType, len(cues)),
		EndTrigger:             make([]sql.NullString, len(cues)),
		ReplacesEncounterMusic: make([]bool, len(cues)),
	}

	for i, c := range cues {
		params.DataHash[i] = generateDataHash(c)
		params.SongID[i] = c.SongID
		params.SceneDescription[i] = c.SceneDescription
		params.TriggerAreaID[i] = h.ObjPtrToNullInt32ID(c.TriggerLocationArea)
		params.ReplacesBgMusic[i] = database.ToNullBgReplacementType(c.ReplacesBGMusic)
		params.EndTrigger[i] = h.GetNullString(c.EndTrigger)
		params.ReplacesEncounterMusic[i] = c.ReplacesEncounterMusic
	}

	dbRows, err := qtx.CreateCueBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create cues: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractCues() ([]Cue, error) {
	cues := []Cue{}
	var err error

	for i := range l.json.songs {
		song := &l.json.songs[i]

		for j := range song.Cues {
			cue := &song.Cues[j]
			cue.SongID = song.ID

			if cue.TriggerLocationArea != nil {
				cue.TriggerLocationArea.ID, err = assignFK(*cue.TriggerLocationArea, l.Areas)
				if err != nil {
					return nil, err
				}
			}

			cues = append(cues, *cue)
		}
	}

	return dedupeRows(cues, l.Hashes), nil
}

func (l *Lookup) getCues() []Cue {
	cues := []Cue{}

	for _, song := range l.json.songs {
		cues = append(cues, song.Cues...)
	}

	return cues
}

func (l *Lookup) getCueIncludedAreas(c Cue) ([]Area, error) {
	return getResources(c.IncludedAreas, l.Areas)
}

func (l *Lookup) seedJuncCuesIncludedAreas(qtx *database.Queries, ctx context.Context) error {
	const desc string = "cues + included areas"
	jParams, err := processJunctions(l, desc, l.getCues(), l.getCueIncludedAreas)
	if err != nil {
		return err
	}

	return qtx.CreateCuesIncludedAreasJunctionBulk(ctx, database.CreateCuesIncludedAreasJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		CueID:          jParams.ParentIDs,
		IncludedAreaID: jParams.ChildIDs,
	})
}
