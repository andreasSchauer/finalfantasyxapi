package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Item struct {
	//id 			int32
	//dataHash		string
	MasterItem
	ItemAbility
	MasterItemsID         int32
	Description           string  `json:"description"`
	Effect                string  `json:"effect"`
	SphereGridDescription *string `json:"sphere_grid_description"`
	Category              string  `json:"category"`
	Usability             *string `json:"usability"`
	BasePrice             *int32  `json:"base_price"`
	SellValue             int32   `json:"sell_value"`
}

func (i Item) ToHashFields() []any {
	return []any{
		i.MasterItemsID,
		i.Description,
		i.Effect,
		derefOrNil(i.SphereGridDescription),
		i.Category,
		derefOrNil(i.Usability),
		derefOrNil(i.BasePrice),
		i.SellValue,
	}
}

type ItemAbility struct {
	Ability
	AbilityAttributes
	ItemID    int32
	AbilityID int32
	Cursor    string `json:"cursor"`
}

func (i ItemAbility) ToHashFields() []any {
	return []any{
		i.ItemID,
		i.AbilityID,
		i.Cursor,
	}
}


func seedItems(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/items.json"

	var items []Item
	err := loadJSONFile(string(srcPath), &items)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, item := range items {
			item.Type = database.ItemTypeItem

			dbMasterItem, err := seedMasterItem(qtx, item.MasterItem)
			if err != nil {
				return err
			}

			dbItem, err := seedItem(qtx, item, dbMasterItem.ID)
			if err != nil {
				return err
			}

			hasBattleData := item.Category != string(database.ItemCategorySphere) && item.Category != string(database.ItemCategoryOther)

			if hasBattleData {
				err = seedItemAbility(qtx, item, dbItem.ID)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}



func seedItem(qtx *database.Queries, item Item, allItemID int32) (database.Item, error) {
	item.MasterItemsID = allItemID

	dbItem, err := qtx.CreateItem(context.Background(), database.CreateItemParams{
		DataHash:              generateDataHash(item),
		MasterItemID:         item.MasterItemsID,
		Description:           item.Description,
		Effect:                item.Effect,
		SphereGridDescription: getNullString(item.SphereGridDescription),
		Category:              database.ItemCategory(item.Category),
		Usability:             nullItemUsability(item.Usability),
		BasePrice:             getNullInt32(item.BasePrice),
		SellValue:             item.SellValue,
	})
	if err != nil {
		return database.Item{}, fmt.Errorf("couldn't create Item: %s: %v", item.Name, err)
	}

	return dbItem, nil
}



func seedItemAbility(qtx *database.Queries, item Item, itemID int32) error {
	itemAbility := item.ItemAbility
	ability := itemAbility.Ability
	ability.Name = item.Name
	ability.Type = database.AbilityTypeItem

	dbAbility, err := seedAbility(qtx, item.AbilityAttributes, ability)
	if err != nil {
		return err
	}

	itemAbility.ItemID = itemID
	itemAbility.AbilityID = dbAbility.ID

	err = qtx.CreateItemAbility(context.Background(), database.CreateItemAbilityParams{
		DataHash:  generateDataHash(itemAbility),
		ItemID:    itemAbility.ItemID,
		AbilityID: itemAbility.AbilityID,
		Cursor:    database.TargetType(itemAbility.Cursor),
	})
	if err != nil {
		return fmt.Errorf("couldn't create Item Ability: %s: %v", itemAbility.Name, err)
	}

	return nil
}
