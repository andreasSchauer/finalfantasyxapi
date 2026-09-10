package api

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type DropChanceResponse struct {
	URL            		string 					`json:"url"`
	TotalDropChances	TotalDropChances		`json:"total_drop_chances"`
	IsolatedDropChances IsolatedDropChances		`json:"isolated_drop_chances"`
}

func (r DropChanceResponse) GetURL() string {
	return r.URL
}

func (r DropChanceResponse) Round(digits int32) DropChanceResponse {
	r.TotalDropChances = r.TotalDropChances.Round(digits)
	r.IsolatedDropChances = r.IsolatedDropChances.Round(digits)

	return r
}

type TotalDropChances struct {
	WithFinBlow			float64		`json:"with_fin_blow"`
	NoFinBlow			float64		`json:"no_fin_blow"`
}

func (c TotalDropChances) Round(digits int32) TotalDropChances {
	c.WithFinBlow = h.FloatRound(c.WithFinBlow, digits)
	c.NoFinBlow = h.FloatRound(c.NoFinBlow, digits)

	return c
}

type IsolatedDropChances struct {
	EquipmentDrop		float64		`json:"equipment_drop"`
	EquipmentMatch		float64		`json:"equipment_match"`
	CharChances
}

func (c IsolatedDropChances) Round(digits int32) IsolatedDropChances {
	c.EquipmentDrop = h.FloatRound(c.EquipmentDrop, digits)
	c.EquipmentMatch = h.FloatRound(c.EquipmentMatch, digits)
	c.CharChances = c.CharChances.Round(digits)

	return c
}

type CharChances struct {
	AnyCharFinBlow		float64		`json:"any_char_fin_blow"`
	AnyCharNoFinBlow	float64		`json:"any_char_no_fin_blow"`
	CharFinBlow			*float64	`json:"char_fin_blow,omitempty"`
	CharNoFinBlow		*float64	`json:"char_no_fin_blow,omitempty"`
}

func (c CharChances) Round(digits int32) CharChances {
	c.AnyCharFinBlow = h.FloatRound(c.AnyCharFinBlow, digits)
	c.AnyCharNoFinBlow = h.FloatRound(c.AnyCharNoFinBlow, digits)
	c.CharFinBlow = h.FloatPtrRound(c.CharFinBlow, digits)
	c.CharNoFinBlow = h.FloatPtrRound(c.CharNoFinBlow, digits)

	return c
}

type AbilityWheel struct {
	PrioritySlot		string
	Abilities			[7]string
}

func calcDropChance(cfg *Config, params DropChanceParams, url string) (DropChanceResponse, error) {
	response := DropChanceResponse{
		URL: url,
	}

	mon, _ := seeding.GetResourceByID(params.Monster, cfg.l.MonstersID)
	charPtr := getCharPtr(cfg, params)

	wantedAbilities, monAbilities, err := vfAutoAbilities(cfg, params, charPtr, mon)
	if err != nil {
		return DropChanceResponse{}, err
	}

	_ = extractAbilitySlots(params, mon)
	_ = mon.Equipment.AttachedAbilities.Chances
	charChances := calcCharRandomChances(cfg, params, wantedAbilities, monAbilities)
	monDropChance := calcMonDropChance(mon)
	
	response.IsolatedDropChances = IsolatedDropChances{
		EquipmentDrop: monDropChance,
		CharChances: charChances,
	}

	return response.Round(4), nil
}

func extractAbilitySlots(params DropChanceParams, mon seeding.Monster) ([]seeding.EquipmentSlotsChance) {
	var minSlots int32 = getRequiredSlots(params)
	var maxSlots int32 = 4
	abilitySlotChances := mon.Equipment.AbilitySlots.Chances
	var chances []seeding.EquipmentSlotsChance

	if params.TotalSlots != nil {
		minSlots = *params.TotalSlots
		maxSlots = *params.TotalSlots
	}

	for _, chance := range abilitySlotChances {
		if chance.Amount >= minSlots && chance.Amount <= maxSlots {
			chances = append(chances, chance)
		}
	}

	return chances
}



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
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s doesn't drop '%s'", mon, ability), nil)
		}
	}

	return nil
}

func getCharPtr(cfg *Config, params DropChanceParams) *seeding.Character {
	if params.Character == nil {
		return nil
	}

	char, _ := seeding.GetResourceByID(*params.Character, cfg.l.CharactersID)
	return &char
}

func vfCharGetsAbility(ability seeding.EquipmentDrop, charPtr *seeding.Character, mon seeding.Monster) error {
	if charPtr == nil {
		return nil
	}
	
	if !charGetsAbility(ability, charPtr) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s doesn't drop '%s' for %s", mon, ability.Ability, charPtr.Name), nil)
	}

	return nil
}

func charGetsAbility(ability seeding.EquipmentDrop, charPtr *seeding.Character) bool {
	if charPtr == nil {
		return true
	}
	
	chars := ability.Characters

	if len(chars) == 0 {
		return true
	}

	return slices.Contains(chars, charPtr.Name)
}

func calcMonDropChance(mon seeding.Monster) float64 {
	return float64(mon.Equipment.DropChance) / 255
}


func calcCharRandomChances(cfg *Config, params DropChanceParams, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop) CharChances {
	var eligibleChars int32
	baseRandomChance := 0.75 / h.FloatLen(params.PartyMembers)
	var chances CharChances

	if params.Character != nil {
		char, _ := seeding.GetResourceByID(*params.Character, cfg.l.CharactersID)

		if charLearnsReqAbilities(char, wantedAbilities, monAbilities) {
			finBlowChance := baseRandomChance + 0.25
			chances.CharNoFinBlow = &baseRandomChance
			chances.CharFinBlow = &finBlowChance
		}
	}

	for _, charID := range params.PartyMembers {
		char, _ := seeding.GetResourceByID(charID, cfg.l.CharactersID)

		if charLearnsReqAbilities(char, wantedAbilities, monAbilities) {
			eligibleChars++
		}
	}

	chances.AnyCharNoFinBlow = float64(eligibleChars) * baseRandomChance

	if eligibleChars > 0 {
		chances.AnyCharFinBlow = chances.AnyCharNoFinBlow + 0.25
	}

	return chances
}


func charLearnsReqAbilities(char seeding.Character, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop) bool {
	for _, ability := range wantedAbilities {
		for _, monAbility := range monAbilities {
			if ability.Name == monAbility.Ability && !charGetsAbility(monAbility, &char) {
				return false
			}
		}
	}

	return true
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