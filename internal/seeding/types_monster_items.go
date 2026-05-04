package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterItems struct {
	ID                  int32
	MonsterID           int32
	DropChance          int32          `json:"drop_chance"`
	DropCondition       *string        `json:"drop_condition"`
	OtherItemsCondition *string        `json:"other_items_condition"`
	OtherItems          []PossibleItem `json:"other_items"`
	StealCommon         *ItemAmount    `json:"steal_common"`
	StealRare           *ItemAmount    `json:"steal_rare"`
	DropCommon          *ItemAmount    `json:"drop_common"`
	DropRare            *ItemAmount    `json:"drop_rare"`
	SecondaryDropCommon *ItemAmount    `json:"secondary_drop_common"`
	SecondaryDropRare   *ItemAmount    `json:"secondary_drop_rare"`
	Bribe               *ItemAmount    `json:"bribe"`
}

func (m MonsterItems) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.MonsterID,
		m.DropChance,
		h.DerefOrNil(m.DropCondition),
		h.DerefOrNil(m.OtherItemsCondition),
		h.ObjPtrToID(m.StealCommon),
		h.ObjPtrToID(m.StealRare),
		h.ObjPtrToID(m.DropCommon),
		h.ObjPtrToID(m.DropRare),
		h.ObjPtrToID(m.SecondaryDropCommon),
		h.ObjPtrToID(m.SecondaryDropRare),
		h.ObjPtrToID(m.Bribe),
	}
}

func (m MonsterItems) GetID() int32 {
	return m.ID
}

func (m *MonsterItems) SetID(id int32) {
	m.ID = id
}

func (m MonsterItems) Error() string {
	return fmt.Sprintf("monster items of monster with id %d", m.MonsterID)
}
