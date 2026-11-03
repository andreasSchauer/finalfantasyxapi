package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop5SeedMonsterFormations(qtx *database.Queries, ctx context.Context) error {
	formations, err := extractMonsterFormations()
	if err != nil {
		return err
	}

	params := database.CreateMonsterFormationBulkParams{
		DataHash:           make([]string, len(formations)),
		Version:            make([]sql.NullInt32, len(formations)),
		MonsterSelectionID: make([]int32, len(formations)),
		FormationDataID:    make([]int32, len(formations)),
	}

	for i, mf := range formations {
		params.DataHash[i] = generateDataHash(mf)
		params.Version[i] = h.GetNullInt32(mf.Version)
		params.MonsterSelectionID[i] = mf.MonsterSelection.ID
		params.FormationDataID[i] = mf.FormationData.ID
	}

	dbRows, err := qtx.CreateMonsterFormationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster-formations: %v", err)
	}

	for i, row := range dbRows {
		formations[i].ID = row.ID
		rawJSON.monsterFormations[i].ID = row.ID
		L.MonsterFormations[seedtypes.Key(formations[i])] = formations[i]
		L.MonsterFormationsID[row.ID] = formations[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterFormations() ([]seedtypes.MonsterFormation, error) {
	formations := []seedtypes.MonsterFormation{}
	var err error

	for i := range rawJSON.monsterFormations {
		mf := &rawJSON.monsterFormations[i]

		mf.FormationData.ID, err = GetHashID(mf.FormationData)
		if err != nil {
			return nil, err
		}

		mf.MonsterSelection.ID, err = GetHashID(mf.MonsterSelection)
		if err != nil {
			return nil, err
		}

		formations = append(formations, *mf)
	}

	return dedupeRows(formations, hashes), nil
}

func completeMonsterFormations() error {
	for i := range rawJSON.monsterFormations {
		formation := &rawJSON.monsterFormations[i]

		err := assignIDs(formation.Monsters)
		if err != nil {
			return err
		}

		err = assignIDs(formation.EncounterAreas)
		if err != nil {
			return err
		}

		err = assignIDs(formation.TriggerCommands)
		if err != nil {
			return err
		}

		L.MonsterFormations[seedtypes.Key(*formation)] = *formation
		L.MonsterFormationsID[formation.ID] = *formation
	}

	return nil
}

func loop1SeedMonsterSelections(qtx *database.Queries, ctx context.Context) error {
	selections, err := extractMonsterSelections()
	if err != nil {
		return err
	}

	dataHashes := make([]string, len(selections))

	for i, s := range selections {
		dataHashes[i] = generateDataHash(s)
	}

	dbRows, err := qtx.CreateMonsterSelectionBulk(ctx, dataHashes)
	if err != nil {
		return fmt.Errorf("couldn't create monster selections: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterSelections() ([]seedtypes.MonsterSelection, error) {
	selections := []seedtypes.MonsterSelection{}
	var err error

	for i := range rawJSON.monsterFormations {
		mf := &rawJSON.monsterFormations[i]

		for j := range mf.MonsterSelection.Monsters {
			ma := &mf.MonsterSelection.Monsters[j]

			key := seedtypes.Key(*ma)
			ma.MonsterID, err = assignFK(key, L.Monsters)
			if err != nil {
				return nil, err
			}
		}

		selections = append(selections, mf.MonsterSelection)
	}

	return dedupeRows(selections, hashes), nil
}
