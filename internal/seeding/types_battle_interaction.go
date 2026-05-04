package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type BattleInteraction struct {
	ID                        int32
	Target                    string            `json:"target"`
	BasedOnUserAttack         bool              `json:"based_on_user_attack"`
	Range                     *int32            `json:"range"`
	Damage                    *Damage           `json:"damage"`
	ShatterRate               int32             `json:"shatter_rate"`
	Accuracy                  Accuracy          `json:"accuracy"`
	AffectedBy                []string          `json:"affected_by"`
	HitAmount                 int32             `json:"hit_amount"`
	SpecialAction             *string           `json:"special_action"`
	InflictedDelay            *InflictedDelay   `json:"inflicted_delay"`
	InflictedStatusConditions []InflictedStatus `json:"inflicted_status_conditions"`
	RemovedStatusConditions   []string          `json:"removed_status_conditions"`
	CopiedStatusConditions    []InflictedStatus `json:"copied_status_conditions"`
	StatChanges               []StatChange      `json:"stat_changes"`
	ModifierChanges           []ModifierChange  `json:"modifier_changes"`
}

func (bi BattleInteraction) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", bi),
		bi.Target,
		bi.BasedOnUserAttack,
		h.DerefOrNil(bi.Range),
		bi.ShatterRate,
		bi.Accuracy.ID,
		bi.HitAmount,
		h.DerefOrNil(bi.SpecialAction),
		h.ObjPtrToID(bi.InflictedDelay),
	}
}

func (bi BattleInteraction) GetID() int32 {
	return bi.ID
}

func (bi *BattleInteraction) SetID(id int32) {
	bi.ID = id
}

func (bi BattleInteraction) Error() string {
	return fmt.Sprintf("battle interaction with target: %s, phys attack: %t, range: %v, damage id: %v, shatter rate: %v, accuracy id: %d, hit amount: %d, special action: %v", bi.Target, bi.BasedOnUserAttack, h.PtrToString(bi.Range), h.ObjPtrToID(bi.Damage), bi.ShatterRate, bi.Accuracy.ID, bi.HitAmount, h.PtrToString(bi.SpecialAction))
}
