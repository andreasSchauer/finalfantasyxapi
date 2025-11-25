package seeding

import (
	"context"
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
	return fmt.Sprintf("inflicted delay with ctb attack type: %s, delay type: %s, constant: %d, condition: %v", id.CTBAttackType, id.DelayType, id.DamageConstant, h.DerefOrNil(id.Condition))
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
		return InflictedDelay{}, h.GetErr(delay.Error(), err, "couldn't create inflicted delay")
	}

	delay.ID = dbDelay.ID

	return delay, nil
}
