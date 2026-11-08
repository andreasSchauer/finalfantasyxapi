package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

type MatchChanceParams struct {
	mon              seedtypes.Monster
	wantedAbilities  []seedtypes.AutoAbility
	monAbilities     []seedtypes.EquipmentDrop
	slotsTable       []seedtypes.EquipmentSlotsChance
	shotsTable       []seedtypes.EquipmentSlotsChance
	characterChances servicetypes.CharacterChances
}

func (p MatchChanceParams) calcChances(params servicetypes.DropChanceParams) servicetypes.EquipmentMatchChances {
	equipTypeChance := 0.5
	wheels := createAbilityWheels(p.monAbilities, params)

	var chanceFinBlow float64
	var chanceNoFinBlow float64

	for _, wheel := range wheels {
		wantedIndices, err := getWantedIndices(p.wantedAbilities, wheel)
		if err != nil {
			continue
		}

		wheelWeightFinBlow, wheelWeightNoFinBlow := getWheelWeights(wheel, p.characterChances)

		for _, slots := range p.slotsTable {
			slotWeight := h.PercentageToDecimal(slots.Chance)

			for _, shots := range p.shotsTable {
				shotWeight := h.PercentageToDecimal(shots.Chance)

				matchParams := EquipmentMatchParams{
					WantedAbilities:   p.wantedAbilities,
					WantedIndices:     wantedIndices,
					OptionalAbilities: params.OptionalAbilities,
					MinEmptySlots:     params.MinEmptySlots,
					EquipmentSlots:    slots.Amount,
					ShotAmt:           shots.Amount,
					AbilityWheel:      wheel,
				}

				matchChance := matchParams.calcChance()

				chanceFinBlow += matchChance * wheelWeightFinBlow * slotWeight * shotWeight
				chanceNoFinBlow += matchChance * wheelWeightNoFinBlow * slotWeight * shotWeight
			}
		}
	}

	chanceFinBlow *= equipTypeChance
	chanceNoFinBlow *= equipTypeChance

	return servicetypes.EquipmentMatchChances{
		FinBlow:   chanceFinBlow,
		NoFinBlow: chanceNoFinBlow,
	}
}

func extractAbilitySlots(params servicetypes.DropChanceParams, mon seedtypes.Monster) []seedtypes.EquipmentSlotsChance {
	var minSlots = params.GetRequiredSlots()
	var maxSlots int32 = 4
	abilitySlotChances := mon.Equipment.AbilitySlots.Chances
	var chances []seedtypes.EquipmentSlotsChance

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

func getWheelWeights(wheel AbilityWheel, cc servicetypes.CharacterChances) (float64, float64) {
	if cc.SpecificChar != nil {
		specificChar := *cc.SpecificChar
		return specificChar.FinBlow, specificChar.NoFinBlow
	}
	eligibleCharsAmt := len(cc.EligibleChars)

	if eligibleCharsAmt == 0 {
		return 0, 0
	}

	charRatio := h.FloatLen(wheel.Characters) / float64(eligibleCharsAmt)
	wheelWeightFinBlow := charRatio * cc.AnyFinBlow
	wheelWeightNoFinBlow := charRatio * cc.AnyNoFinBlow

	return wheelWeightFinBlow, wheelWeightNoFinBlow
}

func getWantedIndices(wantedAbilities []seedtypes.AutoAbility, wheel AbilityWheel) ([][]int32, error) {
	var idxGroups [][]int32

	for _, ability := range wantedAbilities {
		var abilityIdxGroup []int32
		for i, name := range wheel.IndexedNames {
			if ability.Name == name {
				abilityIdxGroup = append(abilityIdxGroup, int32(i))
			}
		}

		if len(abilityIdxGroup) == 0 {
			return nil, errs.ErrContinue
		}

		idxGroups = append(idxGroups, abilityIdxGroup)
	}

	return idxGroups, nil
}
