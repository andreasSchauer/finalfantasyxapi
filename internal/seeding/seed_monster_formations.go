package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop5SeedMonsterFormations(qtx *database.Queries, ctx context.Context) error {
	formations, err := l.extractMonsterFormations()
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
		return fmt.Errorf("couldn't create monster formations: %v", err)
	}

	for i, row := range dbRows {
		formations[i].ID = row.ID
		l.json.monsterFormations[i].ID = row.ID
		l.MonsterFormations[Key(formations[i])] = formations[i]
		l.MonsterFormationsID[row.ID] = formations[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterFormations() ([]MonsterFormation, error) {
	formations := []MonsterFormation{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		mf.FormationData.ID, err = l.GetHashID(mf.FormationData)
		if err != nil {
			return nil, err
		}

		mf.MonsterSelection.ID, err = l.GetHashID(mf.MonsterSelection)
		if err != nil {
			return nil, err
		}

		formations = append(formations, *mf)
	}

	return dedupeRows(formations, l.Hashes), nil
}

func (l *Lookup) completeMonsterFormations() error {
	for i := range l.json.monsterFormations {
		formation := &l.json.monsterFormations[i]

		err := assignIDs(l, formation.Monsters)
		if err != nil {
			return err
		}

		err = assignIDs(l, formation.EncounterAreas)
		if err != nil {
			return err
		}

		err = assignIDs(l, formation.TriggerCommands)
		if err != nil {
			return err
		}

		l.MonsterFormations[Key(*formation)] = *formation
		l.MonsterFormationsID[formation.ID] = *formation
	}

	return nil
}

func (l *Lookup) loop1SeedMonsterSelections(qtx *database.Queries, ctx context.Context) error {
	selections, err := l.extractMonsterSelections()
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterSelections() ([]MonsterSelection, error) {
	selections := []MonsterSelection{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		for j := range mf.MonsterSelection.Monsters {
			ma := &mf.MonsterSelection.Monsters[j]

			key := Key(*ma)
			ma.MonsterID, err = assignFK(key, l.Monsters)
			if err != nil {
				return nil, err
			}
		}

		selections = append(selections, mf.MonsterSelection)
	}

	return dedupeRows(selections, l.Hashes), nil
}
