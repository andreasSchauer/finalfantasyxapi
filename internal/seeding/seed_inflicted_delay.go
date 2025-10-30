package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type InflictedDelay struct {
	ID					int32
	Condition			*string		`json:"condition"`
	CTBAttackType		string		`json:"ctb_attack_type"`
	DelayType			string		`json:"delay_type"`
	DamageConstant		int32		`json:"damage_constant"`
}

func (id InflictedDelay) ToHashFields() []any {
	return []any{
		derefOrNil(id.Condition),
		id.CTBAttackType,
		id.DelayType,
		id.DamageConstant,
	}
}


func (id InflictedDelay) GetID() int32 {
	return id.ID
}


func (l *lookup) seedInflictedDelay(qtx *database.Queries, delay InflictedDelay) (InflictedDelay, error) {
	dbDelay, err := qtx.CreateInflictedDelay(context.Background(), database.CreateInflictedDelayParams{
		DataHash: 		generateDataHash(delay),
		Condition: 		getNullString(delay.Condition),
		CtbAttackType: 	database.CtbAttackType(delay.CTBAttackType),
		DelayType: 		database.DelayType(delay.DelayType),
		DamageConstant: delay.DamageConstant,
	})
	if err != nil {
		return InflictedDelay{}, fmt.Errorf("couldn't create inflicted delay: %v", err)
	}

	delay.ID = dbDelay.ID

	return delay, nil
}