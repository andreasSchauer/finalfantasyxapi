package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Item struct {
	ID						int32
	MasterItem
	ItemAbility
	Description           	string  `json:"description"`
	Effect                	string  `json:"effect"`
	SphereGridDescription 	*string `json:"sphere_grid_description"`
	Category              	string  `json:"category"`
	Usability             	*string `json:"usability"`
	BasePrice             	*int32  `json:"base_price"`
	SellValue             	int32   `json:"sell_value"`
}

func (i Item) ToHashFields() []any {
	return []any{
		i.MasterItem.ID,
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
	ItemID    int32
	Cursor    string `json:"cursor"`
}

func (i ItemAbility) ToHashFields() []any {
	return []any{
		i.ItemID,
		i.Ability.ID,
		i.Cursor,
	}
}

func (l *lookup) seedItems(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/items.json"

	var items []Item
	err := loadJSONFile(string(srcPath), &items)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, item := range items {
			item.Type = database.ItemTypeItem

			dbMasterItem, err := l.seedMasterItem(qtx, item.MasterItem)
			if err != nil {
				return err
			}
			item.MasterItem.ID = dbMasterItem.ID

			dbItem, err := l.seedItem(qtx, item)
			if err != nil {
				return err
			}
			item.ID = dbItem.ID

			hasBattleData := item.Category != string(database.ItemCategorySphere) && item.Category != string(database.ItemCategoryOther)

			if hasBattleData {
				err = l.seedItemAbility(qtx, item)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}


func (l *lookup) seedItem(qtx *database.Queries, item Item) (database.Item, error) {
	dbItem, err := qtx.CreateItem(context.Background(), database.CreateItemParams{
		DataHash:              generateDataHash(item),
		MasterItemID:          item.MasterItem.ID,
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

	item.ID = dbItem.ID
	key := createLookupKey(item.MasterItem)
	l.items[key] = item

	return dbItem, nil
}


func (l *lookup) seedItemAbility(qtx *database.Queries, item Item) error {
	itemAbility := item.ItemAbility
	ability := itemAbility.Ability
	ability.Name = item.Name
	ability.Type = database.AbilityTypeItem

	dbAbility, err := l.seedAbility(qtx, ability)
	if err != nil {
		return err
	}

	itemAbility.ItemID = item.ID
	itemAbility.Ability.ID = dbAbility.ID

	err = qtx.CreateItemAbility(context.Background(), database.CreateItemAbilityParams{
		DataHash:  generateDataHash(itemAbility),
		ItemID:    itemAbility.ItemID,
		AbilityID: itemAbility.Ability.ID,
		Cursor:    database.TargetType(itemAbility.Cursor),
	})
	if err != nil {
		return fmt.Errorf("couldn't create Item Ability: %s: %v", itemAbility.Name, err)
	}

	return nil
}
