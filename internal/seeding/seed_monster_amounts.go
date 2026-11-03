package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (l *Lookup) loop2SeedMonsterAmounts(qtx *database.Queries, ctx context.Context) error {
	mas, err := l.extractMonsterAmounts()
	if err != nil {
		return err
	}

	params := database.CreateMonsterAmountBulkParams{
		DataHash:  make([]string, len(mas)),
		MonsterID: make([]int32, len(mas)),
		Amount:    make([]int32, len(mas)),
	}

	for i, c := range mas {
		params.DataHash[i] = generateDataHash(c)
		params.MonsterID[i] = c.MonsterID
		params.Amount[i] = c.Amount
	}

	dbRows, err := qtx.CreateMonsterAmountBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster amounts: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterAmounts() ([]seedtypes.MonsterAmount, error) {
	mas := []seedtypes.MonsterAmount{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		for j := range mf.Monsters {
			ma := &mf.Monsters[j]

			key := seedtypes.Key(*ma)
			ma.MonsterID, err = assignFK(key, l.Monsters)
			if err != nil {
				return nil, err
			}

			mas = append(mas, *ma)
		}
	}

	return dedupeRows(mas, l.Hashes), nil
}
