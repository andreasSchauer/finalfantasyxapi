package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func vfAutoAbilities(cfg *Config, params DropChanceParams, charPtr *seeding.Character, mon seeding.Monster) ([]seeding.AutoAbility, []seeding.EquipmentDrop, error) {
	if mon.Equipment == nil {
		return nil, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s doesn't drop equipment.", mon), nil)
	}

	autoAbilities, monsterAbilities, err := vfMonsterAutoAbilities(cfg, params, mon)
	if err != nil {
		return nil, nil, err
	}

	err = vfAutoAbilitiesDropped(autoAbilities, monsterAbilities, mon, charPtr)
	if err != nil {
		return nil, nil, err
	}

	return autoAbilities, monsterAbilities, nil
}

func vfMonsterAutoAbilities(cfg *Config, params DropChanceParams, mon seeding.Monster) ([]seeding.AutoAbility, []seeding.EquipmentDrop, error) {
	var autoAbilities []seeding.AutoAbility
	var monsterAbilities []seeding.EquipmentDrop
	equipTypes := make(map[string]bool)

	for _, id := range params.AutoAbilities {
		ability, _ := seeding.GetResourceByID(id, cfg.l.AutoAbilitiesID)
		equipTypes[ability.Type] = true

		if len(equipTypes) > 1 {
			return nil, nil, newHTTPError(http.StatusBadRequest, "weapon- and armor-abilities can't be combined.", nil)
		}

		switch ability.Type {
		case string(database.EquipTypeWeapon):
			monsterAbilities = mon.Equipment.WeaponAbilities

		case string(database.EquipTypeArmor):
			monsterAbilities = mon.Equipment.ArmorAbilities
		}

		autoAbilities = append(autoAbilities, ability)
	}

	return autoAbilities, monsterAbilities, nil
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
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s doesn't drop auto-ability '%s'", mon, ability.Name), nil)
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



