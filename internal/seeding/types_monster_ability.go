package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterAbility struct {
	ID        int32
	AbilityID int32
	AbilityReference
	IsForced bool `json:"is_forced"`
	IsUnused bool `json:"is_unused"`
}

func (m MonsterAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.AbilityID,
		m.IsForced,
		m.IsUnused,
	}
}

func (m MonsterAbility) GetID() int32 {
	return m.ID
}

func (m *MonsterAbility) SetID(id int32) {
	m.ID = id
}

func (m MonsterAbility) Error() string {
	return fmt.Sprintf("monster ability '%s', type: %s, is forced: %t, is unused: %t", h.NameToString(m.Name, m.Version, nil), m.AbilityType, m.IsForced, m.IsUnused)
}
