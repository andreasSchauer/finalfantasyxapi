package main

import (
	"fmt"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type MonsterSub struct {
	ID                   	int32                 	`json:"id"`
	URL						string					`json:"url"`
	Name                 	string                	`json:"name"`
	Version              	*int32                	`json:"version,omitempty"`
	Specification        	*string               	`json:"specification,omitempty"`
	HP						int32				  	`json:"hp"`
	OverkillDamage       	int32                 	`json:"overkill_damage"`
	AP                   	int32                 	`json:"ap"`
	APOverkill           	int32                 	`json:"ap_overkill"`
	Gil                  	int32                 	`json:"gil"`
	MaxBribeAmount			*int32					`json:"max_bribe_amount"`
	RonsoRages           	[]string    		  	`json:"ronso_rages"`
	Items					*MonsterItemsSub		`json:"items"`
	Equipment				*MonsterEquipmentSub	`json:"equipment"`
}

func (m MonsterSub) GetSectionName() string {
	return "monsters"
}

// actually need another itemAmount + logic that determines, if an item is a key item or not, even if this won't be relevant for monsters
// can use the master item lookup and its type for this. or simply the same as I did earlier with the itemAmount constructor
type MonsterItemsSub struct {
	StealCommon         	*ItemAmountSub   	`json:"steal_common"`
	StealRare           	*ItemAmountSub   	`json:"steal_rare"`
	DropCommon          	*ItemAmountSub   	`json:"drop_common"`
	DropRare            	*ItemAmountSub   	`json:"drop_rare"`
	SecondaryDropCommon 	*ItemAmountSub   	`json:"secondary_drop_common"`
	SecondaryDropRare   	*ItemAmountSub   	`json:"secondary_drop_rare"`
	Bribe               	*ItemAmountSub   	`json:"bribe"`
	OtherItems				[]ItemAmountSub		`json:"other_items"`
}

type MonsterEquipmentSub struct {
	WeaponAbilities			[]string				`json:"weapon_abilities"`
	ArmorAbilities			[]string				`json:"armor_abilities"`
}




func getSubMonsters(cfg *Config, dbIDs []int32) []SubResource {
	i := cfg.e.monsters
	monsters := []MonsterSub{}

	for _, monID := range dbIDs {
		mon, _ := seeding.GetResourceByID(monID, i.objLookupID)
		monHP := getMonsterSubHP(mon)

		monSub := MonsterSub{
			ID: mon.ID,
			URL: cfg.createResourceURL(i.endpoint, monID),
			Name: mon.Name,
			Version: mon.Version,
			Specification: mon.Specification,
			HP: monHP,
			OverkillDamage: mon.OverkillDamage,
			AP: mon.AP,
			APOverkill: mon.APOverkill,
			Gil: mon.Gil,
			MaxBribeAmount: getMonsterSubBribeAmount(mon, monHP),
			RonsoRages: mon.RonsoRages,
			Items: cfg.getMonsterSubItems(mon),
			Equipment: getMonsterSubEquipment(mon),
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

func (cfg *Config) getMonsterSubItems(mon seeding.Monster) *MonsterItemsSub {
	if mon.Items == nil {
		return nil
	}

	return &MonsterItemsSub{
		StealCommon: cfg.createSubItemAmountPtr(mon.Items.StealCommon),
		StealRare: cfg.createSubItemAmountPtr(mon.Items.StealRare),
		DropCommon: cfg.createSubItemAmountPtr(mon.Items.DropCommon),
		DropRare: cfg.createSubItemAmountPtr(mon.Items.DropRare),
		SecondaryDropCommon: cfg.createSubItemAmountPtr(mon.Items.SecondaryDropCommon),
		SecondaryDropRare: cfg.createSubItemAmountPtr(mon.Items.SecondaryDropRare),
		Bribe: cfg.createSubItemAmountPtr(mon.Items.Bribe),
		OtherItems: cfg.getMonsterSubOtherItems(mon.Items.OtherItems),
	}
}


func (cfg *Config) getMonsterSubOtherItems(items []seeding.PossibleItem) []ItemAmountSub {
	otherItems := []ItemAmountSub{}

	for _, item := range items {
		otherItem := cfg.createSubItemAmount(item.ItemAmount)
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
		ArmorAbilities: getMonsterSubAutoAbilities(mon.Equipment.ArmorAbilities),
	}
}

func getMonsterSubAutoAbilities(drops []seeding.EquipmentDrop) []string {
	autoAbilities := []string{}

	for _, drop := range drops {
		ability := formatAutoAbility(drop)
		autoAbilities = append(autoAbilities, ability)
	}

	return autoAbilities
}

func formatAutoAbility(drop seeding.EquipmentDrop) string {
	if len(drop.Characters) == 0 {
		return drop.Ability
	}

	formattedChars := strings.Join(drop.Characters, ", ")
	return fmt.Sprintf("%s (%s)", drop.Ability, formattedChars)
}