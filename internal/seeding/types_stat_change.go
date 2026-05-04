package seeding

import "fmt"

type StatChange struct {
	ID              int32
	StatID          int32
	StatName        string  `json:"name"`
	CalculationType string  `json:"calculation_type"`
	Value           float32 `json:"value"`
}

func (s StatChange) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.StatID,
		s.CalculationType,
		s.Value,
	}
}

func (s StatChange) GetID() int32 {
	return s.ID
}

func (s *StatChange) SetID(id int32) {
	s.ID = id
}

func (s StatChange) Error() string {
	return fmt.Sprintf("stat change with stat: %s, calc type: %s, value %f", s.StatName, s.CalculationType, s.Value)
}
