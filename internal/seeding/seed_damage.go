package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop2SeedDamages(qtx *database.Queries, ctx context.Context) error {
	damages, err := extractDamages()
	if err != nil {
		return err
	}

	params := database.CreateDamageBulkParams{
		DataHash:        make([]string, len(damages)),
		AttackType:      make([]database.AttackType, len(damages)),
		DamageType:      make([]database.DamageType, len(damages)),
		HpClassID:       make([]sql.NullInt32, len(damages)),
		MpClassID:       make([]sql.NullInt32, len(damages)),
		CtbClassID:      make([]sql.NullInt32, len(damages)),
		Critical:        make([]database.NullCriticalType, len(damages)),
		CriticalPlusVal: make([]sql.NullInt32, len(damages)),
		IsPiercing:      make([]bool, len(damages)),
		BreakDmgLimit:   make([]database.NullBreakDmgLmtType, len(damages)),
		ElementID:       make([]sql.NullInt32, len(damages)),
	}

	for i, d := range damages {
		params.DataHash[i] = generateDataHash(d)
		params.AttackType[i] = database.AttackType(d.AttackType)
		params.DamageType[i] = database.DamageType(d.DamageType)
		params.HpClassID[i] = h.ObjPtrToNullInt32ID(d.HP)
		params.MpClassID[i] = h.ObjPtrToNullInt32ID(d.MP)
		params.CtbClassID[i] = h.ObjPtrToNullInt32ID(d.CTB)
		params.Critical[i] = database.ToNullCriticalType(d.Critical)
		params.CriticalPlusVal[i] = h.GetNullInt32(d.CriticalPlusVal)
		params.IsPiercing[i] = d.IsPiercing
		params.BreakDmgLimit[i] = database.ToNullBreakDmgLmtType(d.BreakDmgLimit)
		params.ElementID[i] = h.GetNullInt32(d.ElementID)
	}

	dbRows, err := qtx.CreateDamageBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create damages: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractDamages() ([]seedtypes.Damage, error) {
	damages := []seedtypes.Damage{}

	for i := range rawJSON.playerAbilities {
		ability := &rawJSON.playerAbilities[i]

		newDamages, err := prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range rawJSON.overdriveAbilities {
		ability := &rawJSON.overdriveAbilities[i]

		newDamages, err := prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range rawJSON.items {
		item := &rawJSON.items[i]

		newDamages, err := prepareDamages(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range rawJSON.triggerCommands {
		command := &rawJSON.triggerCommands[i]

		newDamages, err := prepareDamages(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range rawJSON.miscAbilities {
		ability := &rawJSON.miscAbilities[i]

		newDamages, err := prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range rawJSON.enemyAbilities {
		ability := &rawJSON.enemyAbilities[i]

		newDamages, err := prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	return dedupeRows(damages, hashes), nil
}

func prepareDamages(battleInteractions []seedtypes.BattleInteraction) ([]seedtypes.Damage, error) {
	damages := []seedtypes.Damage{}
	var err error

	for i := range battleInteractions {
		bi := &battleInteractions[i]

		if bi.Damage != nil {
			if bi.Damage.HP != nil {
				bi.Damage.HP.ID, err = GetHashID(bi.Damage.HP)
				if err != nil {
					return nil, err
				}
			}

			if bi.Damage.MP != nil {
				bi.Damage.MP.ID, err = GetHashID(bi.Damage.MP)
				if err != nil {
					return nil, err
				}
			}
			if bi.Damage.CTB != nil {
				bi.Damage.CTB.ID, err = GetHashID(bi.Damage.CTB)
				if err != nil {
					return nil, err
				}
			}

			bi.Damage.ElementID, err = assignFKPtr(bi.Damage.Element, lookup.L.Elements)
			if err != nil {
				return nil, err
			}

			damages = append(damages, *bi.Damage)
		}
	}

	return damages, nil
}

func completeDamage(damage *seedtypes.Damage) error {
	if damage == nil {
		return nil
	}

	err := assignID(damage)
	if err != nil {
		return err
	}

	return nil
}
