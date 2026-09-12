package api

import (
	"runtime"
	"sync/atomic"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type EquipmentMatchParams struct {
	WantedAbilities  []seeding.AutoAbility
	WantedIndeces 	 []int32
	LenientAbilities bool
	MinEmptySlots    *int32
	EquipmentSlots   int32
	ShotAmt          int32
	AbilityWheel     AbilityWheel
}

func calcEquipmentMatchChances(cfg *Config, params DropChanceParams, mon seeding.Monster, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop, cc CharacterChances) (float64, float64) {
	equipTypeChance := 0.5
	slotsTable := extractAbilitySlots(params, mon)
	shotsTable := mon.Equipment.AttachedAbilities.Chances
	wheels := createAbilityWheels(cfg, monAbilities, params)

	var chanceFinBlow float64
	var chanceNoFinBlow float64

	for _, wheel := range wheels {
		wheelWeightFinBlow, wheelWeightNoFinBlow := getWheelWeights(params, wheel, cc)
		wantedIndeces := getWantedIndeces(wantedAbilities, wheel)

		for _, slots := range slotsTable {
			slotWeight := h.PercentageToDecimal(slots.Chance)

			for _, shots := range shotsTable {
				shotWeight := h.PercentageToDecimal(shots.Chance)

				matchParams := EquipmentMatchParams{
					WantedAbilities:  wantedAbilities,
					WantedIndeces:	  wantedIndeces,
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

func getWantedIndeces(wantedAbilities []seeding.AutoAbility, wheel AbilityWheel) []int32 {
	indeces := make([]int32, 0, len(wantedAbilities))

	for _, ability := range wantedAbilities {
		idx := getTargetIdx(ability.Name, wheel.IndexedNames)
		indeces = append(indeces, idx)
	}

	return indeces
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

func calcEquipmentMatchChance(p EquipmentMatchParams) float64 {
	baseEquipment, baseEqLen := initEquipment(p)
	clashes := p.AbilityWheel.Clashes
	wantedIndeces := p.WantedIndeces
	wantedLen := h.Len32(p.WantedAbilities)
	
	if p.ShotAmt == 0 {
		if isMatch(wantedIndeces, &baseEquipment, p.MinEmptySlots, p.EquipmentSlots, wantedLen, baseEqLen, p.LenientAbilities) {
			return 1
		}

		return 0
	}

	totalCombinations := h.PowInt(7, p.ShotAmt)
	numWorkers := int32(runtime.NumCPU())
	chunkSize := totalCombinations / numWorkers
	
	var globalMatchingRows int32
	workerGate := numWorkers

	for workerID := int32(0); workerID < numWorkers; workerID++ {
		startIdx := workerID * chunkSize
		endIdx := startIdx + chunkSize

		if workerID == numWorkers - 1 {
			endIdx = totalCombinations
		}

		go func(start, end int32) {
			var localMatches int32

			for i := start; i < end; i++ {
				equipment, eqLen := assembleEquipment(clashes, baseEquipment, p.ShotAmt, p.EquipmentSlots, baseEqLen, i)
		
				if isMatch(wantedIndeces, &equipment, p.MinEmptySlots, p.EquipmentSlots, wantedLen, eqLen, p.LenientAbilities) {
					localMatches++
				}
			}

			if localMatches > 0 {
				atomic.AddInt32(&globalMatchingRows, localMatches)
			}

			atomic.AddInt32(&workerGate, -1)
		}(startIdx, endIdx)
	}

	for atomic.LoadInt32(&workerGate) > 0 {
		runtime.Gosched()
	}

	return float64(atomic.LoadInt32(&globalMatchingRows)) / float64(totalCombinations)
}

func assembleEquipment(clashes [8][8]bool, equipment [4]int32, shotAmt, equipmentSlots, eqLen, idx int32) ([4]int32, int32) {
	for range shotAmt {
		if eqLen == equipmentSlots {
			break
		}

		rolledSlot := idx & 7
		idx >>= 3
		rolledIdx := int32(rolledSlot)

		if rolledIdx == 7 {
			return equipment, -1
		}

		var lockedOut bool
		for i := range eqLen {
			if clashes[equipment[i]][rolledIdx] {
				lockedOut = true
				break
			}
		}
		if lockedOut { continue }

		var duplicateAbility bool
		for i := range eqLen {
			if equipment[i] == rolledIdx {
				duplicateAbility = true
				break
			}
		}
		if duplicateAbility { continue }
		
		equipment[eqLen] = rolledIdx
		eqLen++
	}

	return equipment, eqLen
}

func initEquipment(p EquipmentMatchParams) ([4]int32, int32) {
	var equipment [4]int32
	var eqLen int32

	if p.AbilityWheel.PrioritySlot != nil {
		equipment[0] = 7
		eqLen++
	}

	return equipment, eqLen
}


func isMatch(wantedIndices []int32, equipment *[4]int32, minEmptySlots *int32, equipmentSlots, wantedLen, eqLen int32, lenientAbilities bool) bool {
	// all abilities are present
	for _, wantedIdx := range wantedIndices {
		var found bool

		for i := range eqLen {
			if equipment[i] == wantedIdx {
				found = true
				break
			}
		}

		if !found { return false }
	}

	// strict abilities match
	if !lenientAbilities && wantedLen != eqLen {
		return false
	}

	// empty slots amount is possible to get
	actualEmptySlots := equipmentSlots - eqLen

	if actualEmptySlots < 0 {
		return false
	}
	if minEmptySlots != nil && actualEmptySlots < *minEmptySlots {
		return false
	}

	return true
}