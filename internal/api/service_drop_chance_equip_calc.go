package api

import (
	"runtime"
	"slices"
	"sync/atomic"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type EquipmentMatchParams struct {
	WantedAbilities  []seeding.AutoAbility
	WantedIndices 	 [][]int32
	LenientAbilities bool
	MinEmptySlots    *int32
	EquipmentSlots   int32
	ShotAmt          int32
	AbilityWheel     AbilityWheel
}


func calcEquipmentMatchChance(p EquipmentMatchParams) float64 {
	baseEquipment, baseEqLen := initEquipment(p)
	clashes := p.AbilityWheel.Clashes
	wantedIndices := p.WantedIndices
	wantedLen := h.Len32(p.WantedAbilities)
	
	if p.ShotAmt == 0 {
		if isMatch(wantedIndices, &baseEquipment, p.MinEmptySlots, p.EquipmentSlots, wantedLen, baseEqLen, p.LenientAbilities) {
			return 1
		}

		return 0
	}

	totalCombinations := h.PowInt(7, p.ShotAmt)
	var globalMatchingRows atomic.Int32
	
	numWorkers := int32(runtime.NumCPU())
	chunkSize := totalCombinations / numWorkers
	workerGate := numWorkers

	for workerID := range numWorkers {
		startIdx := workerID * chunkSize
		endIdx := startIdx + chunkSize

		if workerID == numWorkers - 1 {
			endIdx = totalCombinations
		}

		go func(start, end int32) {
			var localMatches int32

			for i := start; i < end; i++ {
				equipment, eqLen := assembleEquipment(clashes, baseEquipment, p.ShotAmt, p.EquipmentSlots, baseEqLen, i)
		
				if isMatch(wantedIndices, &equipment, p.MinEmptySlots, p.EquipmentSlots, wantedLen, eqLen, p.LenientAbilities) {
					localMatches++
				}
			}

			if localMatches > 0 {
				globalMatchingRows.Add(localMatches)
			}

			atomic.AddInt32(&workerGate, -1)
		}(startIdx, endIdx)
	}

	for atomic.LoadInt32(&workerGate) > 0 {
		runtime.Gosched()
	}

	return float64(globalMatchingRows.Load()) / float64(totalCombinations)
}

func initEquipment(p EquipmentMatchParams) ([4]int32, int32) {
	equipment := [4]int32{-1, -1, -1, -1}
	var eqLen int32

	if p.AbilityWheel.PrioritySlot != nil {
		equipment[0] = 7
		eqLen++
	}

	return equipment, eqLen
}

func assembleEquipment(clashes [8][8]bool, equipment [4]int32, shotAmt, equipmentSlots, eqLen, idx int32) ([4]int32, int32) {
	for range shotAmt {
		if eqLen == equipmentSlots {
            break
        }

        rolledIdx := int32(idx % 7)
        idx /= 7

		if autoAbilityLockedOut(&clashes, &equipment, eqLen, rolledIdx) {
			continue
		}

		if duplicateAutoAbility(&equipment, eqLen, rolledIdx) {
			continue
		}
		
		equipment[eqLen] = rolledIdx
		eqLen++
	}

	return equipment, eqLen
}

func autoAbilityLockedOut(clashes *[8][8]bool, equipment *[4]int32, eqLen, rolledIdx int32) bool {
    for i := range eqLen {
        if clashes[equipment[i]][rolledIdx] {
            return true
        }
    }
    return false
}

func duplicateAutoAbility(equipment *[4]int32, eqLen, rolledIdx int32) bool {
    for i := range eqLen {
        if equipment[i] == rolledIdx {
            return true
        }
    }
    return false
}


func isMatch(wantedIndices [][]int32, equipment *[4]int32, minEmptySlots *int32, equipmentSlots, wantedLen, eqLen int32, lenientAbilities bool) bool {
    if !allAbilitiesPresent(wantedIndices, equipment, eqLen) {
        return false
    }

    if !strictAbilitiesMatch(lenientAbilities, wantedLen, eqLen) {
        return false
    }

    if !emptySlotsPossible(minEmptySlots, equipmentSlots, eqLen) {
        return false
    }

    return true
}

func allAbilitiesPresent(wantedIndices [][]int32, equipment *[4]int32, eqLen int32) bool {
    for _, abilityGroup := range wantedIndices {
        var found bool

        for i := range eqLen {
            if slices.Contains(abilityGroup, equipment[i]) {
                found = true
                break
            }
        }

        if !found {
            return false
        }
    }

    return true
}

func strictAbilitiesMatch(lenientAbilities bool, wantedLen, eqLen int32) bool {
    if !lenientAbilities && wantedLen != eqLen {
        return false
    }

    return true
}

func emptySlotsPossible(minEmptySlots *int32, equipmentSlots, eqLen int32) bool {
    actualEmptySlots := equipmentSlots - eqLen

    if actualEmptySlots < 0 {
        return false
    }
    
    if minEmptySlots != nil && actualEmptySlots < *minEmptySlots {
        return false
    }

    return true
}