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
		fmt.Sprintf("%T", ma),
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

func (ma *MonsterAmount) SetID(id int32) {
	ma.ID = id
}

func (ma MonsterAmount) GetName() string {
	return ma.MonsterName
}

func (ma MonsterAmount) GetVersion() *int32 {
	return ma.Version
}

func (ma MonsterAmount) GetVal() int32 {
	return ma.Amount
}

func (ma MonsterAmount) Error() string {
	return fmt.Sprintf("monster amount with monster: '%s', amount: %d", h.NameToString(ma.MonsterName, ma.Version, nil), ma.Amount)
}

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

func (l *Lookup) extractMonsterAmounts() ([]MonsterAmount, error) {
	mas := []MonsterAmount{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		for j := range mf.Monsters {
			ma := &mf.Monsters[j]

			key := Key(*ma)
			ma.MonsterID, err = assignFK(key, l.Monsters)
			if err != nil {
				return nil, err
			}

			mas = append(mas, *ma)
		}
	}

	return dedupeRows(mas, l.Hashes), nil
}