package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func calcEquipmentMatchChances(cfg *Config, params DropChanceParams, mon seeding.Monster, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop, cc CharacterChances) (float64, float64) {
	equipTypeChance := 0.5
	slotsTable := extractAbilitySlots(params, mon)
	shotsTable := mon.Equipment.AttachedAbilities.Chances
	wheels := createAbilityWheels(cfg, monAbilities, params)

	var chanceFinBlow float64
	var chanceNoFinBlow float64

	for _, wheel := range wheels {
		wantedIndices, err := getWantedIndices(wantedAbilities, wheel)
		if err != nil {
			continue
		}
		
		wheelWeightFinBlow, wheelWeightNoFinBlow := getWheelWeights(params, wheel, cc)

		for _, slots := range slotsTable {
			slotWeight := h.PercentageToDecimal(slots.Chance)

			for _, shots := range shotsTable {
				shotWeight := h.PercentageToDecimal(shots.Chance)

				matchParams := EquipmentMatchParams{
					WantedAbilities:  wantedAbilities,
					WantedIndices:	  wantedIndices,
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

	chanceFinBlow *= equipTypeChance
	chanceNoFinBlow *= equipTypeChance

	return chanceFinBlow, chanceNoFinBlow
}

func getWantedIndices(wantedAbilities []seeding.AutoAbility, wheel AbilityWheel) ([][]int32, error) {
	var indices [][]int32

	for _, ability := range wantedAbilities {
		var group []int32
		for i, name := range wheel.IndexedNames {
			if ability.Name == name {
				group = append(group, int32(i))
			}
		}

		if len(group) == 0 {
			return nil, errContinue
		}

		indices = append(indices, group)
	}

	return indices, nil
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
		return *cc.CharFinBlow, *cc.CharNoFinBlow
	}

	charRatio := h.FloatLen(wheel.Characters) / float64(cc.EligibleChars)
	wheelWeightFinBlow := charRatio * cc.AnyCharFinBlow
	wheelWeightNoFinBlow := charRatio * cc.AnyCharNoFinBlow

	return wheelWeightFinBlow, wheelWeightNoFinBlow
}