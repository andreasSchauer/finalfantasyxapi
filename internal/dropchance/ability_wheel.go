package dropchance

import (
	"fmt"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

type abilityWheel struct {
	Characters   []seedtypes.Character
	PrioritySlot *string
	Abilities    [7]string
	Clashes      [8][8]bool
	IndexedNames [8]string
}

func (w *abilityWheel) compileClashes() {
	copy(w.IndexedNames[:7], w.Abilities[:])

	if w.PrioritySlot != nil {
		w.IndexedNames[7] = *w.PrioritySlot
	}

	for sourceIdx := range 8 {
		abilityName := w.IndexedNames[sourceIdx]
		if abilityName == "" {
			continue
		}

		ability, _ := lookup.GetResource(abilityName, lookup.L.AutoAbilities)

		for _, lockedOutName := range ability.LockedOutAbilities {
			targetIdxGroup := getTargetIdxGroup(lockedOutName, w.IndexedNames)

			for _, targetIdx := range targetIdxGroup {
				w.Clashes[sourceIdx][targetIdx] = true
				w.Clashes[targetIdx][sourceIdx] = true
			}
		}
	}
}

func getTargetIdxGroup(targetName string, indexedNames [8]string) []int32 {
	var targetIndexGroup []int32
	for targetIdx := range 8 {
		if indexedNames[targetIdx] == targetName {
			targetIndexGroup = append(targetIndexGroup, int32(targetIdx))
		}
	}

	return targetIndexGroup
}

func createAbilityWheels(abilities []seedtypes.EquipmentDrop, params servicetypes.DropChanceParams) []abilityWheel {
	uniqueKeys := getUniqueWheelKeys(abilities, params)
	var wheels []abilityWheel

	if params.CharacterID != nil {
		char, _ := lookup.GetResourceByID(*params.CharacterID, lookup.L.CharactersID)
		wheel := createAbilityWheel(abilities, char)
		wheel.Characters = []seedtypes.Character{char}
		wheels = append(wheels, wheel)
		return wheels
	}

	for _, wheelCharList := range uniqueKeys {
		char := wheelCharList[0]
		wheel := createAbilityWheel(abilities, char)
		wheel.Characters = wheelCharList
		wheels = append(wheels, wheel)
	}

	slices.SortStableFunc(wheels, sortAbilityWheels)

	return wheels
}

func createAbilityWheel(abilities []seedtypes.EquipmentDrop, char seedtypes.Character) abilityWheel {
	var wheel abilityWheel

	idx := 0
	for _, ability := range abilities {
		if !charGetsAbility(ability, &char) {
			continue
		}

		if ability.IsForced {
			abilityName := ability.Ability
			wheel.PrioritySlot = &abilityName
			continue
		}

		wheelSlots := int(*ability.Probability)

		for range wheelSlots {
			wheel.Abilities[idx] = ability.Ability
			idx++
		}
	}

	wheel.compileClashes()
	return wheel
}

func sortAbilityWheels(a, b abilityWheel) int {
	aID := a.Characters[0].ID
	bID := b.Characters[0].ID

	if aID < bID {
		return -1
	}

	if aID > bID {
		return 1
	}

	return 0
}

func getUniqueWheelKeys(abilities []seedtypes.EquipmentDrop, params servicetypes.DropChanceParams) map[string][]seedtypes.Character {
	uniqueKeys := make(map[string][]seedtypes.Character)

	for _, charID := range params.PartyMembers {
		char, _ := lookup.GetResourceByID(charID, lookup.L.CharactersID)
		key := getAbilityWheelKey(abilities, &char)

		_, ok := uniqueKeys[key]
		if ok {
			uniqueKeys[key] = append(uniqueKeys[key], char)
			continue
		}

		uniqueKeys[key] = []seedtypes.Character{char}
	}

	return uniqueKeys
}

func getAbilityWheelKey(abilities []seedtypes.EquipmentDrop, charPtr *seedtypes.Character) string {
	var keys []string

	for _, ability := range abilities {
		if charGetsAbility(ability, charPtr) {
			key := getAbilityKey(ability)
			keys = append(keys, key)
		}
	}

	return strings.Join(keys, "|")
}

func getAbilityKey(ability seedtypes.EquipmentDrop) string {
	if ability.IsForced {
		return fmt.Sprintf("%s_priority", ability.Ability)
	}

	return fmt.Sprintf("%s_%d", ability.Ability, *ability.Probability)
}
