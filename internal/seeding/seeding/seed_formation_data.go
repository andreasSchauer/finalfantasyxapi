package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop4SeedFormationData(qtx *database.Queries, ctx context.Context) error {
	data, err := extractFormationData()
	if err != nil {
		return err
	}

	params := database.CreateFormationDataBulkParams{
		DataHash:       make([]string, len(data)),
		Category:       make([]database.MonsterFormationCategory, len(data)),
		Availability:   make([]database.AvailabilityType, len(data)),
		IsForcedAmbush: make([]bool, len(data)),
		CanEscape:      make([]bool, len(data)),
		BossSongID:     make([]sql.NullInt32, len(data)),
		Notes:          make([]sql.NullString, len(data)),
	}

	for i, d := range data {
		params.DataHash[i] = generateDataHash(d)
		params.Category[i] = database.MonsterFormationCategory(d.Category)
		params.Availability[i] = database.AvailabilityType(d.Availability)
		params.IsForcedAmbush[i] = d.IsForcedAmbush
		params.CanEscape[i] = d.CanEscape
		params.BossSongID[i] = h.ObjPtrToNullInt32ID(d.BossMusic)
		params.Notes[i] = h.GetNullString(d.Notes)
	}

	dbRows, err := qtx.CreateFormationDataBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create formation data: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractFormationData() ([]seedtypes.FormationData, error) {
	data := []seedtypes.FormationData{}
	var err error

	for i := range rawJSON.monsterFormations {
		mfData := &rawJSON.monsterFormations[i].FormationData

		if mfData.BossMusic != nil {
			mfData.BossMusic.ID, err = GetHashID(mfData.BossMusic)
			if err != nil {
				return nil, err
			}
		}

		data = append(data, *mfData)
	}

	return dedupeRows(data, hashes), nil
}

func loop3SeedFormationBossSongs(qtx *database.Queries, ctx context.Context) error {
	songs, err := extractFormationBossSongs()
	if err != nil {
		return err
	}

	params := database.CreateFormationBossSongBulkParams{
		DataHash:         make([]string, len(songs)),
		SongID:           make([]int32, len(songs)),
		CelebrateVictory: make([]bool, len(songs)),
	}

	for i, s := range songs {
		params.DataHash[i] = generateDataHash(s)
		params.SongID[i] = s.SongID
		params.CelebrateVictory[i] = s.CelebrateVictory
	}

	dbRows, err := qtx.CreateFormationBossSongBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create formation boss songs: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractFormationBossSongs() ([]seedtypes.FormationBossSong, error) {
	songs := []seedtypes.FormationBossSong{}
	var err error

	for i := range rawJSON.monsterFormations {
		data := &rawJSON.monsterFormations[i].FormationData

		if data.BossMusic == nil {
			continue
		}

		data.BossMusic.SongID, err = assignFK(data.BossMusic.Song, lookup.L.Songs)
		if err != nil {
			return nil, err
		}

		songs = append(songs, *data.BossMusic)
	}

	return dedupeRows(songs, hashes), nil
}
