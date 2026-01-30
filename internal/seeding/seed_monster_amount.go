package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterAmount struct {
	ID          int32	`json:"-"`
	MonsterID   int32	`json:"-"`
	MonsterName string 	`json:"monster_name"`
	Version     *int32 	`json:"version"`
	Amount      int32  	`json:"amount"`
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
		h.DerefOrNil(ma.Version),
	}
}

func (ma MonsterAmount) ToFormationHashFields() []any {
	return []any{
		ma.MonsterName,
		h.DerefOrNil(ma.Version),
		ma.Amount,
	}
}

func (ma MonsterAmount) GetID() int32 {
	return ma.ID
}

func (ma MonsterAmount) Error() string {
	return fmt.Sprintf("monster amount with monster: %s, version: %v, amount: %d", ma.MonsterName, h.DerefOrNil(ma.Version), ma.Amount)
}

func (l *Lookup) seedMonsterAmount(qtx *database.Queries, monsterAmount MonsterAmount) (MonsterAmount, error) {
	dbMonsterAmount, err := qtx.CreateMonsterAmount(context.Background(), database.CreateMonsterAmountParams{
		DataHash:  generateDataHash(monsterAmount),
		MonsterID: monsterAmount.MonsterID,
		Amount:    monsterAmount.Amount,
	})
	if err != nil {
		return MonsterAmount{}, h.NewErr(monsterAmount.Error(), err, "couldn't create monster amount")
	}
	monsterAmount.ID = dbMonsterAmount.ID

	return monsterAmount, nil
}
