package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type InflictedDelay struct {
	ID             int32
	Condition      *string `json:"condition"`
	CTBAttackType  string  `json:"ctb_attack_type"`
	DelayType      string  `json:"delay_type"`
	DamageConstant int32   `json:"damage_constant"`
}

func (id InflictedDelay) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", id),
		h.DerefOrNil(id.Condition),
		id.CTBAttackType,
		id.DelayType,
		id.DamageConstant,
	}
}

func (id InflictedDelay) GetID() int32 {
	return id.ID
}

func (id InflictedDelay) Error() string {
	return fmt.Sprintf("inflicted delay with ctb attack type: %s, delay type: %s, constant: %d, condition: %v", id.CTBAttackType, id.DelayType, id.DamageConstant, h.PtrToString(id.Condition))
}

func (l *Lookup) seedInflictedDelay(qtx *database.Queries, delay InflictedDelay) (InflictedDelay, error) {
	dbDelay, err := qtx.CreateInflictedDelay(context.Background(), database.CreateInflictedDelayParams{
		DataHash:       generateDataHash(delay),
		Condition:      h.GetNullString(delay.Condition),
		CtbAttackType:  database.CtbAttackType(delay.CTBAttackType),
		DelayType:      database.DelayType(delay.DelayType),
		DamageConstant: delay.DamageConstant,
	})
	if err != nil {
		return InflictedDelay{}, h.NewErr(delay.Error(), err, "couldn't create inflicted delay")
	}

	delay.ID = dbDelay.ID

	return delay, nil
}


func (l *Lookup) loop1SeedInflictedDelays(qtx *database.Queries, ctx context.Context) error {
	delays := l.extractInflictedDelays()

	params := database.CreateInflictedDelayBulkParams{
		DataHash:       make([]string, len(delays)),
		Condition:      make([]sql.NullString, len(delays)),
		CtbAttackType:  make([]database.CtbAttackType, len(delays)),
		DelayType:      make([]database.DelayType, len(delays)),
		DamageConstant: make([]int32, len(delays)),
	}

	for i, d := range delays {
		params.DataHash[i] = generateDataHash(d)
		params.Condition[i] = h.GetNullString(d.Condition)
		params.CtbAttackType[i] = database.CtbAttackType(d.CTBAttackType)
		params.DelayType[i] = database.DelayType(d.DelayType)
		params.DamageConstant[i] = d.DamageConstant
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

	for _, ability := range l.json.unspecifiedAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, status := range l.json.statusConditions {
		if status.CtbOnInfliction != nil {
			delays = append(delays, *status.CtbOnInfliction)
		}
	}

	return dedupeRows(delays, l.Hashes)
}