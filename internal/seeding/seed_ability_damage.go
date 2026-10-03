package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop5SeedAbilityDamages(qtx *database.Queries, ctx context.Context) error {
	damages, err := l.extractAbilityDamages()
	if err != nil {
		return err
	}

	params := database.CreateAbilityDamageBulkParams{
		DataHash:       make([]string, len(damages)),
		Condition:      make([]sql.NullString, len(damages)),
		AttackType:     make([]database.AttackType, len(damages)),
		TargetClass:    make([]database.TargetClass, len(damages)),
		DamageType:     make([]database.DamageType, len(damages)),
		DamageFormula:  make([]database.DamageFormula, len(damages)),
		DamageConstant: make([]int32, len(damages)),
	}

	for i, d := range damages {
		params.DataHash[i] = generateDataHash(d)
		params.Condition[i] = h.GetNullString(d.Condition)
		params.AttackType[i] = database.AttackType(d.AttackType)
		params.TargetClass[i] = database.TargetClass(d.TargetClass)
		params.DamageType[i] = database.DamageType(d.DamageType)
		params.DamageFormula[i] = database.DamageFormula(d.DamageFormula)
		params.DamageConstant[i] = d.DamageConstant
	}

	dbRows, err := qtx.CreateAbilityDamageBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability damages: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAbilityDamages() ([]AbilityDamage, error) {
	damages := []AbilityDamage{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		newDamages, err := l.getAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}
		damages = append(damages, newDamages...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		newDamages, err := l.getAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		newDamages, err := l.getAbilityDamages(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		newDamages, err := l.getAbilityDamages(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.miscAbilities {
		ability := &l.json.miscAbilities[i]

		newDamages, err := l.getAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		newDamages, err := l.getAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	return dedupeRows(damages, l.Hashes), nil
}

func (l *Lookup) getAbilityDamages(battleInteractions []BattleInteraction) ([]AbilityDamage, error) {
	damages := []AbilityDamage{}

	for j := range battleInteractions {
		bi := &battleInteractions[j]

		if bi.Damage == nil {
			continue
		}

		damages = append(damages, bi.Damage.DamageCalc...)
	}

	return damages, nil
}