package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop1SeedBackgroundMusic(qtx *database.Queries, ctx context.Context) error {
	bm := extractBackgroundMusic()

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
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractBackgroundMusic() []seedtypes.BackgroundMusic {
	bgMusic := []seedtypes.BackgroundMusic{}

	for _, song := range rawJSON.songs {
		bgMusic = append(bgMusic, song.BackgroundMusic...)
	}

	return dedupeRows(bgMusic, hashes)
}

func loop4SeedCues(qtx *database.Queries, ctx context.Context) error {
	cues, err := extractCues()
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
		params.TriggerAreaID[i] = h.GetNullInt32(c.TriggerAreaID)
		params.ReplacesBgMusic[i] = database.ToNullBgReplacementType(c.ReplacesBGMusic)
		params.EndTrigger[i] = h.GetNullString(c.EndTrigger)
		params.ReplacesEncounterMusic[i] = c.ReplacesEncounterMusic
	}

	dbRows, err := qtx.CreateCueBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create cues: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractCues() ([]seedtypes.Cue, error) {
	cues := []seedtypes.Cue{}
	var err error

	for i := range rawJSON.songs {
		song := &rawJSON.songs[i]

		for j := range song.Cues {
			cue := &song.Cues[j]
			cue.SongID = song.ID

			if cue.TriggerLocationArea != nil {
				cue.TriggerAreaID, err = assignFKPtr(cue.TriggerLocationArea, lookup.L.Areas)
				if err != nil {
					return nil, err
				}
			}

			cues = append(cues, *cue)
		}
	}

	return dedupeRows(cues, hashes), nil
}

func getCues() []seedtypes.Cue {
	cues := []seedtypes.Cue{}

	for _, song := range rawJSON.songs {
		cues = append(cues, song.Cues...)
	}

	return cues
}

func getCueIncludedAreas(c seedtypes.Cue) ([]seedtypes.Area, error) {
	return lookup.GetResources(c.IncludedAreas, lookup.L.Areas)
}

func seedJuncCuesIncludedAreas(qtx *database.Queries, ctx context.Context) error {
	const desc string = "cues + included areas"
	jParams, err := processJunctions(desc, getCues(), getCueIncludedAreas)
	if err != nil {
		return err
	}

	return qtx.CreateCuesIncludedAreasJunctionBulk(ctx, database.CreateCuesIncludedAreasJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		CueID:          jParams.ParentIDs,
		IncludedAreaID: jParams.ChildIDs,
	})
}
