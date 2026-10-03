package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop1SeedInflictedDelays(qtx *database.Queries, ctx context.Context) error {
	delays := l.extractInflictedDelays()

	params := database.CreateInflictedDelayBulkParams{
		DataHash:       make([]string, len(delays)),
		Condition:      make([]sql.NullString, len(delays)),
		DelayStrength:  make([]database.DelayStrength, len(delays)),
	}

	for i, d := range delays {
		params.DataHash[i] = generateDataHash(d)
		params.Condition[i] = h.GetNullString(d.Condition)
		params.DelayStrength[i] = database.DelayStrength(d.DelayStrength)
	}

	dbRows, err := qtx.CreateInflictedDelayBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create inflicted delays: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractInflictedDelays() []InflictedDelay {
	delays := []InflictedDelay{}

	for _, ability := range l.json.enemyAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.items {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.overdriveAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.playerAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.triggerCommands {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.miscAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	return dedupeRows(delays, l.Hashes)
}
