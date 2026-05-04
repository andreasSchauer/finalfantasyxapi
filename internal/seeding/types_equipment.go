package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EquipmentTable struct {
	ID                      int32
	Type                    string `json:"type"`
	Classification          string `json:"classification"`
	SpecificCharacterID     *int32
	SpecificCharacter       *string         `json:"specific_character"`
	Version                 *int32          `json:"version"`
	Priority                *int32          `json:"priority"`
	RequiredAutoAbilities   []string        `json:"required_auto_abilities"`
	SelectableAutoAbilities []AbilityPool   `json:"selectable_auto_abilities"`
	RequiredSlots           *int32          `json:"required_slots"`
	EquipmentNames          []EquipmentName `json:"names"`
}

func (e EquipmentTable) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", e),
		e.Type,
		e.Classification,
		h.DerefOrNil(e.SpecificCharacterID),
		h.DerefOrNil(e.Version),
		h.DerefOrNil(e.Priority),
		h.DerefOrNil(e.RequiredSlots),
	}
}

func (e EquipmentTable) ToKeyFields() []any {
	return []any{
		e.Type,
		e.Classification,
		h.DerefOrNil(e.SpecificCharacter),
		h.DerefOrNil(e.Version),
		h.DerefOrNil(e.Priority),
	}
}

func (e EquipmentTable) GetID() int32 {
	return e.ID
}

func (e EquipmentTable) Error() string {
	return fmt.Sprintf("equipment table with type: %s, classification: %s, specific character: %v, version: %v, priority: %v", e.Type, e.Classification, h.PtrToString(e.SpecificCharacter), h.PtrToString(e.Version), h.PtrToString(e.Priority))
}

func (e EquipmentTable) GetResParamsUnnamed() h.ResParamsUnnamed {
	return h.ResParamsUnnamed{
		ID: e.ID,
	}
}

type AbilityPool struct {
	ID               int32
	EquipmentTableID int32
	PoolIdx          int32
	AutoAbilities    []string `json:"auto_abilities"`
	ReqAmount        int32    `json:"req_amount"`
}

func (p AbilityPool) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", p),
		p.EquipmentTableID,
		p.PoolIdx,
		p.ReqAmount,
	}
}

func (p AbilityPool) GetID() int32 {
	return p.ID
}

func (p *AbilityPool) SetID(id int32) {
	p.ID = id
}

func (p AbilityPool) Error() string {
	return fmt.Sprintf("ability pool with equipment table id: %d, req amount: %d", p.EquipmentTableID, p.ReqAmount)
}

type EquipmentName struct {
	ID            int32
	CharacterID   int32
	CharacterName string `json:"character"`
	Name          string `json:"name"`
}

func (e EquipmentName) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", e),
		e.CharacterID,
		e.Name,
	}
}

func (e EquipmentName) GetID() int32 {
	return e.ID
}

func (e *EquipmentName) SetID(id int32) {
	e.ID = id
}

func (e EquipmentName) Error() string {
	return fmt.Sprintf("equipment name %s, character name: %s", e.Name, e.CharacterName)
}

func (e EquipmentName) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   e.ID,
		Name: e.Name,
	}
}
