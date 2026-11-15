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

func Loop1SeedInflictedDelays(qtx *database.Queries, ctx context.Context) error {
	delays := extractInflictedDelays()

	params := database.CreateInflictedDelayBulkParams{
		DataHash:      make([]string, len(delays)),
		Condition:     make([]sql.NullString, len(delays)),
		DelayStrength: make([]database.DelayStrength, len(delays)),
	}

	for i, d := range delays {
		params.DataHash[i] = seedhelpers.GenerateDataHash(d)
		params.Condition[i] = h.GetNullString(d.Condition)
		params.DelayStrength[i] = database.DelayStrength(d.DelayStrength)
	}

	dbRows, err := qtx.CreateInflictedDelayBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create inflicted delays: %v", err)
	}

	for _, row := range dbRows {
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractInflictedDelays() []seedtypes.InflictedDelay {
	delays := []seedtypes.InflictedDelay{}

	for _, ability := range seedhelpers.RawJSON.EnemyAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range seedhelpers.RawJSON.Items {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range seedhelpers.RawJSON.OverdriveAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range seedhelpers.RawJSON.PlayerAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range seedhelpers.RawJSON.TriggerCommands {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range seedhelpers.RawJSON.MiscAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	return dedupeRows(delays, seedhelpers.Hashes)
}
