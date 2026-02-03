package main

import (
	"fmt"
	"net/http"

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
	IsStoryBased   bool                 `json:"is_story_based"`
	IsRepeatable   bool                 `json:"is_repeatable"`
	CanBeCaptured  bool                 `json:"can_be_captured"`
	RonsoRages     []string             `json:"ronso_rages"`
	Items          *MonsterItemsSub     `json:"items"`
	Equipment      *MonsterEquipmentSub `json:"equipment"`
}

func (m MonsterSub) GetSectionName() string {
	return "monsters"
}

func (m MonsterSub) GetURL() string {
	return m.URL
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

func convertMonsterSubItems(cfg *Config, items seeding.MonsterItems) MonsterItemsSub {
	return MonsterItemsSub{
		StealCommon:         convertObjPtr(cfg, items.StealCommon, convertSubItemAmount),
		StealRare:           convertObjPtr(cfg, items.StealRare, convertSubItemAmount),
		DropCommon:          convertObjPtr(cfg, items.DropCommon, convertSubItemAmount),
		DropRare:            convertObjPtr(cfg, items.DropRare, convertSubItemAmount),
		SecondaryDropCommon: convertObjPtr(cfg, items.SecondaryDropCommon, convertSubItemAmount),
		SecondaryDropRare:   convertObjPtr(cfg, items.SecondaryDropRare, convertSubItemAmount),
		Bribe:               convertObjPtr(cfg, items.Bribe, convertSubItemAmount),
		OtherItems:          convertObjSlice(cfg, items.OtherItems, posItemToItemAmtSub),
	}
}

type MonsterEquipmentSub struct {
	WeaponAbilities []string `json:"weapon_abilities"`
	ArmorAbilities  []string `json:"armor_abilities"`
}

func convertMonsterSubEquipment(_ *Config, equipment seeding.MonsterEquipment) MonsterEquipmentSub {
	return MonsterEquipmentSub{
		WeaponAbilities: getMonsterSubAutoAbilities(equipment.WeaponAbilities),
		ArmorAbilities:  getMonsterSubAutoAbilities(equipment.ArmorAbilities),
	}
}

func handleMonstersSection(cfg *Config, r *http.Request, dbIDs []int32) ([]SubResource, error) {
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
			IsStoryBased:   mon.IsStoryBased,
			IsRepeatable:   mon.IsRepeatable,
			CanBeCaptured:  mon.CanBeCaptured,
			RonsoRages:     mon.RonsoRages,
			Items:          convertObjPtr(cfg, mon.Items, convertMonsterSubItems),
			Equipment:      convertObjPtr(cfg, mon.Equipment, convertMonsterSubEquipment),
		}

		monsters = append(monsters, monSub)
	}

	return toSubResourceSlice(monsters), nil
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
