package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Mix struct {
	ID                   int32
	OverdriveID          int32
	Name                 string           `json:"name"`
	Category             string           `json:"category"`
	BestCombinations     []MixCombination `json:"best_combinations"`
	PossibleCombinations []MixCombination `json:"possible_combinations"`
}

func (m Mix) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.OverdriveID,
		m.Category,
	}
}

func (m Mix) ToKeyFields() []any {
	return []any{
		m.Name,
	}
}

func (m Mix) GetID() int32 {
	return m.ID
}

func (m Mix) Error() string {
	return fmt.Sprintf("mix %s", m.Name)
}

func (m Mix) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   m.ID,
		Name: m.Name,
	}
}

type MixCombination struct {
	ID           int32
	MixID        int32
	FirstItem    string `json:"first_item"`
	SecondItem   string `json:"second_item"`
	FirstItemID  int32
	SecondItemID int32
	IsBestCombo  bool
}

func (m MixCombination) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.MixID,
		m.FirstItemID,
		m.SecondItemID,
		m.IsBestCombo,
	}
}

func (m MixCombination) ToKeyFields() []any {
	return []any{
		m.FirstItem,
		m.SecondItem,
	}
}

func (m MixCombination) GetID() int32 {
	return m.ID
}

func (m *MixCombination) SetID(id int32) {
	m.ID = id
}

func (m MixCombination) Error() string {
	return fmt.Sprintf("mix combination with first item: %s, second item: %s", m.FirstItem, m.SecondItem)
}
