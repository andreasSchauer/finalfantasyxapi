package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedBattleInteractions(qtx *database.Queries, ctx context.Context) error {
	bis, err := l.extractBattleInteractions()
	if err != nil {
		return err
	}

	params := database.CreateBattleInteractionBulkParams{
		DataHash:          make([]string, len(bis)),
		Target:            make([]database.TargetType, len(bis)),
		BasedOnUserAttack: make([]bool, len(bis)),
		Range:             make([]sql.NullInt32, len(bis)),
		ShatterRate:       make([]int32, len(bis)),
		AccuracyID:        make([]int32, len(bis)),
		InflictedDelayID:  make([]sql.NullInt32, len(bis)),
		HitAmount:         make([]int32, len(bis)),
		SpecialAction:     make([]database.NullSpecialActionType, len(bis)),
	}

	for i, bi := range bis {
		params.DataHash[i] = generateDataHash(bi)
		params.Target[i] = database.TargetType(bi.Target)
		params.BasedOnUserAttack[i] = bi.BasedOnUserAttack
		params.Range[i] = h.GetNullInt32(bi.Range)
		params.ShatterRate[i] = bi.ShatterRate
		params.AccuracyID[i] = bi.Accuracy.ID
		params.InflictedDelayID[i] = h.ObjPtrToNullInt32ID(bi.InflictedDelay)
		params.HitAmount[i] = bi.HitAmount
		params.SpecialAction[i] = database.ToNullSpecialActionType(bi.SpecialAction)
	}

	dbRows, err := qtx.CreateBattleInteractionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create battle interactions: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractBattleInteractions() ([]BattleInteraction, error) {
	bis := []BattleInteraction{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		bisNew, err := l.prepareBattleInteractions(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		bisNew, err := l.prepareBattleInteractions(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.miscAbilities {
		ability := &l.json.miscAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	return dedupeRows(bis, l.Hashes), nil
}

func (l *Lookup) prepareBattleInteractions(bis []BattleInteraction) ([]BattleInteraction, error) {
	bisNew := []BattleInteraction{}
	var err error

	for j := range bis {
		bi := &bis[j]

		bi.Accuracy.ID, err = l.GetHashID(bi.Accuracy)
		if err != nil {
			return nil, err
		}

		if bi.InflictedDelay != nil {
			bi.InflictedDelay.ID, err = l.GetHashID(bi.InflictedDelay)
		}

		bisNew = append(bisNew, *bi)
	}

	return bisNew, nil
}

func (l *Lookup) completeBattleInteractions(bis []BattleInteraction) error {
	for i := range bis {
		bi := &bis[i]

		err := l.assignID(bi)
		if err != nil {
			return err
		}

		err = l.completeDamage(bi.Damage)
		if err != nil {
			return err
		}

		err = assignIDs(l, bi.InflictedStatusConditions)
		if err != nil {
			return err
		}

		err = assignIDs(l, bi.CopiedStatusConditions)
		if err != nil {
			return err
		}

		err = assignIDs(l, bi.StatChanges)
		if err != nil {
			return err
		}

		err = assignIDs(l, bi.ModifierChanges)
		if err != nil {
			return err
		}
	}

	return nil
}
