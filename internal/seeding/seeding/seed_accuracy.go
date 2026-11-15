package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1SeedAccuracies(qtx *database.Queries, ctx context.Context) error {
	accuracies := extractAccuracies()

	params := database.CreateAbilityAccuracyBulkParams{
		DataHash:    make([]string, len(accuracies)),
		AccSource:   make([]database.AccSourceType, len(accuracies)),
		HitChance:   make([]sql.NullInt32, len(accuracies)),
		AccModifier: make([]sql.NullFloat64, len(accuracies)),
	}

	for i, a := range accuracies {
		params.DataHash[i] = seedhelpers.GenerateDataHash(a)
		params.AccSource[i] = database.AccSourceType(a.AccSource)
		params.HitChance[i] = h.GetNullInt32(a.HitChance)
		params.AccModifier[i] = h.GetNullFloat64(a.AccModifier)
	}

	dbRows, err := qtx.CreateAbilityAccuracyBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability accuracies: %v", err)
	}

	for _, row := range dbRows {
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractAccuracies() []seedtypes.Accuracy {
	accuracies := []seedtypes.Accuracy{}

	for _, ability := range seedhelpers.RawJSON.EnemyAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range seedhelpers.RawJSON.Items {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range seedhelpers.RawJSON.OverdriveAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range seedhelpers.RawJSON.PlayerAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range seedhelpers.RawJSON.TriggerCommands {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range seedhelpers.RawJSON.MiscAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, aeon := range seedhelpers.RawJSON.Aeons {
		if aeon.PhysAtkAccuracy != nil {
			accuracies = append(accuracies, *aeon.PhysAtkAccuracy)
		}
	}

	return dedupeRows(accuracies, seedhelpers.Hashes)
}
