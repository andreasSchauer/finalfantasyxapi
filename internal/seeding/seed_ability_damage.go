package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop1SeedClassDamages(qtx *database.Queries, ctx context.Context) error {
	damages, err := l.extractClassDamages()
	if err != nil {
		return err
	}

	params := database.CreateClassDamageBulkParams{
		DataHash:       make([]string, len(damages)),
		Condition:      make([]sql.NullString, len(damages)),
		DamageFormula:  make([]database.DamageFormula, len(damages)),
		DamageConstant: make([]int32, len(damages)),
	}

	for i, d := range damages {
		params.DataHash[i] = generateDataHash(d)
		params.Condition[i] = h.GetNullString(d.Condition)
		params.DamageFormula[i] = database.DamageFormula(d.DamageFormula)
		params.DamageConstant[i] = d.DamageConstant
	}

	dbRows, err := qtx.CreateClassDamageBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability damages: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractClassDamages() ([]ClassDamage, error) {
	damages := []ClassDamage{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		newDamages, err := l.getClassDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}
		damages = append(damages, newDamages...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		newDamages, err := l.getClassDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		newDamages, err := l.getClassDamages(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		newDamages, err := l.getClassDamages(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.miscAbilities {
		ability := &l.json.miscAbilities[i]

		newDamages, err := l.getClassDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		newDamages, err := l.getClassDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	return dedupeRows(damages, l.Hashes), nil
}

func (l *Lookup) getClassDamages(battleInteractions []BattleInteraction) ([]ClassDamage, error) {
	damages := []ClassDamage{}

	for j := range battleInteractions {
		bi := &battleInteractions[j]

		if bi.Damage == nil {
			continue
		}

		if bi.Damage.HP != nil {
			damages = append(damages, *bi.Damage.HP)
		}

		if bi.Damage.MP != nil {
			damages = append(damages, *bi.Damage.MP)
		}

		if bi.Damage.CTB != nil {
			damages = append(damages, *bi.Damage.CTB)
		}
	}

	return damages, nil
}
