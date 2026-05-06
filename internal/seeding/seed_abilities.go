package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop2SeedAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractAbilities()
	if err != nil {
		return err
	}

	params := database.CreateAbilityBulkParams{
		DataHash:      make([]string, len(abilities)),
		Name:          make([]string, len(abilities)),
		Version:       make([]sql.NullInt32, len(abilities)),
		Specification: make([]sql.NullString, len(abilities)),
		AttributesID:  make([]int32, len(abilities)),
		Type:          make([]database.AbilityType, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.Name[i] = a.Name
		params.Version[i] = h.GetNullInt32(a.Version)
		params.Specification[i] = h.GetNullString(a.Specification)
		params.AttributesID[i] = a.Attributes.ID
		params.Type[i] = a.Type
	}

	dbRows, err := qtx.CreateAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		l.Abilities[Key(abilities[i])] = abilities[i]
		l.AbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAbilities() ([]Ability, error) {
	abilities := []Ability{}
	var err error

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]
		ability.Attributes.ID, err = l.getHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypePlayerAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]
		ability.Attributes.ID, err = l.getHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypeOverdriveAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		if len(item.BattleInteractions) == 0 {
			continue
		}

		item.Attributes.ID, err = l.getHashID(item.Attributes)
		if err != nil {
			return nil, err
		}
		item.Ability.Name = item.Name
		item.Ability.Type = database.AbilityTypeItemAbility
		abilities = append(abilities, item.Ability)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]
		command.Attributes.ID, err = l.getHashID(command.Attributes)
		if err != nil {
			return nil, err
		}

		command.Type = database.AbilityTypeTriggerCommand
		abilities = append(abilities, command.Ability)
	}

	for i := range l.json.miscAbilities {
		ability := &l.json.miscAbilities[i]
		ability.Attributes.ID, err = l.getHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypeMiscAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]
		ability.Attributes.ID, err = l.getHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypeEnemyAbility
		abilities = append(abilities, ability.Ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}

func (l *Lookup) loop1SeedAbilityAttributes(qtx *database.Queries, ctx context.Context) error {
	attributes := l.extractAbilityAttributes()

	params := database.CreateAbilityAttributesBulkParams{
		DataHash:         make([]string, len(attributes)),
		Rank:             make([]sql.NullInt32, len(attributes)),
		AppearsInHelpBar: make([]bool, len(attributes)),
		CanCopycat:       make([]bool, len(attributes)),
	}

	for i, a := range attributes {
		params.DataHash[i] = generateDataHash(a)
		params.Rank[i] = h.GetNullInt32(a.Rank)
		params.AppearsInHelpBar[i] = a.AppearsInHelpBar
		params.CanCopycat[i] = a.CanCopycat
	}

	dbRows, err := qtx.CreateAbilityAttributesBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability attributes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAbilityAttributes() []Attributes {
	attributes := []Attributes{}

	for _, ability := range l.json.enemyAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.items {
		if len(ability.BattleInteractions) == 0 {
			continue
		}
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.overdrives {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.overdriveAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.playerAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.triggerCommands {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.miscAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	return dedupeRows(attributes, l.Hashes)
}
