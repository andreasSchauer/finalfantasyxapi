package api

type EquipmentAssembler struct {
	indexedNames		[8]string
	clashes				[8][8]bool
	equipment			[4]int32
	eqLen				int32
	equipmentSlots  	int32
	shotAmt          	int32
}

func (a EquipmentAssembler) assemble(idx int32) ([4]int32, int32) {
	for range a.shotAmt {
		if a.eqLen == a.equipmentSlots {
            break
        }

        rolledIdx := int32(idx % 7) // can only ever be 0 to 6 | 55 % 7 = 6 (49 + 6)
        idx /= 7 // 55 / 7 = 7.87 => 7 | oldIdx = newIdx * 7 + rolledIdx | 55 = 7 * 7 + 6
		// newIdx will be moduloed further to get a new rolledIdx (7 % 7 = 0)
		
		if !a.autoAbilityValid(rolledIdx) {
			continue
		}

		a.equipment[a.eqLen] = rolledIdx
		a.eqLen++
	}

	return a.equipment, a.eqLen
}

func (a *EquipmentAssembler) autoAbilityValid(rolledIdx int32) bool {
    for i := range a.eqLen {
		// autoAbilityLockedOut
        if a.clashes[a.equipment[i]][rolledIdx] {
            return false
        }

		// duplicateAutoAbility
		if a.indexedNames[a.equipment[i]] == a.indexedNames[rolledIdx] {
            return false
        }
    }
    return true
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