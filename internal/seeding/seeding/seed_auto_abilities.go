package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop5SeedAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := extractAutoAbilities()
	if err != nil {
		return err
	}

	params := database.CreateAutoAbilityBulkParams{
		DataHash:             make([]string, len(abilities)),
		Name:                 make([]string, len(abilities)),
		Description:          make([]sql.NullString, len(abilities)),
		Effect:               make([]string, len(abilities)),
		Type:                 make([]database.EquipType, len(abilities)),
		Category:             make([]database.AutoAbilityCategory, len(abilities)),
		AbilityValue:         make([]sql.NullInt32, len(abilities)),
		ActivationCondition:  make([]database.AaActivationCondition, len(abilities)),
		Counter:              make([]database.NullCounterType, len(abilities)),
		RequiredItemAmountID: make([]sql.NullInt32, len(abilities)),
		GradRcvryStatID:      make([]sql.NullInt32, len(abilities)),
		OnHitElementID:       make([]sql.NullInt32, len(abilities)),
		AddedElemResistID:    make([]sql.NullInt32, len(abilities)),
		OnHitStatusID:        make([]sql.NullInt32, len(abilities)),
		AddedPropertyID:      make([]sql.NullInt32, len(abilities)),
		CnvrsnFromModID:      make([]sql.NullInt32, len(abilities)),
		CnvrsnToModID:        make([]sql.NullInt32, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = sh.GenerateDataHash(a)
		params.Name[i] = a.Name
		params.Description[i] = h.GetNullString(a.Description)
		params.Effect[i] = a.Effect
		params.Type[i] = database.EquipType(a.Type)
		params.Category[i] = database.AutoAbilityCategory(a.Category)
		params.AbilityValue[i] = h.GetNullInt32(a.AbilityValue)
		params.ActivationCondition[i] = database.AaActivationCondition(a.ActivationCondition)
		params.Counter[i] = database.ToNullCounterType(a.Counter)
		params.RequiredItemAmountID[i] = h.ObjPtrToNullInt32ID(a.RequiredItem)
		params.GradRcvryStatID[i] = h.GetNullInt32(a.GradRecoveryStatID)
		params.OnHitElementID[i] = h.GetNullInt32(a.OnHitElementID)
		params.AddedElemResistID[i] = h.ObjPtrToNullInt32ID(a.AddedElemResist)
		params.OnHitStatusID[i] = h.ObjPtrToNullInt32ID(a.OnHitStatus)
		params.AddedPropertyID[i] = h.GetNullInt32(a.AddedPropertyID)
		params.CnvrsnFromModID[i] = h.GetNullInt32(a.CnvrsnFromModID)
		params.CnvrsnToModID[i] = h.GetNullInt32(a.CnvrsnToModID)
	}

	dbRows, err := qtx.CreateAutoAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create auto-abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		sh.RawJSON.AutoAbilities[i].ID = row.ID
		lookup.L.AutoAbilities[seedtypes.Key(abilities[i])] = abilities[i]
		lookup.L.AutoAbilitiesID[row.ID] = abilities[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractAutoAbilities() ([]seedtypes.AutoAbility, error) {
	abilities := []seedtypes.AutoAbility{}
	var err error

	for i := range sh.RawJSON.AutoAbilities {
		ability := &sh.RawJSON.AutoAbilities[i]

		ability.GradRecoveryStatID, err = sh.AssignFKPtr(ability.GradualRecovery, lookup.L.Stats)
		if err != nil {
			return nil, err
		}

		ability.OnHitElementID, err = sh.AssignFKPtr(ability.OnHitElement, lookup.L.Elements)
		if err != nil {
			return nil, err
		}

		ability.AddedPropertyID, err = sh.AssignFKPtr(ability.AddedProperty, lookup.L.Properties)
		if err != nil {
			return nil, err
		}

		ability.CnvrsnFromModID, err = sh.AssignFKPtr(ability.ConversionFrom, lookup.L.Modifiers)
		if err != nil {
			return nil, err
		}

		ability.CnvrsnToModID, err = sh.AssignFKPtr(ability.ConversionTo, lookup.L.Modifiers)
		if err != nil {
			return nil, err
		}

		if ability.RequiredItem != nil {
			ability.RequiredItem.ID, err = sh.GetHashID(ability.RequiredItem)
			if err != nil {
				return nil, err
			}
		}

		if ability.AddedElemResist != nil {
			ability.AddedElemResist.ID, err = sh.GetHashID(ability.AddedElemResist)
			if err != nil {
				return nil, err
			}
		}

		if ability.OnHitStatus != nil {
			ability.OnHitStatus.ID, err = sh.GetHashID(ability.OnHitStatus)
			if err != nil {
				return nil, err
			}
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, sh.Hashes), nil
}

func completeAutoAbilities() error {
	for i := range sh.RawJSON.AutoAbilities {
		autoAbility := &sh.RawJSON.AutoAbilities[i]

		err := assignIDs(autoAbility.AddedStatusResists)
		if err != nil {
			return err
		}

		err = assignIDs(autoAbility.StatChanges)
		if err != nil {
			return err
		}

		err = assignIDs(autoAbility.ModifierChanges)
		if err != nil {
			return err
		}

		lookup.L.AutoAbilities[seedtypes.Key(autoAbility)] = *autoAbility
		lookup.L.AutoAbilitiesID[autoAbility.ID] = *autoAbility
	}

	return nil
}
