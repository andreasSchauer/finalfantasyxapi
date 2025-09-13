package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type MasterItem struct {
	Name string 			`json:"name"`
	Type database.ItemType
}

func (i MasterItem) ToHashFields() []any {
	return []any{
		i.Name,
		i.Type,
	}
}


func seedMasterItem(qtx *database.Queries, item MasterItem) (database.MasterItem, error) {
	dbMasterItem, err := qtx.CreateMasterItem(context.Background(), database.CreateMasterItemParams{
		DataHash: generateDataHash(item),
		Name:     item.Name,
		Type:     item.Type,
	})
	if err != nil {
		return database.MasterItem{}, fmt.Errorf("couldn't create Master Item: %s: %v", item.Name, err)
	}

	return dbMasterItem, nil
}