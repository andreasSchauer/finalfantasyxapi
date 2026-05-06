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
		StatID:         make([]int32, len(damages)),
		DamageType:     make([]database.DamageType, len(damages)),
		DamageFormula:  make([]database.DamageFormula, len(damages)),
		DamageConstant: make([]int32, len(damages)),
	}

	for i, d := range damages {
		params.DataHash[i] = generateDataHash(d)
		params.Condition[i] = h.GetNullString(d.Condition)
		params.AttackType[i] = database.AttackType(d.AttackType)
		params.StatID[i] = d.StatID
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

		newDamages, err := l.prepareAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}
		damages = append(damages, newDamages...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		newDamages, err := l.prepareAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		newDamages, err := l.prepareAbilityDamages(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		newDamages, err := l.prepareAbilityDamages(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.miscAbilities {
		ability := &l.json.miscAbilities[i]

		newDamages, err := l.prepareAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		newDamages, err := l.prepareAbilityDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	return dedupeRows(damages, l.Hashes), nil
}

func (l *Lookup) prepareAbilityDamages(battleInteractions []BattleInteraction) ([]AbilityDamage, error) {
	damages := []AbilityDamage{}

	for j := range battleInteractions {
		bi := &battleInteractions[j]

		if bi.Damage == nil {
			continue
		}

		newDamages, err := l.prepareAbilityDamage(bi.Damage.DamageCalc)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	return damages, nil
}

func (l *Lookup) prepareAbilityDamage(abilityDamages []AbilityDamage) ([]AbilityDamage, error) {
	damages := []AbilityDamage{}
	var err error

	for i := range abilityDamages {
		ad := &abilityDamages[i]

		ad.StatID, err = assignFK(ad.TargetStat, l.Stats)
		if err != nil {
			return nil, err
		}

		damages = append(damages, *ad)
	}

	return damages, nil
}
