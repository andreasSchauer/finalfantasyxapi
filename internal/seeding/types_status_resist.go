package seeding

import "fmt"

type StatusResist struct {
	ID                int32
	StatusConditionID int32
	StatusCondition   string `json:"name"`
	Resistance        int32  `json:"resistance"`
}

func (sr StatusResist) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", sr),
		sr.StatusConditionID,
		sr.Resistance,
	}
}

func (sr StatusResist) GetID() int32 {
	return sr.ID
}

func (sr *StatusResist) SetID(id int32) {
	sr.ID = id
}

func (sr StatusResist) GetName() string {
	return sr.StatusCondition
}

func (sr StatusResist) GetVersion() *int32 {
	return nil
}

func (sr StatusResist) GetVal() int32 {
	return sr.Resistance
}

func (sr StatusResist) Error() string {
	return fmt.Sprintf("status resist with status %s, resistance %d", sr.StatusCondition, sr.Resistance)
}
