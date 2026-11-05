package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type DroppedAutoAbilities struct {
	Forced  *EqAbilitySimple  `json:"forced"`
	Dropped []EqAbilityChance `json:"dropped"`
}

func (d DroppedAutoAbilities) Percent() DroppedAutoAbilities {
	d.Dropped = h.SlicePercent(d.Dropped)

	return d
}

func (d DroppedAutoAbilities) Round(n int32) DroppedAutoAbilities {
	d.Dropped = h.SliceRound(d.Dropped, n)

	return d
}

func getDroppedAutoAbilities(abilities []seedtypes.EquipmentDrop) DroppedAutoAbilities {
	if abilities[0].IsForced {
		forcedAbility := convertEqAbilitySimple(abilities[0])

		return DroppedAutoAbilities{
			Forced:  &forcedAbility,
			Dropped: convertObjSlice(abilities[1:], convertEqAbilityChance),
		}
	}

	return DroppedAutoAbilities{
		Dropped: convertObjSlice(abilities, convertEqAbilityChance),
	}
}

type EqAbilitySimple struct {
	ID         int32    `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

func convertEqAbilitySimple(ed seedtypes.EquipmentDrop) EqAbilitySimple {
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

func convertEqAbilityChance(ed seedtypes.EquipmentDrop) EqAbilityChance {
	return EqAbilityChance{
		EqAbilitySimple: convertEqAbilitySimple(ed),
		Chance:          calcAutoAbilityDropChance(ed),
	}
}

func calcAutoAbilityDropChance(ed seedtypes.EquipmentDrop) float64 {
	return float64(*ed.Probability) / 7
}
