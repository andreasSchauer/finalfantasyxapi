package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type DroppedAutoAbilities struct {
	ForcedAbility    *DroppedAutoAbilityInfo `json:"forced_ability"`
	DroppedAbilities []DroppedAutoAbility    `json:"dropped_abilities"`
}

func (d DroppedAutoAbilities) Percent() DroppedAutoAbilities {
	d.DroppedAbilities = h.SlicePercent(d.DroppedAbilities)

	return d
}

func (d DroppedAutoAbilities) Round(n int32) DroppedAutoAbilities {
	d.DroppedAbilities = h.SliceRound(d.DroppedAbilities, n)

	return d
}

func getDroppedAutoAbilities(cfg *Config, abilities []seeding.EquipmentDrop) DroppedAutoAbilities {
	if abilities[0].IsForced {
		forcedAbility := convertDroppedAutoAbilityInfo(abilities[0])

		return DroppedAutoAbilities{
			ForcedAbility:    &forcedAbility,
			DroppedAbilities: convertObjSlice(cfg, abilities[1:], convertDroppedAutoAbility),
		}
	}

	return DroppedAutoAbilities{
		DroppedAbilities: convertObjSlice(cfg, abilities, convertDroppedAutoAbility),
	}
}

type DroppedAutoAbilityInfo struct {
	ID         int32    `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

func convertDroppedAutoAbilityInfo(ed seeding.EquipmentDrop) DroppedAutoAbilityInfo {
	return DroppedAutoAbilityInfo{
		ID:         ed.AutoAbilityID,
		Name:       ed.Ability,
		Characters: ed.Characters,
	}
}

type DroppedAutoAbility struct {
	DroppedAutoAbilityInfo
	Chance float64 `json:"chance"`
}

func (a DroppedAutoAbility) Percent() DroppedAutoAbility {
	a.Chance = h.DecimalToPercent(a.Chance)

	return a
}

func (a DroppedAutoAbility) Round(n int32) DroppedAutoAbility {
	a.Chance = h.FloatRound(a.Chance, n)

	return a
}

func convertDroppedAutoAbility(_ *Config, ed seeding.EquipmentDrop) DroppedAutoAbility {
	return DroppedAutoAbility{
		DroppedAutoAbilityInfo: convertDroppedAutoAbilityInfo(ed),
		Chance:                 calcAutoAbilityDropChance(ed),
	}
}

func calcAutoAbilityDropChance(ed seeding.EquipmentDrop) float64 {
	return float64(*ed.Probability) / 7
}