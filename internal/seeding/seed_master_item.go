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


func (l *Lookup) loop1SeedMasterItems(qtx *database.Queries, ctx context.Context) error {
	items := l.extractMasterItems()

	params := database.CreateMasterItemBulkParams{
		DataHash: make([]string, len(items)),
		Name:     make([]string, len(items)),
		Type:     make([]database.ItemType, len(items)),
	}

	for i, mi := range items {
		params.DataHash[i] = generateDataHash(mi)
		params.Name[i] = mi.Name
		params.Type[i] = mi.Type
	}

	dbRows, err := qtx.CreateMasterItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create master items: %v", err)
	}

	for i, row := range dbRows {
		items[i].ID = row.ID
		key := CreateLookupKey(items[i])
		l.MasterItems[key] = items[i]
		l.MasterItemsID[row.ID] = items[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMasterItems() []MasterItem {
	masterItems := []MasterItem{}

	for _, i := range l.json.items {
		i.MasterItem.Type = database.ItemTypeItem
		masterItems = append(masterItems, i.MasterItem)
	}

	for _, i := range l.json.keyItems {
		i.MasterItem.Type = database.ItemTypeKeyItem
		masterItems = append(masterItems, i.MasterItem)
	}

	return dedupeRows(masterItems, l.Hashes)
}

