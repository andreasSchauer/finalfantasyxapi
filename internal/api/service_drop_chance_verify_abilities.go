package api

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func vfAutoAbilities(cfg *Config, params DropChanceParams, charPtr *seeding.Character, mon seeding.Monster) ([]seeding.AutoAbility, []seeding.EquipmentDrop, error) {
	if mon.Equipment == nil {
		return nil, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("monster '%s' doesn't drop equipment.", mon), nil)
	}

	params, autoAbilities, err := vfWantedAutoAbilities(cfg, params)
	if err != nil {
		return nil, nil, err
	}

	monsterAbilities, err := getMonAutoAbilities(params, mon)
	if err != nil {
		return nil, nil, err
	}

	err = vfAutoAbilitiesDropped(autoAbilities, monsterAbilities, mon, charPtr)
	if err != nil {
		return nil, nil, err
	}

	return autoAbilities, monsterAbilities, nil
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
		return DropChanceParams{}, nil, newHTTPError(http.StatusBadRequest, "can't discern equipment's type.", nil)
	}

	return params, autoAbilities, nil
}

func getMonAutoAbilities(params DropChanceParams, mon seeding.Monster) ([]seeding.EquipmentDrop, error) {
	switch *params.EquipType {
	case string(database.EquipTypeWeapon):
		return mon.Equipment.WeaponAbilities, nil

	case string(database.EquipTypeArmor):
		return mon.Equipment.ArmorAbilities, nil

	default:
		return nil, newHTTPError(http.StatusBadRequest, "equip type must be set.", nil)
	}
}

func vfAutoAbilitiesDropped(autoAbilities []seeding.AutoAbility, monsterAbilities []seeding.EquipmentDrop, mon seeding.Monster, charPtr *seeding.Character) error {
	for _, ability := range autoAbilities {
		var isPresent bool

		for _, monAbility := range monsterAbilities {
			if ability.Name == monAbility.Ability {
				isPresent = true

				err := vfCharGetsAbility(monAbility, charPtr, mon)
				if err != nil {
					return err
				}
			}
		}

		if !isPresent {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s doesn't drop auto-ability '%s'. it only drops the following auto-abilities: %s.", mon, ability.Name, formatMonAutoAbilities(mon)), nil)
		}
	}

	return nil
}

func vfCharGetsAbility(ability seeding.EquipmentDrop, charPtr *seeding.Character, mon seeding.Monster) error {
	if charPtr == nil {
		return nil
	}
	
	if !charGetsAbility(ability, charPtr) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s doesn't drop auto-ability '%s' for %s", mon, ability.Ability, charPtr.Name), nil)
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
	for _, ability := range monsterAbilities {
		name := fmt.Sprintf("'%s' (%d)", ability.Ability, ability.AutoAbilityID)
		names = append(names, name)
	}

	return strings.Join(names, ", ")
}