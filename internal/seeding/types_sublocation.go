package seeding

import (
	"fmt"
)

type Sublocation struct {
	ID   int32
	Name string `json:"sublocation"`

	Areas    []Area `json:"areas"`
	Location Location
}

func (s Sublocation) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Location.ID,
		s.Name,
	}
}

func (s Sublocation) ToKeyFields() []any {
	return []any{
		s.Name,
	}
}

func (s Sublocation) GetID() int32 {
	return s.ID
}

func (s *Sublocation) SetID(id int32) {
	s.ID = id
}

func (s Sublocation) Error() string {
	return fmt.Sprintf("sublocation %s", s.Name)
}

func (s Sublocation) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
		ID:   s.ID,
		Name: s.Name,
	}
}
