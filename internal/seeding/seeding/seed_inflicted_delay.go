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

func Loop1SeedInflictedDelays(qtx *database.Queries, ctx context.Context) error {
	delays := extractInflictedDelays()

	params := database.CreateInflictedDelayBulkParams{
		DataHash:      make([]string, len(delays)),
		Condition:     make([]sql.NullString, len(delays)),
		DelayStrength: make([]database.DelayStrength, len(delays)),
	}

	for i, d := range delays {
		params.DataHash[i] = sh.GenerateDataHash(d)
		params.Condition[i] = h.GetNullString(d.Condition)
		params.DelayStrength[i] = database.DelayStrength(d.DelayStrength)
	}

	dbRows, err := qtx.CreateInflictedDelayBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create inflicted delays: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractInflictedDelays() []seedtypes.InflictedDelay {
	delays := []seedtypes.InflictedDelay{}

	for _, ability := range sh.RawJSON.EnemyAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range sh.RawJSON.Items {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range sh.RawJSON.OverdriveAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range sh.RawJSON.PlayerAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range sh.RawJSON.TriggerCommands {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range sh.RawJSON.MiscAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	return dedupeRows(delays, sh.Hashes)
}
