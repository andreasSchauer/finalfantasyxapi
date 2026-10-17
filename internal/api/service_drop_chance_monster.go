package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type MonsterDropChances struct {
	MonsterID           int32                `json:"monster_id"`
	MonsterName         string               `json:"monster_name"`
	EquipmentDropChance float64              `json:"equipment_drop_chance"`
	SlotsChances        []ChanceAmount       `json:"slots_chances"`
	ShotsChances        []ChanceAmount       `json:"shots_chances"`
	AutoAbilities       DroppedAutoAbilities `json:"auto_abilities"`
}

func (c MonsterDropChances) Percent() MonsterDropChances {
	c.EquipmentDropChance = h.DecimalToPercent(c.EquipmentDropChance)
	c.SlotsChances = h.SlicePercent(c.SlotsChances)
	c.ShotsChances = h.SlicePercent(c.ShotsChances)
	c.AutoAbilities = c.AutoAbilities.Percent()

	return c
}

func (c MonsterDropChances) Round(n int32) MonsterDropChances {
	c.EquipmentDropChance = h.FloatRound(c.EquipmentDropChance, n)
	c.SlotsChances = h.SliceRound(c.SlotsChances, n)
	c.ShotsChances = h.SliceRound(c.ShotsChances, n)
	c.AutoAbilities = c.AutoAbilities.Round(n)

	return c
}

func getMonsterDropChances(cfg *Config, mon seeding.Monster, wantedAbilities []seeding.EquipmentDrop) MonsterDropChances {
	return MonsterDropChances{
		MonsterID:           mon.ID,
		MonsterName:         h.NameToString(mon.Name, mon.Version, mon.Specification),
		EquipmentDropChance: calcMonDropChance(mon),
		SlotsChances:        convertObjSlice(cfg, mon.Equipment.AbilitySlots.Chances, convertChanceAmount),
		ShotsChances:        convertObjSlice(cfg, mon.Equipment.AttachedAbilities.Chances, convertChanceAmount),
		AutoAbilities:       getDroppedAutoAbilities(cfg, wantedAbilities),
	}
}

func calcMonDropChance(mon seeding.Monster) float64 {
	return float64(mon.Equipment.DropChance) / 255
}

type ChanceAmount struct {
	Amount int32   `json:"amount"`
	Chance float64 `json:"chance"`
}

func (c ChanceAmount) Percent() ChanceAmount {
	c.Chance = h.DecimalToPercent(c.Chance)

	return c
}

func (c ChanceAmount) Round(n int32) ChanceAmount {
	c.Chance = h.FloatRound(c.Chance, n)

	return c
}

func convertChanceAmount(_ *Config, c seeding.EquipmentSlotsChance) ChanceAmount {
	return ChanceAmount{
		Amount: c.Amount,
		Chance: h.PercentToDecimal(float64(c.Chance)),
	}
}
