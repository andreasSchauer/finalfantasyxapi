package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop2SeedMonsterAmounts(qtx *database.Queries, ctx context.Context) error {
	mas, err := extractMonsterAmounts()
	if err != nil {
		return err
	}

	params := database.CreateMonsterAmountBulkParams{
		DataHash:  make([]string, len(mas)),
		MonsterID: make([]int32, len(mas)),
		Amount:    make([]int32, len(mas)),
	}

	for i, c := range mas {
		params.DataHash[i] = sh.GenerateDataHash(c)
		params.MonsterID[i] = c.MonsterID
		params.Amount[i] = c.Amount
	}

	dbRows, err := qtx.CreateMonsterAmountBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster amounts: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterAmounts() ([]seedtypes.MonsterAmount, error) {
	mas := []seedtypes.MonsterAmount{}
	var err error

	for i := range sh.RawJSON.MonsterFormations {
		mf := &sh.RawJSON.MonsterFormations[i]

		for j := range mf.Monsters {
			ma := &mf.Monsters[j]

			key := seedtypes.Key(*ma)
			ma.MonsterID, err = sh.AssignFK(key, lookup.L.Monsters)
			if err != nil {
				return nil, err
			}

			mas = append(mas, *ma)
		}
	}

	return dedupeRows(mas, sh.Hashes), nil
}
