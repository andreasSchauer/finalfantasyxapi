package seeding

import "fmt"

type ElementalResist struct {
	ID        int32	 `json:"elemental_resist_id"`
	ElementID int32  `json:"element_id"`
	Element   string `json:"name"`
	Affinity  string `json:"affinity"`
}

func (er ElementalResist) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", er),
		er.ElementID,
		er.Affinity,
	}
}

func (er ElementalResist) ToKeyFields() []any {
	return []any{
		er.ElementID,
		er.Affinity,
	}
}

func (er ElementalResist) GetID() int32 {
	return er.ID
}

func (er *ElementalResist) SetID(id int32) {
	er.ID = id
}

func (er ElementalResist) Error() string {
	return fmt.Sprintf("elemental resist with element: %s, affinity: %s", er.Element, er.Affinity)
}
