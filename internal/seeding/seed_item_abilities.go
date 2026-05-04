package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

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
