package seeding

import "fmt"

type ModifierChange struct {
	ID              int32
	ModifierID      int32
	ModifierName    string  `json:"name"`
	CalculationType string  `json:"calculation_type"`
	Value           float32 `json:"value"`
}

func (m ModifierChange) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.ModifierID,
		m.CalculationType,
		m.Value,
	}
}

func (m ModifierChange) GetID() int32 {
	return m.ID
}

func (m *ModifierChange) SetID(id int32) {
	m.ID = id
}

func (m ModifierChange) Error() string {
	return fmt.Sprintf("modifier change with modifier: %s, calc type: %s, value %f", m.ModifierName, m.CalculationType, m.Value)
}
