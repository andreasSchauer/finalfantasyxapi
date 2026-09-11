package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type EquipmentMatchParams struct {
	WantedAbilities  []seeding.AutoAbility
	LenientAbilities bool
	MinEmptySlots    *int32
	EquipmentSlots   int32
	ShotAmt          int32
	AbilityWheel     AbilityWheel
}

func calcEquipmentMatchChances(cfg *Config, params DropChanceParams, mon seeding.Monster, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop, cc CharacterChances) (float64, float64) {
	slotsTable := extractAbilitySlots(params, mon)
	shotsTable := mon.Equipment.AttachedAbilities.Chances
	wheels := createAbilityWheels(cfg, monAbilities, params)

	var chanceFinBlow float64
	var chanceNoFinBlow float64

	for _, wheel := range wheels {
		wheelWeightFinBlow, wheelWeightNoFinBlow := getWheelWeights(params, wheel, cc)

		for _, slots := range slotsTable {
			slotWeight := h.PercentageToDecimal(slots.Chance)

			for _, shots := range shotsTable {
				shotWeight := h.PercentageToDecimal(shots.Chance)

				matchParams := EquipmentMatchParams{
					WantedAbilities:  wantedAbilities,
					LenientAbilities: params.LenientAbilities,
					MinEmptySlots:    params.MinEmptySlots,
					EquipmentSlots:   slots.Amount,
					ShotAmt:          shots.Amount,
					AbilityWheel:     wheel,
				}

				matchChance := calcEquipmentMatchChance(matchParams)

				chanceFinBlow += matchChance * wheelWeightFinBlow * slotWeight * shotWeight
				chanceNoFinBlow += matchChance * wheelWeightNoFinBlow * slotWeight * shotWeight
			}
		}
	}

	return chanceFinBlow, chanceNoFinBlow
}


func extractAbilitySlots(params DropChanceParams, mon seeding.Monster) []seeding.EquipmentSlotsChance {
	var minSlots int32 = getRequiredSlots(params)
	var maxSlots int32 = 4
	abilitySlotChances := mon.Equipment.AbilitySlots.Chances
	var chances []seeding.EquipmentSlotsChance

	if params.TotalSlots != nil {
		minSlots = *params.TotalSlots
		maxSlots = *params.TotalSlots
	}

	for _, chance := range abilitySlotChances {
		if chance.Amount >= minSlots && chance.Amount <= maxSlots {
			chances = append(chances, chance)
		}
	}

	return chances
}

func getWheelWeights(params DropChanceParams, wheel AbilityWheel, cc CharacterChances) (float64, float64) {
	if params.Character != nil {
		return *cc.CharNoFinBlow, *cc.CharFinBlow
	}

	charRatio := h.FloatLen(wheel.Characters) / float64(cc.EligibleChars)
	wheelWeightFinBlow := charRatio * cc.AnyCharFinBlow
	wheelWeightNoFinBlow := charRatio * cc.AnyCharNoFinBlow

	return wheelWeightFinBlow, wheelWeightNoFinBlow
}

func calcEquipmentMatchChance(p EquipmentMatchParams) float64 {
	if p.ShotAmt == 0 {
		equipment := initEquipment(p)

		if isMatch(p, equipment) {
			return 1
		}

		return 0
	}

	totalCombinations := h.PowInt(7, p.ShotAmt)
	var matchingRows int32 = 0

	for i := range totalCombinations {
		equipment := assembleEquipment(p, i)

		if isMatch(p, equipment) {
			matchingRows++
		}
	}

	return float64(matchingRows) / float64(totalCombinations)
}

func assembleEquipment(p EquipmentMatchParams, idx int32) map[string]bool {
	equipment := initEquipment(p)

	for range p.ShotAmt {
		if len(equipment) == int(p.EquipmentSlots) {
			break
		}

		rolledSlot := idx % 7
		idx /= 7

		rolledAbility := p.AbilityWheel.Abilities[rolledSlot]
		equipment[rolledAbility] = true
	}

	return equipment
}

func initEquipment(p EquipmentMatchParams) map[string]bool {
	equipment := make(map[string]bool)

	if p.AbilityWheel.PrioritySlot != nil {
		equipment[*p.AbilityWheel.PrioritySlot] = true
	}

	return equipment
}

func isMatch(p EquipmentMatchParams, equipment map[string]bool) bool {
	if !allAbilitiesPresent(p, equipment) {
		return false
	}

	if !strictAbilitiesMatch(p, equipment) {
		return false
	}

	if !emptySlotsPossible(p, equipment) {
		return false
	}

	return true
}

func allAbilitiesPresent(p EquipmentMatchParams, equipment map[string]bool) bool {
	for _, ability := range p.WantedAbilities {
		if !equipment[ability.Name] {
			return false
		}
	}

	return true
}

func strictAbilitiesMatch(p EquipmentMatchParams, equipment map[string]bool) bool {
	if !p.LenientAbilities && len(p.WantedAbilities) != len(equipment) {
		return false
	}

	return true
}

func emptySlotsPossible(p EquipmentMatchParams, equipment map[string]bool) bool {
	actualEmptySlots := p.EquipmentSlots - int32(len(equipment))

	if actualEmptySlots < 0 {
		return false
	}

	if p.MinEmptySlots != nil && actualEmptySlots < *p.MinEmptySlots {
		return false
	}

	return true
}
