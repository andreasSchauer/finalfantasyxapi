package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type KeyItem struct {
	//id 			int32
	//dataHash		string
	MasterItem
	MasterItemsID int32
	Category      string `json:"category"`
	Description   string `json:"description"`
	Effect        string `json:"effect"`
}

func (k KeyItem) ToHashFields() []any {
	return []any{
		k.MasterItemsID,
		k.Category,
		k.Description,
		k.Effect,
	}
}

type KeyItemLookup struct {
	KeyItem
	ID 		int32
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
			keyItem.Type = database.ItemTypeKeyItem
			dbMasterItem, err := l.seedMasterItem(qtx, keyItem.MasterItem)
			if err != nil {
				return err
			}

			keyItem.MasterItemsID = dbMasterItem.ID

			dbKeyItem, err := qtx.CreateKeyItem(context.Background(), database.CreateKeyItemParams{
				DataHash:     generateDataHash(keyItem),
				MasterItemID: keyItem.MasterItemsID,
				Category:     database.KeyItemCategory(keyItem.Category),
				Description:  keyItem.Description,
				Effect:       keyItem.Effect,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Key Item: %s: %v", keyItem.Name, err)
			}

			key := createLookupKey(keyItem.MasterItem)
			l.keyItems[key] = KeyItemLookup{
				KeyItem: 	keyItem,
				ID: 		dbKeyItem.ID,
			}
		}
		return nil
	})
}
