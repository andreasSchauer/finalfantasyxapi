package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type MonsterEquipment struct {
	DropChance        int32                 `json:"drop_chance"`
	Power             int32                 `json:"power"`
	CriticalPlus      int32                 `json:"critical_plus"`
	AbilitySlots      MonsterEquipmentSlots `json:"ability_slots"`
	AttachedAbilities MonsterEquipmentSlots `json:"attached_abilities"`
	WeaponAbilities   []EquipmentDrop       `json:"weapon_abilities"`
	ArmorAbilities    []EquipmentDrop       `json:"armor_abilities"`
}

func (me MonsterEquipment) IsZero() bool {
	return me.DropChance == 0
}

type MonsterEquipmentSlots struct {
	MinAmount int32                  `json:"min_amount"`
	MaxAmount int32                  `json:"max_amount"`
	Chances   []EquipmentSlotsChance `json:"chances"`
}

type EquipmentSlotsChance struct {
	Amount int32 `json:"amount"`
	Chance int32 `json:"chance"`
}

type EquipmentDrop struct {
	AutoAbility NamedAPIResource   `json:"auto_ability"`
	ForcedChars []NamedAPIResource `json:"forced_characters"`
	IsForced    bool               `json:"is_forced"`
	Probability *int32             `json:"probability,omitempty"`
}

func (ed EquipmentDrop) GetAPIResource() IsAPIResource {
	return ed.AutoAbility.GetAPIResource()
}

func (cfg *Config) getMonsterEquipment(r *http.Request, mon database.Monster) (MonsterEquipment, error) {
	dbEquipment, err := cfg.db.GetMonsterEquipment(r.Context(), mon.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return MonsterEquipment{}, nil
		}
		return MonsterEquipment{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve equipment of %s.", getMonsterName(mon)), err)
	}

	abilitySlots, attachedAbilities, err := cfg.getMonsterEquipmentSlots(r, mon, dbEquipment)
	if err != nil {
		return MonsterEquipment{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve equipment slots of %s.", getMonsterName(mon)), err)
	}

	weaponAbilities, err := cfg.getEquipmentDrops(r, mon, dbEquipment, database.EquipTypeWeapon)
	if err != nil {
		return MonsterEquipment{}, err
	}

	armorAbilities, err := cfg.getEquipmentDrops(r, mon, dbEquipment, database.EquipTypeArmor)
	if err != nil {
		return MonsterEquipment{}, err
	}

	return MonsterEquipment{
		DropChance:        anyToInt32(dbEquipment.DropChance),
		Power:             anyToInt32(dbEquipment.Power),
		CriticalPlus:      dbEquipment.CriticalPlus,
		AbilitySlots:      abilitySlots,
		AttachedAbilities: attachedAbilities,
		WeaponAbilities:   weaponAbilities,
		ArmorAbilities:    armorAbilities,
	}, nil
}

func (cfg *Config) getMonsterEquipmentSlots(r *http.Request, mon database.Monster, equipment database.MonsterEquipment) (MonsterEquipmentSlots, MonsterEquipmentSlots, error) {
	dbEquipmentSlots, err := cfg.db.GetMonsterEquipmentSlots(r.Context(), equipment.ID)
	if err != nil {
		return MonsterEquipmentSlots{}, MonsterEquipmentSlots{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve equipment slots of %s.", getMonsterName(mon)), err)
	}

	dbAbilitySlots := dbEquipmentSlots[0]
	dbAttachedAbilities := dbEquipmentSlots[1]

	abilitySlots, err := cfg.assembleMonsterEquipmentSlots(r, mon, equipment, dbAbilitySlots)
	if err != nil {
		return MonsterEquipmentSlots{}, MonsterEquipmentSlots{}, err
	}

	attachedAbilities, err := cfg.assembleMonsterEquipmentSlots(r, mon, equipment, dbAttachedAbilities)
	if err != nil {
		return MonsterEquipmentSlots{}, MonsterEquipmentSlots{}, err
	}

	return abilitySlots, attachedAbilities, nil
}

func (cfg *Config) assembleMonsterEquipmentSlots(r *http.Request, mon database.Monster, equipment database.MonsterEquipment, dbEquipmentSlots database.MonsterEquipmentSlot) (MonsterEquipmentSlots, error) {
	equipmentSlotsChances, err := cfg.getMonsterEquipmentSlotsChances(r, mon, equipment, dbEquipmentSlots)
	if err != nil {
		return MonsterEquipmentSlots{}, err
	}

	equipmentSlots := MonsterEquipmentSlots{
		MinAmount: anyToInt32(dbEquipmentSlots.MinAmount),
		MaxAmount: anyToInt32(dbEquipmentSlots.MaxAmount),
		Chances:   equipmentSlotsChances,
	}

	return equipmentSlots, nil
}

func (cfg *Config) getMonsterEquipmentSlotsChances(r *http.Request, mon database.Monster, equipment database.MonsterEquipment, slots database.MonsterEquipmentSlot) ([]EquipmentSlotsChance, error) {
	dbSlotsChances, err := cfg.db.GetMonsterEquipmentSlotsChances(r.Context(), database.GetMonsterEquipmentSlotsChancesParams{
		MonsterEquipmentID: equipment.ID,
		EquipmentSlotsID:   slots.ID,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't assemble equipment slots of %s.", getMonsterName(mon)), err)
	}

	chances := []EquipmentSlotsChance{}

	for _, dbChance := range dbSlotsChances {
		chance := EquipmentSlotsChance{
			Amount: anyToInt32(dbChance.Amount),
			Chance: anyToInt32(dbChance.Chance),
		}

		chances = append(chances, chance)
	}

	return chances, nil
}

func (cfg *Config) getEquipmentDrops(r *http.Request, mon database.Monster, equipment database.MonsterEquipment, equipType database.EquipType) ([]EquipmentDrop, error) {
	dbDrops, err := cfg.db.GetMonsterEquipmentAbilities(r.Context(), database.GetMonsterEquipmentAbilitiesParams{
		MonsterEquipmentID: equipment.ID,
		Type:               equipType,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve dropped %s abilities of %s.", string(equipType), getMonsterName(mon)), err)
	}

	drops := []EquipmentDrop{}

	for _, dbDrop := range dbDrops {
		forcedChars, err := cfg.getEquipmentDropForcedChars(r, mon, equipment, dbDrop)
		if err != nil {
			return nil, err
		}
		autoAbility := cfg.newNamedAPIResourceSimple(cfg.e.autoAbilities.endpoint, dbDrop.AutoAbilityID, dbDrop.AutoAbility)

		drop := EquipmentDrop{
			AutoAbility: autoAbility,
			ForcedChars: forcedChars,
			IsForced:    dbDrop.IsForced,
			Probability: anyToInt32Ptr(dbDrop.Probability),
		}

		drops = append(drops, drop)
	}

	return drops, nil
}

func (cfg *Config) getEquipmentDropForcedChars(r *http.Request, mon database.Monster, equipment database.MonsterEquipment, drop database.GetMonsterEquipmentAbilitiesRow) ([]NamedAPIResource, error) {
	dbChars, err := cfg.db.GetEquipmentDropCharacters(r.Context(), database.GetEquipmentDropCharactersParams{
		MonsterEquipmentID: equipment.ID,
		EquipmentDropID:    drop.ID,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve characters of auto-ability '%s' dropped by %s.", drop.AutoAbility, getMonsterName(mon)), err)
	}

	characters := createNamedAPIResourcesSimple(cfg, dbChars, cfg.e.characters.endpoint, func(char database.GetEquipmentDropCharactersRow) (int32, string) {
		return char.CharacterID, char.CharacterName
	})

	return characters, nil
}
