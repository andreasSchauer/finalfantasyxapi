package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AlteredState struct {
	ID          int32
	MonsterID   int32
	Condition   string           `json:"condition"`
	IsTemporary bool             `json:"is_temporary"`
	Changes     []AltStateChange `json:"changes"`
}

func (a AlteredState) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.MonsterID,
		a.Condition,
		a.IsTemporary,
	}
}

func (a AlteredState) GetID() int32 {
	return a.ID
}

func (a *AlteredState) SetID(id int32) {
	a.ID = id
}

func (a AlteredState) Error() string {
	return fmt.Sprintf("altered state with monster id: %d, is temporary: %t, condition: %s", a.MonsterID, a.IsTemporary, a.Condition)
}

type AltStateChange struct {
	ID               int32
	AlteredStateID   int32
	AlterationType   string            `json:"alteration_type"`
	Distance         *int32            `json:"distance"`
	Properties       []string          `json:"properties"`
	AutoAbilities    []string          `json:"auto_abilities"`
	BaseStats        []BaseStat        `json:"base_stats"`
	ElemResists      []ElementalResist `json:"elem_resists"`
	StatusImmunities []string          `json:"status_immunities"`
	AddedStatus      *InflictedStatus  `json:"added_status"`
}

func (a AltStateChange) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.AlteredStateID,
		a.AlterationType,
		h.DerefOrNil(a.Distance),
		h.ObjPtrToID(a.AddedStatus),
	}
}

func (a AltStateChange) GetID() int32 {
	return a.ID
}

func (a *AltStateChange) SetID(id int32) {
	a.ID = id
}

func (a AltStateChange) Error() string {
	return fmt.Sprintf("alt stat change with altered state id: %d, alteration type: %s", a.AlteredStateID, a.AlterationType)
}
