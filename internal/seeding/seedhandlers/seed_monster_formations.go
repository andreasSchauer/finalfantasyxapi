package seedhandlers

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

func Loop5SeedMonsterFormations(qtx *database.Queries, ctx context.Context) error {
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
		params.DataHash[i] = sh.GenerateDataHash(mf)
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
		sh.RawJSON.MonsterFormations[i].ID = row.ID
		lookup.L.MonsterFormations[seedtypes.Key(formations[i])] = formations[i]
		lookup.L.MonsterFormationsID[row.ID] = formations[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterFormations() ([]seedtypes.MonsterFormation, error) {
	formations := []seedtypes.MonsterFormation{}
	var err error

	for i := range sh.RawJSON.MonsterFormations {
		mf := &sh.RawJSON.MonsterFormations[i]

		mf.FormationData.ID, err = sh.GetHashID(mf.FormationData)
		if err != nil {
			return nil, err
		}

		mf.MonsterSelection.ID, err = sh.GetHashID(mf.MonsterSelection)
		if err != nil {
			return nil, err
		}

		formations = append(formations, *mf)
	}

	return sh.DedupeRows(formations, sh.Hashes), nil
}

func CompleteMonsterFormations() error {
	for i := range sh.RawJSON.MonsterFormations {
		formation := &sh.RawJSON.MonsterFormations[i]

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

		lookup.L.MonsterFormations[seedtypes.Key(*formation)] = *formation
		lookup.L.MonsterFormationsID[formation.ID] = *formation
	}

	return nil
}

func Loop1SeedMonsterSelections(qtx *database.Queries, ctx context.Context) error {
	selections, err := extractMonsterSelections()
	if err != nil {
		return err
	}

	dataHashes := make([]string, len(selections))

	for i, s := range selections {
		dataHashes[i] = sh.GenerateDataHash(s)
	}

	dbRows, err := qtx.CreateMonsterSelectionBulk(ctx, dataHashes)
	if err != nil {
		return fmt.Errorf("couldn't create monster selections: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterSelections() ([]seedtypes.MonsterSelection, error) {
	selections := []seedtypes.MonsterSelection{}
	var err error

	for i := range sh.RawJSON.MonsterFormations {
		mf := &sh.RawJSON.MonsterFormations[i]

		for j := range mf.MonsterSelection.Monsters {
			ma := &mf.MonsterSelection.Monsters[j]

			key := seedtypes.Key(*ma)
			ma.MonsterID, err = sh.AssignFK(key, lookup.L.Monsters)
			if err != nil {
				return nil, err
			}
		}

		selections = append(selections, mf.MonsterSelection)
	}

	return sh.DedupeRows(selections, sh.Hashes), nil
}
