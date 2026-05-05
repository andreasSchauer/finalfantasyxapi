package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop5SeedAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractAutoAbilities()
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
		params.DataHash[i] = generateDataHash(a)
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
		l.json.autoAbilities[i].ID = row.ID
		l.AutoAbilities[Key(abilities[i])] = abilities[i]
		l.AutoAbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAutoAbilities() ([]AutoAbility, error) {
	abilities := []AutoAbility{}
	var err error

	for i := range l.json.autoAbilities {
		ability := &l.json.autoAbilities[i]

		ability.GradRecoveryStatID, err = assignFKPtr(ability.GradualRecovery, l.Stats)
		if err != nil {
			return nil, err
		}

		ability.OnHitElementID, err = assignFKPtr(ability.OnHitElement, l.Elements)
		if err != nil {
			return nil, err
		}

		ability.AddedPropertyID, err = assignFKPtr(ability.AddedProperty, l.Properties)
		if err != nil {
			return nil, err
		}

		ability.CnvrsnFromModID, err = assignFKPtr(ability.ConversionFrom, l.Modifiers)
		if err != nil {
			return nil, err
		}

		ability.CnvrsnToModID, err = assignFKPtr(ability.ConversionTo, l.Modifiers)
		if err != nil {
			return nil, err
		}

		if ability.RequiredItem != nil {
			ability.RequiredItem.ID, err = l.getHashID(ability.RequiredItem)
			if err != nil {
				return nil, err
			}
		}

		if ability.AddedElemResist != nil {
			ability.AddedElemResist.ID, err = l.getHashID(ability.AddedElemResist)
			if err != nil {
				return nil, err
			}
		}

		if ability.OnHitStatus != nil {
			ability.OnHitStatus.ID, err = l.getHashID(ability.OnHitStatus)
			if err != nil {
				return nil, err
			}
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}

func (l *Lookup) completeAutoAbilities() error {
	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		err := assignIDs(l, autoAbility.AddedStatusResists)
		if err != nil {
			return err
		}

		err = assignIDs(l, autoAbility.StatChanges)
		if err != nil {
			return err
		}

		err = assignIDs(l, autoAbility.ModifierChanges)
		if err != nil {
			return err
		}

		l.AutoAbilities[Key(autoAbility)] = *autoAbility
		l.AutoAbilitiesID[autoAbility.ID] = *autoAbility
	}

	return nil
}
