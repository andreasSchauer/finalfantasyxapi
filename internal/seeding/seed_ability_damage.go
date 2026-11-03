package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop1SeedClassDamages(qtx *database.Queries, ctx context.Context) error {
	damages, err := extractClassDamages()
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
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractClassDamages() ([]seedtypes.ClassDamage, error) {
	damages := []seedtypes.ClassDamage{}

	for i := range rawJSON.playerAbilities {
		ability := &rawJSON.playerAbilities[i]

		newDamages, err := getClassDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}
		damages = append(damages, newDamages...)
	}

	for i := range rawJSON.overdriveAbilities {
		ability := &rawJSON.overdriveAbilities[i]

		newDamages, err := getClassDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range rawJSON.items {
		item := &rawJSON.items[i]

		newDamages, err := getClassDamages(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range rawJSON.triggerCommands {
		command := &rawJSON.triggerCommands[i]

		newDamages, err := getClassDamages(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range rawJSON.miscAbilities {
		ability := &rawJSON.miscAbilities[i]

		newDamages, err := getClassDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range rawJSON.enemyAbilities {
		ability := &rawJSON.enemyAbilities[i]

		newDamages, err := getClassDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	return dedupeRows(damages, hashes), nil
}

func getClassDamages(battleInteractions []seedtypes.BattleInteraction) ([]seedtypes.ClassDamage, error) {
	damages := []seedtypes.ClassDamage{}

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
