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

func (i MasterItem) Error() string {
	return fmt.Sprintf("master item %s, type %s", i.Name, i.Type)
}

func (l *Lookup) seedMasterItem(qtx *database.Queries, masterItem MasterItem) (MasterItem, error) {
	dbMasterItem, err := qtx.CreateMasterItem(context.Background(), database.CreateMasterItemParams{
		DataHash: generateDataHash(masterItem),
		Name:     masterItem.Name,
		Type:     masterItem.Type,
	})
	if err != nil {
		return MasterItem{}, h.GetErr(masterItem.Error(), err)
	}

	masterItem.ID = dbMasterItem.ID
	l.masterItems[masterItem.Name] = masterItem

	return masterItem, nil
}
