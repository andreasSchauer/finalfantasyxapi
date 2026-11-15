package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop2SeedAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := extractAbilities()
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
		params.DataHash[i] = seedhelpers.GenerateDataHash(a)
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
		lookup.L.Abilities[seedtypes.Key(abilities[i])] = abilities[i]
		lookup.L.AbilitiesID[row.ID] = abilities[i]
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractAbilities() ([]seedtypes.Ability, error) {
	abilities := []seedtypes.Ability{}
	var err error

	for i := range seedhelpers.RawJSON.PlayerAbilities {
		ability := &seedhelpers.RawJSON.PlayerAbilities[i]
		ability.Attributes.ID, err = seedhelpers.GetHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypePlayerAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range seedhelpers.RawJSON.OverdriveAbilities {
		ability := &seedhelpers.RawJSON.OverdriveAbilities[i]
		ability.Attributes.ID, err = seedhelpers.GetHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypeOverdriveAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range seedhelpers.RawJSON.Items {
		item := &seedhelpers.RawJSON.Items[i]

		if len(item.BattleInteractions) == 0 {
			continue
		}

		item.Attributes.ID, err = seedhelpers.GetHashID(item.Attributes)
		if err != nil {
			return nil, err
		}
		item.Ability.Name = item.Name
		item.Ability.Type = database.AbilityTypeItemAbility
		abilities = append(abilities, item.Ability)
	}

	for i := range seedhelpers.RawJSON.TriggerCommands {
		command := &seedhelpers.RawJSON.TriggerCommands[i]
		command.Attributes.ID, err = seedhelpers.GetHashID(command.Attributes)
		if err != nil {
			return nil, err
		}

		command.Type = database.AbilityTypeTriggerCommand
		abilities = append(abilities, command.Ability)
	}

	for i := range seedhelpers.RawJSON.MiscAbilities {
		ability := &seedhelpers.RawJSON.MiscAbilities[i]
		ability.Attributes.ID, err = seedhelpers.GetHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypeMiscAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range seedhelpers.RawJSON.EnemyAbilities {
		ability := &seedhelpers.RawJSON.EnemyAbilities[i]
		ability.Attributes.ID, err = seedhelpers.GetHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypeEnemyAbility
		abilities = append(abilities, ability.Ability)
	}

	return dedupeRows(abilities, seedhelpers.Hashes), nil
}

func Loop1SeedAbilityAttributes(qtx *database.Queries, ctx context.Context) error {
	attributes := extractAbilityAttributes()

	params := database.CreateAbilityAttributesBulkParams{
		DataHash:         make([]string, len(attributes)),
		Rank:             make([]sql.NullInt32, len(attributes)),
		AppearsInHelpBar: make([]bool, len(attributes)),
		CanCopycat:       make([]bool, len(attributes)),
	}

	for i, a := range attributes {
		params.DataHash[i] = seedhelpers.GenerateDataHash(a)
		params.Rank[i] = h.GetNullInt32(a.Rank)
		params.AppearsInHelpBar[i] = a.AppearsInHelpBar
		params.CanCopycat[i] = a.CanCopycat
	}

	dbRows, err := qtx.CreateAbilityAttributesBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability attributes: %v", err)
	}

	for _, row := range dbRows {
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractAbilityAttributes() []seedtypes.Attributes {
	attributes := []seedtypes.Attributes{}

	for _, ability := range seedhelpers.RawJSON.EnemyAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range seedhelpers.RawJSON.Items {
		if len(ability.BattleInteractions) == 0 {
			continue
		}
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range seedhelpers.RawJSON.Overdrives {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range seedhelpers.RawJSON.OverdriveAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range seedhelpers.RawJSON.PlayerAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range seedhelpers.RawJSON.TriggerCommands {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range seedhelpers.RawJSON.MiscAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	return dedupeRows(attributes, seedhelpers.Hashes)
}
