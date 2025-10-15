package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type MasterItem struct {
	ID		int32
	Name 	string 				`json:"name"`
	Type 	database.ItemType
}

func (i MasterItem) ToHashFields() []any {
	return []any{
		i.Name,
		i.Type,
	}
}

func (i MasterItem) ToKeyFields() []any {
	return []any{
		i.Name,
		i.Type,
	}
}

func (i MasterItem) GetID() int32 {
	return i.ID
}


func (l *lookup) seedMasterItem(qtx *database.Queries, masterItem MasterItem) (MasterItem, error) {
	dbMasterItem, err := qtx.CreateMasterItem(context.Background(), database.CreateMasterItemParams{
		DataHash: generateDataHash(masterItem),
		Name:     masterItem.Name,
		Type:     masterItem.Type,
	})
	if err != nil {
		return MasterItem{}, fmt.Errorf("couldn't create Master Item: %s: %v", masterItem.Name, err)
	}

	masterItem.ID = dbMasterItem.ID

	return masterItem, nil
}