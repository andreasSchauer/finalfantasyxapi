package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Item struct {
	ID int32
	MasterItem
	ItemAbility
	Description           string   `json:"description"`
	Effect                string   `json:"effect"`
	RelatedStats          []string `json:"related_stats"`
	SphereGridDescription *string  `json:"sphere_grid_description"`
	Category              string   `json:"category"`
	Usability             string   `json:"usability"`
	AvailableMenus        []string `json:"available_menus"`
	BasePrice             *int32   `json:"base_price"`
	SellValue             int32    `json:"sell_value"`
}

func (i Item) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", i),
		i.MasterItem.ID,
		i.Description,
		i.Effect,
		h.DerefOrNil(i.SphereGridDescription),
		i.Category,
		i.Usability,
		h.DerefOrNil(i.BasePrice),
		i.SellValue,
	}
}

func (i Item) GetID() int32 {
	return i.ID
}

func (i Item) Error() string {
	return fmt.Sprintf("item %s", i.Name)
}

func (i Item) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   i.ID,
		Name: i.Name,
	}
}

type ItemAbility struct {
	ID int32
	Ability
	ItemID             int32
	Cursor             string              `json:"cursor"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (a ItemAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.ItemID,
		a.Ability.ID,
		a.Cursor,
	}
}

func (a ItemAbility) GetID() int32 {
	return a.ID
}

func (a *ItemAbility) SetID(id int32) {
	a.ID = id
}

func (a ItemAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        a.Name,
		Version:     a.Version,
		AbilityType: string(database.AbilityTypeItemAbility),
	}
}

func (a ItemAbility) Error() string {
	return fmt.Sprintf("item ability %s", a.Name)
}

func (a ItemAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   a.ID,
		Name: a.Name,
	}
}

func (l *Lookup) loop2SeedItems(qtx *database.Queries, ctx context.Context) error {
	items, err := l.extractItems()

	params := database.CreateItemBulkParams{
		DataHash:     make([]string, len(items)),
		MasterItemID: make([]int32, len(items)),
		Description:  make([]string, len(items)),
		Effect:       make([]string, len(items)),
		Category:     make([]database.ItemCategory, len(items)),
		Usability:    make([]database.ItemUsability, len(items)),
		BasePrice:    make([]sql.NullInt32, len(items)),
		SellValue:    make([]int32, len(items)),
	}

	for i, item := range items {
		params.DataHash[i] = generateDataHash(item)
		params.MasterItemID[i] = item.MasterItem.ID
		params.Description[i] = item.Description
		params.Effect[i] = item.Effect
		params.Category[i] = database.ItemCategory(item.Category)
		params.Usability[i] = database.ItemUsability(item.Usability)
		params.BasePrice[i] = h.GetNullInt32(item.BasePrice)
		params.SellValue[i] = item.SellValue
	}

	dbRows, err := qtx.CreateItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create items: %v", err)
	}

	for i, row := range dbRows {
		items[i].ID = row.ID
		l.json.items[i].ID = row.ID
		l.Items[items[i].Name] = items[i]
		l.ItemsID[row.ID] = items[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractItems() ([]Item, error) {
	items := []Item{}
	var err error

	for i := range l.json.items {
		item := &l.json.items[i]

		item.MasterItem.ID, err = assignFK(item.Name, l.MasterItems)
		if err != nil {
			return nil, err
		}
		item.MasterItem.Type = database.ItemTypeItem

		items = append(items, *item)
	}

	return dedupeRows(items, l.Hashes), nil
}

func (l *Lookup) completeItems() error {
	for i := range l.json.items {
		item := &l.json.items[i]

		if len(item.BattleInteractions) > 0 {
			item.ItemAbility.ID = item.ID

			err := l.completeBattleInteractions(item.BattleInteractions)
			if err != nil {
				return err
			}

			l.ItemAbilities[Key(item)] = item.ItemAbility
			l.ItemAbilitiesID[item.ID] = item.ItemAbility
		}

		l.Items[item.Name] = *item
		l.ItemsID[item.ID] = *item
	}

	return nil
}

func (l *Lookup) getItemAvailableMenus(i Item) ([]Submenu, error) {
	return getResources(i.AvailableMenus, l.Submenus)
}

func (l *Lookup) getItemRelatedStats(i Item) ([]Stat, error) {
	return getResources(i.RelatedStats, l.Stats)
}

func (l *Lookup) seedJuncItemsAvailableMenus(qtx *database.Queries, ctx context.Context) error {
	const desc string = "items + available menus"
	jParams, err := processJunctions(l, desc, l.json.items, l.getItemAvailableMenus)
	if err != nil {
		return err
	}

	return qtx.CreateItemsAvailableMenusJunctionBulk(ctx, database.CreateItemsAvailableMenusJunctionBulkParams{
		DataHash:  jParams.DataHashes,
		ItemID:    jParams.ParentIDs,
		SubmenuID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncItemsRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "items + related stats"
	jParams, err := processJunctions(l, desc, l.json.items, l.getItemRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateItemsRelatedStatsJunctionBulk(ctx, database.CreateItemsRelatedStatsJunctionBulkParams{
		DataHash: jParams.DataHashes,
		ItemID:   jParams.ParentIDs,
		StatID:   jParams.ChildIDs,
	})
}

func (l *Lookup) loop3SeedItemAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractItemAbilities()
	if err != nil {
		return err
	}

	params := database.CreateItemAbilityBulkParams{
		DataHash:  make([]string, len(abilities)),
		ItemID:    make([]int32, len(abilities)),
		AbilityID: make([]int32, len(abilities)),
		Cursor:    make([]database.TargetType, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
		params.ItemID[i] = a.ItemID
		params.Cursor[i] = database.TargetType(a.Cursor)
	}

	dbRows, err := qtx.CreateItemAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create item abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		l.ItemAbilities[abilities[i].Name] = abilities[i]
		l.ItemAbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractItemAbilities() ([]ItemAbility, error) {
	abilities := []ItemAbility{}
	var err error

	for i := range l.json.items {
		item := &l.json.items[i]

		if len(item.BattleInteractions) == 0 {
			continue
		}

		ability := &item.ItemAbility
		ability.Name = item.Name
		ability.Type = database.AbilityTypeItemAbility
		ability.ItemID = item.ID

		ability.Ability.ID, err = l.getHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}
