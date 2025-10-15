package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type KeyItem struct {
	ID int32
	MasterItem
	Category    string `json:"category"`
	Description string `json:"description"`
	Effect      string `json:"effect"`
}

func (k KeyItem) ToHashFields() []any {
	return []any{
		k.MasterItem.ID,
		k.Category,
		k.Description,
		k.Effect,
	}
}

func (k KeyItem) GetID() int32 {
	return k.ID
}

func (l *lookup) seedKeyItems(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/key_items.json"

	var keyItems []KeyItem
	err := loadJSONFile(string(srcPath), &keyItems)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, keyItem := range keyItems {
			var err error
			keyItem.Type = database.ItemTypeKeyItem

			keyItem.MasterItem, err = seedObjAssignFK(qtx, keyItem.MasterItem, l.seedMasterItem)
			if err != nil {
				return err
			}

			dbKeyItem, err := qtx.CreateKeyItem(context.Background(), database.CreateKeyItemParams{
				DataHash:     generateDataHash(keyItem),
				MasterItemID: keyItem.MasterItem.ID,
				Category:     database.KeyItemCategory(keyItem.Category),
				Description:  keyItem.Description,
				Effect:       keyItem.Effect,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Key Item: %s: %v", keyItem.Name, err)
			}

			keyItem.ID = dbKeyItem.ID
			key := createLookupKey(keyItem.MasterItem)
			l.keyItems[key] = keyItem
		}
		return nil
	})
}
