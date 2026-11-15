package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3SeedBattleInteractions(qtx *database.Queries, ctx context.Context) error {
	bis, err := extractBattleInteractions()
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
		params.DataHash[i] = sh.GenerateDataHash(bi)
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
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractBattleInteractions() ([]seedtypes.BattleInteraction, error) {
	bis := []seedtypes.BattleInteraction{}

	for i := range sh.RawJSON.PlayerAbilities {
		ability := &sh.RawJSON.PlayerAbilities[i]

		bisNew, err := prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range sh.RawJSON.OverdriveAbilities {
		ability := &sh.RawJSON.OverdriveAbilities[i]

		bisNew, err := prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range sh.RawJSON.Items {
		item := &sh.RawJSON.Items[i]

		bisNew, err := prepareBattleInteractions(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range sh.RawJSON.TriggerCommands {
		command := &sh.RawJSON.TriggerCommands[i]

		bisNew, err := prepareBattleInteractions(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range sh.RawJSON.MiscAbilities {
		ability := &sh.RawJSON.MiscAbilities[i]

		bisNew, err := prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range sh.RawJSON.EnemyAbilities {
		ability := &sh.RawJSON.EnemyAbilities[i]

		bisNew, err := prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	return dedupeRows(bis, sh.Hashes), nil
}

func prepareBattleInteractions(bis []seedtypes.BattleInteraction) ([]seedtypes.BattleInteraction, error) {
	bisNew := []seedtypes.BattleInteraction{}
	var err error

	for j := range bis {
		bi := &bis[j]

		bi.Accuracy.ID, err = sh.GetHashID(bi.Accuracy)
		if err != nil {
			return nil, err
		}

		if bi.InflictedDelay != nil {
			bi.InflictedDelay.ID, err = sh.GetHashID(bi.InflictedDelay)
			if err != nil {
				return nil, err
			}
		}

		bisNew = append(bisNew, *bi)
	}

	return bisNew, nil
}

func completeBattleInteractions(bis []seedtypes.BattleInteraction) error {
	for i := range bis {
		bi := &bis[i]

		err := assignID(bi)
		if err != nil {
			return err
		}

		err = completeDamage(bi.Damage)
		if err != nil {
			return err
		}

		err = assignIDs(bi.InflictedStatusConditions)
		if err != nil {
			return err
		}

		err = assignIDs(bi.CopiedStatusConditions)
		if err != nil {
			return err
		}

		err = assignIDs(bi.StatChanges)
		if err != nil {
			return err
		}

		err = assignIDs(bi.ModifierChanges)
		if err != nil {
			return err
		}
	}

	return nil
}
