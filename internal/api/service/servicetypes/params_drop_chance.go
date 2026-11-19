// Package servicetypes ajsidoaüsdji
package servicetypes

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

type DropChanceParams struct {
	Monster           int32   `json:"monster_id,omitempty"`
	CharacterID       *int32  `json:"character_id,omitempty"`
	PartyMembers      []int32 `json:"party_members,omitempty"`
	AutoAbilities     []int32 `json:"auto_abilities,omitempty"`
	OptionalAbilities bool    `json:"optional_abilities,omitempty"`
	EquipType         *string `json:"equip_type,omitempty"`
	MinEmptySlots     *int32  `json:"min_empty_slots,omitempty"`
	TotalSlots        *int32  `json:"total_slots,omitempty"`
	Decimals          bool    `json:"decimals,omitempty"`
}

func (p DropChanceParams) GetRequiredSlots() int32 {
	return h.Len32(p.AutoAbilities) + h.DerefPtr(p.MinEmptySlots)
}