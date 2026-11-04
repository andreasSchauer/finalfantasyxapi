package api

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func vfAutoAbilities(params DropChanceParams, mon seedtypes.Monster) ([]seedtypes.AutoAbility, []seedtypes.EquipmentDrop, error) {
	if mon.Equipment == nil {
		return nil, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("monster '%s' doesn't drop equipment", mon), nil)
	}

	monMaxSlots := getMonMaxSlots(mon)
	reqSlots := params.GetRequiredSlots()

	if reqSlots > monMaxSlots {
		return nil, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("query requires %d slots, but the max amount of slots, monster '%s' can drop, is %d", reqSlots, mon, monMaxSlots), nil)
	}

	params, autoAbilities, err := vfWantedAutoAbilities(params)
	if err != nil {
		return nil, nil, err
	}

	monsterAbilities := getMonAutoAbilities(*params.EquipType, mon) // equipType will never be nil at this point

	err = vfAutoAbilitiesDropped(autoAbilities, monsterAbilities, mon)
	if err != nil {
		return nil, nil, err
	}

	return autoAbilities, monsterAbilities, nil
}

func getEquipMaxSlots(slotsChances []seedtypes.EquipmentSlotsChance) int32 {
	return slotsChances[len(slotsChances)-1].Amount
}

func getMonMaxSlots(mon seedtypes.Monster) int32 {
	if mon.Equipment == nil {
		return 0
	}

	slotsChances := mon.Equipment.AbilitySlots.Chances
	return slotsChances[len(slotsChances)-1].Amount
}

func vfWantedAutoAbilities(params DropChanceParams) (DropChanceParams, []seedtypes.AutoAbility, error) {
	var autoAbilities []seedtypes.AutoAbility

	for _, id := range params.AutoAbilities {
		ability, _ := lookup.GetResourceByID(id, lookup.L.AutoAbilitiesID)

		if params.EquipType == nil {
			params.EquipType = &ability.Type
		}

		if *params.EquipType != ability.Type {
			return DropChanceParams{}, nil, newHTTPError(http.StatusBadRequest, "weapon- and armor-abilities can't be combined", nil)
		}

		autoAbilities = append(autoAbilities, ability)
	}

	if params.EquipType == nil {
		return DropChanceParams{}, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("can't discern equipment's type. use '%s' to discern the type, or set it manually by using '%s'", typestrings.PfnAutoAbilities, typestrings.PfnEquipType), nil)
	}

	return params, autoAbilities, nil
}

func getMonAutoAbilities(equipType string, mon seedtypes.Monster) []seedtypes.EquipmentDrop {
	switch equipType {
	case string(database.EquipTypeWeapon):
		return mon.Equipment.WeaponAbilities

	case string(database.EquipTypeArmor):
		return mon.Equipment.ArmorAbilities
	}

	return nil
}

func vfAutoAbilitiesDropped(wantedAutoAbilities []seedtypes.AutoAbility, monsterAbilities []seedtypes.EquipmentDrop, mon seedtypes.Monster) error {
	for _, ability := range wantedAutoAbilities {
		var isPresent bool

		for _, monAbility := range monsterAbilities {
			if ability.Name == monAbility.Ability {
				isPresent = true
				break
			}
		}

		if !isPresent {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s doesn't drop auto-ability '%s' (%d). it only drops the following auto-abilities: %s", mon, ability.Name, ability.ID, formatMonAutoAbilities(mon)), nil)
		}
	}

	return nil
}

func formatMonAutoAbilities(mon seedtypes.Monster) string {
	if mon.Equipment == nil {
		return ""
	}

	weaponAbilities := mon.Equipment.WeaponAbilities
	armorAbilities := mon.Equipment.ArmorAbilities
	allAbilities := slices.Concat(weaponAbilities, armorAbilities)

	return formatEquipmentDrops(allAbilities)
}

func formatEquipmentDrops(monsterAbilities []seedtypes.EquipmentDrop) string {
	var names []string
	var forcedName string

	for _, ability := range monsterAbilities {
		if ability.Ability == forcedName {
			continue
		}

		if ability.IsForced {
			forcedName = ability.Ability
		}

		name := fmt.Sprintf("'%s' (%d)", ability.Ability, ability.AutoAbilityID)
		names = append(names, name)
	}

	return strings.Join(names, ", ")
}
