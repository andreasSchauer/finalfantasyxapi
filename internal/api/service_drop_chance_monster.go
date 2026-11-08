package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

func getMonsterDropChances(mon seedtypes.Monster, wantedAbilities []seedtypes.EquipmentDrop) servicetypes.MonsterDropChances {
	return servicetypes.MonsterDropChances{
		MonsterID:           mon.ID,
		MonsterName:         h.NameToString(mon.Name, mon.Version, mon.Specification),
		EquipmentDropChance: calcMonDropChance(mon),
		SlotsChances:        ah.ConvertObjSlice(mon.Equipment.AbilitySlots.Chances, convertChanceAmount),
		ShotsChances:        ah.ConvertObjSlice(mon.Equipment.AttachedAbilities.Chances, convertChanceAmount),
		AutoAbilities:       getDroppedAutoAbilities(wantedAbilities),
	}
}

func calcMonDropChance(mon seedtypes.Monster) float64 {
	return float64(mon.Equipment.DropChance) / 255
}

func convertChanceAmount(c seedtypes.EquipmentSlotsChance) servicetypes.ChanceAmount {
	return servicetypes.ChanceAmount{
		Amount: c.Amount,
		Chance: h.PercentToDecimal(float64(c.Chance)),
	}
}
