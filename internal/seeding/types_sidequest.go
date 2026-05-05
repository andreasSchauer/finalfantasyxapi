package seeding

import (
	"fmt"
)

type Sidequest struct {
	ID int32
	Quest
	Subquests []Subquest `json:"subquests"`
}

func (s Sidequest) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Quest.ID,
	}
}

func (s Sidequest) GetID() int32 {
	return s.ID
}

func (s Sidequest) ToKeyFields() []any {
	return []any{
		s.Name,
	}
}

func (s Sidequest) Error() string {
	return fmt.Sprintf("sidequest %s", s.Name)
}

func (s Sidequest) GetResParamsQuest() ResParamsQuest {
	return ResParamsQuest{
		ID:        s.ID,
		Sidequest: &s.Name,
		Subquest:  nil,
		Type:      string(s.Quest.Type),
	}
}

type Subquest struct {
	ID int32
	Quest
	SidequestID int32
}

func (s Subquest) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Quest.ID,
		s.SidequestID,
	}
}

func (s Subquest) ToKeyFields() []any {
	return []any{
		s.Name,
	}
}

func (s Subquest) GetID() int32 {
	return s.ID
}

func (s *Subquest) SetID(id int32) {
	s.ID = id
}

func (s Subquest) Error() string {
	return fmt.Sprintf("subquest %s", s.Name)
}

func (s Subquest) GetResParamsQuest() ResParamsQuest {
	return ResParamsQuest{
		ID:        s.ID,
		Sidequest: nil,
		Subquest:  &s.Name,
		Type:      string(s.Quest.Type),
	}
}
