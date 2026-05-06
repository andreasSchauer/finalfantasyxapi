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
		Critical:        make([]database.NullCriticalType, len(damages)),
		CriticalPlusVal: make([]sql.NullInt32, len(damages)),
		IsPiercing:      make([]bool, len(damages)),
		BreakDmgLimit:   make([]database.NullBreakDmgLmtType, len(damages)),
		ElementID:       make([]sql.NullInt32, len(damages)),
	}

	for i, d := range damages {
		params.DataHash[i] = generateDataHash(d)
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

	err = assignIDs(l, damage.DamageCalc)
	if err != nil {
		return err
	}

	return nil
}

func (l *Lookup) seedJuncDamagesDamageCalc(qtx *database.Queries, ctx context.Context) error {
	const desc string = "damages + damage calc"
	params := database.CreateDamagesDamageCalcJunctionBulkParams{
		DataHash:            make([]string, 0),
		AbilityID:           make([]int32, 0),
		BattleInteractionID: make([]int32, 0),
		DamageID:            make([]int32, 0),
		AbilityDamageID:     make([]int32, 0),
	}

	for _, ability := range l.getAbilities() {
		bis, err := l.getAbilityBattleInteractions(ability)
		if err != nil {
			return err
		}

		for _, bi := range bis {
			if bi.Damage == nil {
				continue
			}

			for _, ad := range bi.Damage.DamageCalc {
				j := FourWayJunction{}
				j.GreatGrandparentID = ability.ID
				j.GrandparentID = bi.ID
				j.ParentID = bi.Damage.ID
				j.ChildID = ad.ID
				dataHash := generateJunctionHash(j, desc)

				params.DataHash = append(params.DataHash, dataHash)
				params.AbilityID = append(params.AbilityID, ability.ID)
				params.BattleInteractionID = append(params.BattleInteractionID, bi.ID)
				params.DamageID = append(params.DamageID, bi.Damage.ID)
				params.AbilityDamageID = append(params.AbilityDamageID, ad.ID)
			}
		}
	}

	return qtx.CreateDamagesDamageCalcJunctionBulk(ctx, params)
}
