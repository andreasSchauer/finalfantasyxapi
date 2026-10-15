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
	wantedLen := h.Len32(p.WantedAbilities)
	
	if p.ShotAmt == 0 {
		if isMatch(p, &baseEquipment, wantedLen, baseEqLen) {
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
				equipment, eqLen := assembleEquipment(p, baseEquipment, baseEqLen, i)
		
				if isMatch(p, &equipment, wantedLen, eqLen) {
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
	const idxPrioritySlot int32 = 7
	const idxEmpty int32 = -1
	var eqLen int32
	equipment := [4]int32{idxEmpty, idxEmpty, idxEmpty, idxEmpty}

	if p.AbilityWheel.PrioritySlot != nil {
		equipment[0] = idxPrioritySlot
		eqLen++
	}

	return equipment, eqLen
}

func assembleEquipment(p EquipmentMatchParams, equipment [4]int32, eqLen, idx int32) ([4]int32, int32) {
	for range p.ShotAmt {
		if eqLen == p.EquipmentSlots {
            break
        }

        rolledIdx := int32(idx % 7) // can only ever be 0 to 6 | 55 % 7 = 6 (49 + 6)
        idx /= 7 // 55 / 7 = 7.87 => 7 | oldIdx = newIdx * 7 + rolledIdx | 55 = 7 * 7 + 6
		// newIdx will be moduloed further to get a new rolledIdx (7 % 7 = 0)
		
		if autoAbilityLockedOut(&p.AbilityWheel.Clashes, &equipment, eqLen, rolledIdx) {
			continue
		}
		
		if duplicateAutoAbility(p.AbilityWheel.IdxGroups, &equipment, eqLen, rolledIdx) {
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

func duplicateAutoAbility(wheelIndices [][]int32, equipment *[4]int32, eqLen, rolledIdx int32) bool {
    for i := range eqLen {
		eqGroupIdx := getIdxGroupIndex(wheelIndices, equipment[i])
		rolledGroupIdx := getIdxGroupIndex(wheelIndices, rolledIdx)

        if eqGroupIdx == rolledGroupIdx {
            return true
        }
    }
    return false
}

func getIdxGroupIndex(wantedIndices [][]int32, idx int32) int32 {
	for i, wantedIdxGroup := range wantedIndices {
		if slices.Contains(wantedIdxGroup, idx) {
			return int32(i)
		}
	}

	return -1
}


func isMatch(p EquipmentMatchParams, equipment *[4]int32, wantedLen, eqLen int32) bool {
    if !allWantedAbilitiesPresent(p.WantedIndices, equipment, eqLen) {
        return false
    }

	// if I want piercing and have piercing and strength +5%,
	// then all wanted abilities are present
	// but the abilities don't match with what I want (just piercing)
    if !strictAbilitiesMatch(p.LenientAbilities, wantedLen, eqLen) {
        return false
    }

    if !emptySlotsPossible(p.MinEmptySlots, p.EquipmentSlots, eqLen) {
        return false
    }

    return true
}

func allWantedAbilitiesPresent(wantedIndices [][]int32, equipment *[4]int32, eqLen int32) bool {
    for _, wantedIdxGroup := range wantedIndices {
        if !wantedAbilityPresent(wantedIdxGroup, equipment, eqLen) {
            return false
        }
    }

    return true
}

func wantedAbilityPresent(wantedIdxGroup []int32, equipment *[4]int32, eqLen int32) bool {
	for i := range eqLen {
		if slices.Contains(wantedIdxGroup, equipment[i]) {
			return true
		}
	}

	return false
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