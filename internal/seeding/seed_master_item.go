package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MasterItem struct {
	ID   int32
	Name string `json:"name"`
	Type database.ItemType
}

func (i MasterItem) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", i),
		i.Name,
		i.Type,
	}
}

func (i MasterItem) ToKeyFields() []any {
	return []any{
		i.Name,
	}
}

func (i MasterItem) GetID() int32 {
	return i.ID
}

func (i MasterItem) Error() string {
	return fmt.Sprintf("master item %s, type %s", i.Name, i.Type)
}

func (i MasterItem) GetResParamsTyped() h.ResParamsTyped {
	return h.ResParamsTyped{
		ID: 	i.ID,
		Name: 	i.Name,
		Type: 	string(i.Type),
	}
}

func (l *Lookup) seedMasterItem(qtx *database.Queries, masterItem MasterItem) (MasterItem, error) {
	dbMasterItem, err := qtx.CreateMasterItem(context.Background(), database.CreateMasterItemParams{
		DataHash: generateDataHash(masterItem),
		Name:     masterItem.Name,
		Type:     masterItem.Type,
	})
	if err != nil {
		return MasterItem{}, h.NewErr(masterItem.Error(), err)
	}

	masterItem.ID = dbMasterItem.ID
	l.MasterItems[masterItem.Name] = masterItem
	l.MasterItemsID[masterItem.ID] = masterItem

	return masterItem, nil
}
