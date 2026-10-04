package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop2SeedDamages(qtx *database.Queries, ctx context.Context) error {
	damages, err := l.extractDamages()
	if err != nil {
		return err
	}

	params := database.CreateDamageBulkParams{
		DataHash:        make([]string, len(damages)),
		AttackType: 	 make([]database.AttackType, len(damages)),
		DamageType: 	 make([]database.DamageType, len(damages)),
		HpClassID: 		 make([]sql.NullInt32, len(damages)),
		MpClassID: 		 make([]sql.NullInt32, len(damages)),
		CtbClassID: 	 make([]sql.NullInt32, len(damages)),
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractDamages() ([]Damage, error) {
	damages := []Damage{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		newDamages, err := l.prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		newDamages, err := l.prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		newDamages, err := l.prepareDamages(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		newDamages, err := l.prepareDamages(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.miscAbilities {
		ability := &l.json.miscAbilities[i]

		newDamages, err := l.prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		newDamages, err := l.prepareDamages(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		damages = append(damages, newDamages...)
	}

	return dedupeRows(damages, l.Hashes), nil
}

func (l *Lookup) prepareDamages(battleInteractions []BattleInteraction) ([]Damage, error) {
	damages := []Damage{}
	var err error

	for i := range battleInteractions {
		bi := &battleInteractions[i]

		if bi.Damage != nil {
			if bi.Damage.HP != nil {
				bi.Damage.HP.ID, err = l.GetHashID(bi.Damage.HP)
				if err != nil {
					return nil, err
				}
			}

			if bi.Damage.MP != nil {
				bi.Damage.MP.ID, err = l.GetHashID(bi.Damage.MP)
				if err != nil {
					return nil, err
				}
			}
			if bi.Damage.CTB != nil {
				bi.Damage.CTB.ID, err = l.GetHashID(bi.Damage.CTB)
				if err != nil {
					return nil, err
				}
			}

			bi.Damage.ElementID, err = assignFKPtr(bi.Damage.Element, l.Elements)
			if err != nil {
				return nil, err
			}

			damages = append(damages, *bi.Damage)
		}
	}

	return damages, nil
}

func (l *Lookup) completeDamage(damage *Damage) error {
	if damage == nil {
		return nil
	}

	err := l.assignID(damage)
	if err != nil {
		return err
	}

	return nil
}