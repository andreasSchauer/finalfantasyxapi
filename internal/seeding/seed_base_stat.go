package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type BaseStat struct {
	ID			int32
	StatID		int32
	StatName	string	`json:"name"`
	Value		int32	`json:"value"`
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


func (l *lookup) seedBaseStat(qtx *database.Queries, baseStat BaseStat) (BaseStat, error) {
	var err error

	baseStat.StatID, err = assignFK(baseStat.StatName, l.getStat)
	if err != nil {
		return BaseStat{}, err
	}
	
	dbBaseStat, err := qtx.CreateBaseStat(context.Background(), database.CreateBaseStatParams{
		DataHash: 	generateDataHash(baseStat),
		StatID: 	baseStat.StatID,
		Value: 		baseStat.Value,
	})
	if err != nil {
		return BaseStat{}, err
	}

	baseStat.ID = dbBaseStat.ID

	return baseStat, nil
}