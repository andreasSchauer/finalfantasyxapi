package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Item struct {
	ID int32
	MasterItem
	ItemAbility
	Description           	string  	`json:"description"`
	Effect               	string  	`json:"effect"`
	RelatedStats			[]string	`json:"related_stats"`
	SphereGridDescription 	*string 	`json:"sphere_grid_description"`
	Category              	string  	`json:"category"`
	Usability             	*string 	`json:"usability"`
	AvailableMenus			[]string	`json:"available_menus"`
	BasePrice             	*int32  	`json:"base_price"`
	SellValue             	int32   	`json:"sell_value"`
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

func (i Item) GetID() int32 {
	return i.ID
}

type ItemAbility struct {
	ID					int32
	Ability
	ItemID 				int32
	Cursor 				string `json:"cursor"`
	BattleInteractions  []BattleInteraction `json:"battle_interactions"`
}

func (a ItemAbility) ToHashFields() []any {
	return []any{
		a.ItemID,
		a.Ability.ID,
		a.Cursor,
	}
}

func (a ItemAbility) GetID() int32 {
	return a.ID
}

func (a ItemAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        a.Name,
		Version:     a.Version,
		AbilityType: string(database.AbilityTypeItemAbility),
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
			var err error
			item.Type = database.ItemTypeItem

			item.MasterItem, err = seedObjAssignID(qtx, item.MasterItem, l.seedMasterItem)
			if err != nil {
				return err
			}

			item, err = seedObjAssignID(qtx, item, l.seedItem)
			if err != nil {
				return err
			}

			if len(item.BattleInteractions) > 0 {
				err = l.seedItemAbility(qtx, item)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (l *lookup) seedItem(qtx *database.Queries, item Item) (Item, error) {
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
		return Item{}, fmt.Errorf("couldn't create Item: %s: %v", item.Name, err)
	}

	item.ID = dbItem.ID
	key := createLookupKey(item.MasterItem)
	l.items[key] = item

	return item, nil
}

func (l *lookup) seedItemAbility(qtx *database.Queries, item Item) error {
	var err error
	itemAbility := item.ItemAbility
	itemAbility.Name = item.Name
	itemAbility.Type = database.AbilityTypeItemAbility
	itemAbility.ItemID = item.ID

	itemAbility.Ability, err = seedObjAssignID(qtx, itemAbility.Ability, l.seedAbility)
	if err != nil {
		return err
	}

	dbItemAbility, err := qtx.CreateItemAbility(context.Background(), database.CreateItemAbilityParams{
		DataHash:  generateDataHash(itemAbility),
		ItemID:    itemAbility.ItemID,
		AbilityID: itemAbility.Ability.ID,
		Cursor:    database.TargetType(itemAbility.Cursor),
	})
	if err != nil {
		return fmt.Errorf("couldn't create Item Ability: %s: %v", itemAbility.Name, err)
	}

	itemAbility.ID = dbItemAbility.ID

	item.ItemAbility = itemAbility
	key := createLookupKey(item.MasterItem)
	l.items[key] = item

	return nil
}


func (l *lookup) createItemsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/items.json"

	var items []Item
	err := loadJSONFile(string(srcPath), &items)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonItem := range items {
			item, err := l.getItem(jsonItem.Name)
			if err != nil {
				return err
			}

			err = l.seedItemRelatedStats(qtx, item)
			if err != nil {
				return fmt.Errorf("item %s: %v", item.Name, err)
			}

			err = l.seedItemAvailableMenus(qtx, item)
			if err != nil {
				return fmt.Errorf("item %s: %v", item.Name, err)
			}

			if len(item.BattleInteractions) > 0 {
				l.currentAbility = item.Ability

				err = l.seedBattleInteractions(qtx, l.currentAbility, item.BattleInteractions)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}


func (l *lookup) seedItemRelatedStats(qtx *database.Queries, item Item) error {
	for _, jsonStat := range item.RelatedStats {
		junction, err := createJunction(item, jsonStat, l.getStat)
		if err != nil {
			return err
		}

		err = qtx.CreateItemsRelatedStatsJunction(context.Background(), database.CreateItemsRelatedStatsJunctionParams{
			DataHash: 	generateDataHash(junction),
			ItemID: 	junction.ParentID,
			StatID: 	junction.ChildID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}


func (l *lookup) seedItemAvailableMenus(qtx *database.Queries, item Item) error {
	for _, jsonSubmenu := range item.AvailableMenus {
		junction, err := createJunction(item, jsonSubmenu, l.getSubmenu)
		if err != nil {
			return err
		}

		err = qtx.CreateItemsAvailableMenusJunction(context.Background(), database.CreateItemsAvailableMenusJunctionParams{
			DataHash: 	generateDataHash(junction),
			ItemID: 	junction.ParentID,
			SubmenuID: 	junction.ChildID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}