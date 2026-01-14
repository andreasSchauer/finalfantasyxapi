package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type BaseStat struct {
	ID       int32
	StatID   int32
	StatName string `json:"name"`
	Value    int32  `json:"value"`
}

func (bs BaseStat) ToHashFields() []any {
	return []any{
		bs.StatID,
		bs.Value,
	}
}

func (bs BaseStat) GetID() int32 {
	return bs.ID
}

func (bs BaseStat) Error() string {
	return fmt.Sprintf("base stat %s, value: %d", bs.StatName, bs.Value)
}

func (l *Lookup) seedBaseStat(qtx *database.Queries, baseStat BaseStat) (BaseStat, error) {
	var err error

	baseStat.StatID, err = assignFK(baseStat.StatName, l.Stats)
	if err != nil {
		return BaseStat{}, h.NewErr(baseStat.Error(), err)
	}

	dbBaseStat, err := qtx.CreateBaseStat(context.Background(), database.CreateBaseStatParams{
		DataHash: generateDataHash(baseStat),
		StatID:   baseStat.StatID,
		Value:    baseStat.Value,
	})
	if err != nil {
		return BaseStat{}, h.NewErr(baseStat.Error(), err, "couldn't create base stat")
	}

	baseStat.ID = dbBaseStat.ID

	return baseStat, nil
}
