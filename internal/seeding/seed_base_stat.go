package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type BaseStat struct {
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


func (l *lookup) seedBaseStat(qtx *database.Queries, baseStat BaseStat) (database.BaseStat, error) {
	stat, err := l.getStat(baseStat.StatName)
	if err != nil {
		return database.BaseStat{}, err
	}
	baseStat.StatID = stat.ID
	
	dbBaseStat, err := qtx.CreateBaseStat(context.Background(), database.CreateBaseStatParams{
		DataHash: 	generateDataHash(baseStat),
		StatID: 	baseStat.StatID,
		Value: 		baseStat.Value,
	})
	if err != nil {
		return database.BaseStat{}, err
	}

	return dbBaseStat, nil
}