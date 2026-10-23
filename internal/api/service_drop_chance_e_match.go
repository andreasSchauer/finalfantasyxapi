package api

import "slices"

type EquipmentMatcher struct {
	wantedIndices     [][]int32
	wantedLen         int32
	minEmptySlots     *int32
	equipmentSlots    int32
	optionalAbilities bool
}

func (m *EquipmentMatcher) Match(equipment *[4]int32, eqLen int32) bool {
	if !m.allWantedAbilitiesPresent(equipment, eqLen) {
		return false
	}

	// if I want piercing and have piercing and strength +5%,
	// then all wanted abilities are present
	// but the abilities don't match with what I want (just piercing)
	if !m.strictAbilitiesMatch(eqLen) {
		return false
	}

	if !m.emptySlotsPossible(eqLen) {
		return false
	}

	return true
}

func (m *EquipmentMatcher) allWantedAbilitiesPresent(equipment *[4]int32, eqLen int32) bool {
	for _, wantedIdxGroup := range m.wantedIndices {
		var found bool

		for i := range eqLen {
			if slices.Contains(wantedIdxGroup, equipment[i]) {
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

func (m *EquipmentMatcher) strictAbilitiesMatch(eqLen int32) bool {
	if !m.optionalAbilities && m.wantedLen != eqLen {
		return false
	}

	return true
}

func (m *EquipmentMatcher) emptySlotsPossible(eqLen int32) bool {
	actualEmptySlots := m.equipmentSlots - eqLen

	if actualEmptySlots < 0 {
		return false
	}

	if m.minEmptySlots != nil && actualEmptySlots < *m.minEmptySlots {
		return false
	}

	return true
}
