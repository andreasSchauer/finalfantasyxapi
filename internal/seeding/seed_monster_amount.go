package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type MonsterAmount struct {
	ID          int32
	MonsterID   int32
	MonsterName string `json:"monster_name"`
	Version     *int32 `json:"version"`
	Amount      int32  `json:"amount"`
}

func (ma MonsterAmount) ToHashFields() []any {
	return []any{
		ma.MonsterID,
		ma.Amount,
	}
}

func (ma MonsterAmount) ToKeyFields() []any {
	return []any{
		ma.MonsterName,
		derefOrNil(ma.Version),
	}
}

func (ma MonsterAmount) GetID() int32 {
	return ma.ID
}


func (l *lookup) seedMonsterAmount(qtx *database.Queries, monsterAmount MonsterAmount) (MonsterAmount, error) {
	dbMonsterAmount, err := qtx.CreateMonsterAmount(context.Background(), database.CreateMonsterAmountParams{
		DataHash:  generateDataHash(monsterAmount),
		MonsterID: monsterAmount.MonsterID,
		Amount:    monsterAmount.Amount,
	})
	if err != nil {
		return MonsterAmount{}, fmt.Errorf("couldn't create monster amount for %s: %v", createLookupKey(monsterAmount), err)
	}
	monsterAmount.ID = dbMonsterAmount.ID

	return monsterAmount, nil
}