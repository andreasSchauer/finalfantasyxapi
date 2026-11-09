package dropchance

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

func getDroppedAutoAbilities(abilities []seedtypes.EquipmentDrop) servicetypes.DroppedAutoAbilities {
	if abilities[0].IsForced {
		forcedAbility := convertEqAbilitySimple(abilities[0])

		return servicetypes.DroppedAutoAbilities{
			Forced:  &forcedAbility,
			Dropped: ah.ConvertObjSlice(abilities[1:], convertEqAbilityChance),
		}
	}

	return servicetypes.DroppedAutoAbilities{
		Dropped: ah.ConvertObjSlice(abilities, convertEqAbilityChance),
	}
}

func convertEqAbilitySimple(ed seedtypes.EquipmentDrop) servicetypes.EqAbilitySimple {
	return servicetypes.EqAbilitySimple{
		ID:         ed.AutoAbilityID,
		Name:       ed.Ability,
		Characters: ed.Characters,
	}
}

func convertEqAbilityChance(ed seedtypes.EquipmentDrop) servicetypes.EqAbilityChance {
	return servicetypes.EqAbilityChance{
		EqAbilitySimple: convertEqAbilitySimple(ed),
		Chance:          calcAutoAbilityDropChance(ed),
	}
}

func calcAutoAbilityDropChance(ed seedtypes.EquipmentDrop) float64 {
	return float64(*ed.Probability) / 7
}
