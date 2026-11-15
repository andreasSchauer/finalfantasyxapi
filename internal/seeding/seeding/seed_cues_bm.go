package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1SeedBackgroundMusic(qtx *database.Queries, ctx context.Context) error {
	bm := extractBackgroundMusic()

	params := database.CreateBackgroundMusicBulkParams{
		DataHash:               make([]string, len(bm)),
		Condition:              make([]sql.NullString, len(bm)),
		ReplacesEncounterMusic: make([]bool, len(bm)),
	}

	for i, music := range bm {
		params.DataHash[i] = sh.GenerateDataHash(music)
		params.Condition[i] = h.GetNullString(music.Condition)
		params.ReplacesEncounterMusic[i] = music.ReplacesEncounterMusic
	}

	dbRows, err := qtx.CreateBackgroundMusicBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create background music: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractBackgroundMusic() []seedtypes.BackgroundMusic {
	bgMusic := []seedtypes.BackgroundMusic{}

	for _, song := range sh.RawJSON.Songs {
		bgMusic = append(bgMusic, song.BackgroundMusic...)
	}

	return dedupeRows(bgMusic, sh.Hashes)
}

func Loop4SeedCues(qtx *database.Queries, ctx context.Context) error {
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
		params.DataHash[i] = sh.GenerateDataHash(c)
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
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractCues() ([]seedtypes.Cue, error) {
	cues := []seedtypes.Cue{}
	var err error

	for i := range sh.RawJSON.Songs {
		song := &sh.RawJSON.Songs[i]

		for j := range song.Cues {
			cue := &song.Cues[j]
			cue.SongID = song.ID

			if cue.TriggerLocationArea != nil {
				cue.TriggerAreaID, err = sh.AssignFKPtr(cue.TriggerLocationArea, lookup.L.Areas)
				if err != nil {
					return nil, err
				}
			}

			cues = append(cues, *cue)
		}
	}

	return dedupeRows(cues, sh.Hashes), nil
}
