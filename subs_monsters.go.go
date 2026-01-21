package main

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type MonsterSub struct {
	ID             int32                `json:"id"`
	URL            string               `json:"url"`
	Name           string               `json:"name"`
	Version        *int32               `json:"version,omitempty"`
	Specification  *string              `json:"specification,omitempty"`
	HP             int32                `json:"hp"`
	OverkillDamage int32                `json:"overkill_damage"`
	AP             int32                `json:"ap"`
	APOverkill     int32                `json:"ap_overkill"`
	Gil            int32                `json:"gil"`
	MaxBribeAmount *int32               `json:"max_bribe_amount"`
	RonsoRages     []string             `json:"ronso_rages"`
	Items          *MonsterItemsSub     `json:"items"`
	Equipment      *MonsterEquipmentSub `json:"equipment"`
}

func (m MonsterSub) GetSectionName() string {
	return "monsters"
}

type MonsterItemsSub struct {
	StealCommon         *ItemAmountSub  `json:"steal_common"`
	StealRare           *ItemAmountSub  `json:"steal_rare"`
	DropCommon          *ItemAmountSub  `json:"drop_common"`
	DropRare            *ItemAmountSub  `json:"drop_rare"`
	SecondaryDropCommon *ItemAmountSub  `json:"secondary_drop_common"`
	SecondaryDropRare   *ItemAmountSub  `json:"secondary_drop_rare"`
	Bribe               *ItemAmountSub  `json:"bribe"`
	OtherItems          []ItemAmountSub `json:"other_items"`
}

type MonsterEquipmentSub struct {
	WeaponAbilities []string `json:"weapon_abilities"`
	ArmorAbilities  []string `json:"armor_abilities"`
}

func getSubMonsters(cfg *Config, dbIDs []int32) []SubResource {
	i := cfg.e.monsters
	monsters := []MonsterSub{}

	for _, monID := range dbIDs {
		mon, _ := seeding.GetResourceByID(monID, i.objLookupID)
		monHP := getMonsterSubHP(mon)

		monSub := MonsterSub{
			ID:             mon.ID,
			URL:            createResourceURL(cfg, i.endpoint, monID),
			Name:           mon.Name,
			Version:        mon.Version,
			Specification:  mon.Specification,
			HP:             monHP,
			OverkillDamage: mon.OverkillDamage,
			AP:             mon.AP,
			APOverkill:     mon.APOverkill,
			Gil:            mon.Gil,
			MaxBribeAmount: getMonsterSubBribeAmount(mon, monHP),
			RonsoRages:     mon.RonsoRages,
			Items:          getMonsterSubItems(cfg, mon),
			Equipment:      getMonsterSubEquipment(mon),
		}

		monsters = append(monsters, monSub)
	}

	return toSubResourceSlice(monsters)
}

func getMonsterSubHP(mon seeding.Monster) int32 {
	for _, stat := range mon.BaseStats {
		if stat.StatName == "hp" {
			return stat.Value
		}
	}

	return 0
}

func getMonsterSubBribeAmount(mon seeding.Monster, hp int32) *int32 {
	if mon.Items == nil || mon.Items.Bribe == nil {
		return nil
	}

	bribeAmount := hp * 25
	return &bribeAmount
}

func getMonsterSubItems(cfg *Config, mon seeding.Monster) *MonsterItemsSub {
	if mon.Items == nil {
		return nil
	}

	return &MonsterItemsSub{
		StealCommon:         createSubItemAmountPtr(cfg, mon.Items.StealCommon),
		StealRare:           createSubItemAmountPtr(cfg, mon.Items.StealRare),
		DropCommon:          createSubItemAmountPtr(cfg, mon.Items.DropCommon),
		DropRare:            createSubItemAmountPtr(cfg, mon.Items.DropRare),
		SecondaryDropCommon: createSubItemAmountPtr(cfg, mon.Items.SecondaryDropCommon),
		SecondaryDropRare:   createSubItemAmountPtr(cfg, mon.Items.SecondaryDropRare),
		Bribe:               createSubItemAmountPtr(cfg, mon.Items.Bribe),
		OtherItems:          getMonsterSubOtherItems(cfg, mon.Items.OtherItems),
	}
}

func getMonsterSubOtherItems(cfg *Config, items []seeding.PossibleItem) []ItemAmountSub {
	otherItems := []ItemAmountSub{}

	for _, item := range items {
		otherItem := createSubItemAmount(cfg, item.ItemAmount)
		otherItems = append(otherItems, otherItem)
	}

	return otherItems
}

func getMonsterSubEquipment(mon seeding.Monster) *MonsterEquipmentSub {
	if mon.Equipment == nil {
		return nil
	}

	return &MonsterEquipmentSub{
		WeaponAbilities: getMonsterSubAutoAbilities(mon.Equipment.WeaponAbilities),
		ArmorAbilities:  getMonsterSubAutoAbilities(mon.Equipment.ArmorAbilities),
	}
}

func getMonsterSubAutoAbilities(drops []seeding.EquipmentDrop) []string {
	autoAbilities := []string{}

	for _, drop := range drops {
		autoAbility := formatMonsterAutoAbility(drop)
		autoAbilities = append(autoAbilities, autoAbility)
	}

	return autoAbilities
}

func formatMonsterAutoAbility(drop seeding.EquipmentDrop) string {
	if len(drop.Characters) == 0 {
		return drop.Ability
	}

	formattedChars := h.StringSliceToListString(drop.Characters)
	return fmt.Sprintf("%s (%s)", drop.Ability, formattedChars)
}
