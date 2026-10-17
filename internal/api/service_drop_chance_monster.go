package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type MonsterDropChances struct {
	MonsterID            int32                `json:"monster_id"`
	MonsterName          string               `json:"monster_name"`
	EquipmentDropChance  float64              `json:"equipment_drop_chance"`
	SlotsChances         []SlotsChance        `json:"slots_chances"`
	ShotsChances         []SlotsChance        `json:"shots_chances"`
	DroppedAutoAbilities DroppedAutoAbilities `json:"dropped_auto_abilities"`
}

func (c MonsterDropChances) Percent() MonsterDropChances {
	c.EquipmentDropChance = h.DecimalToPercent(c.EquipmentDropChance)
	c.SlotsChances = h.SlicePercent(c.SlotsChances)
	c.ShotsChances = h.SlicePercent(c.ShotsChances)
	c.DroppedAutoAbilities = c.DroppedAutoAbilities.Percent()

	return c
}

func (c MonsterDropChances) Round(n int32) MonsterDropChances {
	c.EquipmentDropChance = h.FloatRound(c.EquipmentDropChance, n)
	c.SlotsChances = h.SliceRound(c.SlotsChances, n)
	c.ShotsChances = h.SliceRound(c.ShotsChances, n)
	c.DroppedAutoAbilities = c.DroppedAutoAbilities.Round(n)

	return c
}

func getMonsterDropChances(cfg *Config, mon seeding.Monster, wantedAbilities []seeding.EquipmentDrop) MonsterDropChances {
	return MonsterDropChances{
		MonsterID:            mon.ID,
		MonsterName:          h.NameToString(mon.Name, mon.Version, mon.Specification),
		EquipmentDropChance:  calcMonDropChance(mon),
		SlotsChances:         convertObjSlice(cfg, mon.Equipment.AbilitySlots.Chances, convertSlotsChance),
		ShotsChances:         convertObjSlice(cfg, mon.Equipment.AttachedAbilities.Chances, convertSlotsChance),
		DroppedAutoAbilities: getDroppedAutoAbilities(cfg, wantedAbilities),
	}
}

func calcMonDropChance(mon seeding.Monster) float64 {
	return float64(mon.Equipment.DropChance) / 255
}

type SlotsChance struct {
	Amount int32   `json:"amount"`
	Chance float64 `json:"chance"`
}

func (c SlotsChance) Percent() SlotsChance {
	c.Chance = h.DecimalToPercent(c.Chance)

	return c
}

func (c SlotsChance) Round(n int32) SlotsChance {
	c.Chance = h.FloatRound(c.Chance, n)

	return c
}

func convertSlotsChance(_ *Config, c seeding.EquipmentSlotsChance) SlotsChance {
	return SlotsChance{
		Amount: c.Amount,
		Chance: h.PercentToDecimal(float64(c.Chance)),
	}
}