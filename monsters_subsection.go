package main

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

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
	MaxBribeAmount			int32					`json:"max_bribe_amount"`
	RonsoRages           	[]string    		  	`json:"ronso_rages"`
	Items					*MonsterItemsSub		`json:"items"`
	WeaponAbilities			[]string				`json:"weapon_abilities"`
	ArmorAbilities			[]string				`json:"armor_abilities"`
}

type MonsterItemsSub struct {
	StealCommon         	*seeding.ItemAmount   	`json:"steal_common"`
	StealRare           	*seeding.ItemAmount   	`json:"steal_rare"`
	DropCommon          	*seeding.ItemAmount   	`json:"drop_common"`
	DropRare            	*seeding.ItemAmount   	`json:"drop_rare"`
	SecondaryDropCommon 	*seeding.ItemAmount   	`json:"secondary_drop_common"`
	SecondaryDropRare   	*seeding.ItemAmount   	`json:"secondary_drop_rare"`
	Bribe               	*seeding.ItemAmount   	`json:"bribe"`
}