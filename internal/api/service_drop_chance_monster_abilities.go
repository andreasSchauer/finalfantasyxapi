package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type DroppedAbilities struct {
	Forced  *EqAbilitySimple  `json:"forced"`
	Dropped []EqAbilityChance `json:"dropped"`
}

func (d DroppedAbilities) Percent() DroppedAbilities {
	d.Dropped = h.SlicePercent(d.Dropped)

	return d
}

func (d DroppedAbilities) Round(n int32) DroppedAbilities {
	d.Dropped = h.SliceRound(d.Dropped, n)

	return d
}

func getDroppedAutoAbilities(cfg *Config, abilities []seeding.EquipmentDrop) DroppedAbilities {
	if abilities[0].IsForced {
		forcedAbility := convertEqAbilitySimple(abilities[0])

		return DroppedAbilities{
			Forced:  &forcedAbility,
			Dropped: convertObjSlice(cfg, abilities[1:], convertEqAbilityChance),
		}
	}

	return DroppedAbilities{
		Dropped: convertObjSlice(cfg, abilities, convertEqAbilityChance),
	}
}

type EqAbilitySimple struct {
	ID         int32    `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

func convertEqAbilitySimple(ed seeding.EquipmentDrop) EqAbilitySimple {
	return EqAbilitySimple{
		ID:         ed.AutoAbilityID,
		Name:       ed.Ability,
		Characters: ed.Characters,
	}
}

type EqAbilityChance struct {
	EqAbilitySimple
	Chance float64 `json:"chance"`
}

func (a EqAbilityChance) Percent() EqAbilityChance {
	a.Chance = h.DecimalToPercent(a.Chance)

	return a
}

func (a EqAbilityChance) Round(n int32) EqAbilityChance {
	a.Chance = h.FloatRound(a.Chance, n)

	return a
}

func convertEqAbilityChance(_ *Config, ed seeding.EquipmentDrop) EqAbilityChance {
	return EqAbilityChance{
		EqAbilitySimple: convertEqAbilitySimple(ed),
		Chance:          calcAutoAbilityDropChance(ed),
	}
}

func calcAutoAbilityDropChance(ed seeding.EquipmentDrop) float64 {
	return float64(*ed.Probability) / 7
}
