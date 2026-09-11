package api


import (
	"fmt"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AbilityWheel struct {
	Characters			[]seeding.Character
	PrioritySlot		*string
	Abilities			[7]string
}

func createAbilityWheels(cfg *Config, abilities []seeding.EquipmentDrop, params DropChanceParams) []AbilityWheel {
	uniqueKeys := getUniqueWheelKeys(cfg, abilities, params)
	var wheels []AbilityWheel

	if params.Character != nil {
		char, _ := seeding.GetResourceByID(*params.Character, cfg.l.CharactersID)
		wheel := createAbilityWheel(abilities, char)
		wheel.Characters = []seeding.Character{char}
		wheels = append(wheels, wheel)
		return wheels
	}

	for _, chars := range uniqueKeys {
		char := chars[0]
		wheel := createAbilityWheel(abilities, char)
		wheel.Characters = chars
		wheels = append(wheels, wheel)
	}

	slices.SortStableFunc(wheels, sortAbilityWheels)

	return wheels
}


func createAbilityWheel(abilities []seeding.EquipmentDrop, char seeding.Character) AbilityWheel {
	var wheel AbilityWheel

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

	return wheel
}

func sortAbilityWheels(a, b AbilityWheel) int {
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

func getUniqueWheelKeys(cfg *Config, abilities []seeding.EquipmentDrop, params DropChanceParams) map[string][]seeding.Character {
	uniqueKeys := make(map[string][]seeding.Character)

	for _, charID := range params.PartyMembers {
		char, _ := seeding.GetResourceByID(charID, cfg.l.CharactersID)
		key := getAbilityWheelKey(abilities, &char)

		_, ok := uniqueKeys[key]
		if ok {
			uniqueKeys[key] = append(uniqueKeys[key], char)
			continue
		}

		uniqueKeys[key] = []seeding.Character{char}
	}

	return uniqueKeys
}


func getAbilityWheelKey(abilities []seeding.EquipmentDrop, charPtr *seeding.Character) string {
	var keys []string

	for _, ability := range abilities {
		if charGetsAbility(ability, charPtr) {
			key := getAbilityKey(ability)
			keys = append(keys, key)
		}
	}

	return strings.Join(keys, "|")
}

func getAbilityKey(ability seeding.EquipmentDrop) string {
	if ability.IsForced {
		return fmt.Sprintf("%s_priority", ability.Ability)
	}

	return fmt.Sprintf("%s_%d", ability.Ability, *ability.Probability)
}