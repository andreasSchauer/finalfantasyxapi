package api

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func vfAutoAbilities(cfg *Config, params DropChanceParams, mon seeding.Monster) ([]seeding.AutoAbility, []seeding.EquipmentDrop, error) {
	if mon.Equipment == nil {
		return nil, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("monster '%s' doesn't drop equipment.", mon), nil)
	}
	
	monMaxSlots := getMonMaxSlots(mon)
	reqSlots := params.GetRequiredSlots()

	if reqSlots > monMaxSlots {
		return nil, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("query requires %d slots, but the max amount of slots, monster '%s' can drop, is %d.", reqSlots, mon, monMaxSlots), nil)
	}

	params, autoAbilities, err := vfWantedAutoAbilities(cfg, params)
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

func getEquipMaxSlots(slotsChances []seeding.EquipmentSlotsChance) int32 {
	return slotsChances[len(slotsChances)-1].Amount
}

func getMonMaxSlots(mon seeding.Monster) int32 {
	if mon.Equipment == nil {
		return 0
	}

	slotsChances := mon.Equipment.AbilitySlots.Chances
	return slotsChances[len(slotsChances)-1].Amount
}

func vfWantedAutoAbilities(cfg *Config, params DropChanceParams) (DropChanceParams, []seeding.AutoAbility, error) {
	var autoAbilities []seeding.AutoAbility
	
	for _, id := range params.AutoAbilities {
		ability, _ := seeding.GetResourceByID(id, cfg.l.AutoAbilitiesID)
		
		if params.EquipType == nil {
			params.EquipType = &ability.Type
		}
		
		if *params.EquipType != ability.Type {
			return DropChanceParams{}, nil, newHTTPError(http.StatusBadRequest, "weapon- and armor-abilities can't be combined.", nil)
		}
		
		autoAbilities = append(autoAbilities, ability)
	}

	if params.EquipType == nil {
		return DropChanceParams{}, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("can't discern equipment's type. use '%s' to discern the type, or set it manually by using '%s'.", pfnAutoAbilities, pfnEquipType), nil)
	}

	return params, autoAbilities, nil
}

func getMonAutoAbilities(equipType string, mon seeding.Monster) []seeding.EquipmentDrop {
	switch equipType {
	case string(database.EquipTypeWeapon):
		return mon.Equipment.WeaponAbilities

	case string(database.EquipTypeArmor):
		return mon.Equipment.ArmorAbilities
	}

	return nil
}

func vfAutoAbilitiesDropped(wantedAutoAbilities []seeding.AutoAbility, monsterAbilities []seeding.EquipmentDrop, mon seeding.Monster) error {
	for _, ability := range wantedAutoAbilities {
		var isPresent bool

		for _, monAbility := range monsterAbilities {
			if ability.Name == monAbility.Ability {
				isPresent = true
				break
			}
		}

		if !isPresent {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s doesn't drop auto-ability '%s' (%d). it only drops the following auto-abilities: %s.", mon, ability.Name, ability.ID, formatMonAutoAbilities(mon)), nil)
		}
	}

	return nil
}

func formatMonAutoAbilities(mon seeding.Monster) string {
	if mon.Equipment == nil {
		return ""
	}

	weaponAbilities := mon.Equipment.WeaponAbilities
	armorAbilities := mon.Equipment.ArmorAbilities
	allAbilities := slices.Concat(weaponAbilities, armorAbilities)

	return formatEquipmentDrops(allAbilities)
}

func formatEquipmentDrops(monsterAbilities []seeding.EquipmentDrop) string {
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