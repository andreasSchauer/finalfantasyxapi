package api

import (
	"runtime"
	"sync"
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

	matcher := EquipmentMatcher{
		wantedIndices: 		p.WantedIndices,
		wantedLen:			h.Len32(p.WantedAbilities),
		minEmptySlots: 		p.MinEmptySlots,
		equipmentSlots: 	p.EquipmentSlots,
		lenientAbilities: 	p.LenientAbilities,
	}
	
	if p.ShotAmt == 0 {
		if matcher.Match(&baseEquipment, baseEqLen) {
			return 1
		}

		return 0
	}

	assembler := EquipmentAssembler{
		indexedNames: 	p.AbilityWheel.IndexedNames,
		clashes: 		p.AbilityWheel.Clashes,
		equipment: 		baseEquipment,
		eqLen: 			baseEqLen,
		equipmentSlots: p.EquipmentSlots,
		shotAmt: 		p.ShotAmt,
	}

	totalCombinations := h.PowInt(7, p.ShotAmt)
	var globalMatchingRows atomic.Int32
	numWorkers := int32(runtime.NumCPU())
	chunkSize := totalCombinations / numWorkers
	var wg sync.WaitGroup

	for workerID := range numWorkers {
		startIdx := workerID * chunkSize
		endIdx := startIdx + chunkSize

		if workerID == numWorkers - 1 {
			endIdx = totalCombinations
		}

		wg.Add(1)
		go func(start, end int32) {
			defer wg.Done()
			var localMatches int32

			for i := start; i < end; i++ {
				equipment, eqLen := assembler.assemble(i)
		
				if matcher.Match(&equipment, eqLen) {
					localMatches++
				}
			}

			if localMatches > 0 {
				globalMatchingRows.Add(localMatches)
			}
		}(startIdx, endIdx)
	}

	wg.Wait()

	return float64(globalMatchingRows.Load()) / float64(totalCombinations)
}