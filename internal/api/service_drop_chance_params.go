package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type DropChanceParams struct {
	Monster				int32		`json:"monster"`
	Character			*int32		`json:"character"`
	PartyMembers		[]int32		`json:"party_members"`
	AutoAbilities		[]int32		`json:"auto_abilities"`
	LenientAbilities	bool		`json:"lenient_abilities"`
	MinEmptySlots		*int32		`json:"min_empty_slots"`
	TotalSlots			*int32		`json:"total_slots"`
}

func (p DropChanceParams) GetDoc(cfg *Config) ParamsDoc {
	return cfg.getDropChanceParamsDoc()
}


func (cfg *Config) getDropChanceParamsDoc() ParamsDoc {
	return ParamsDoc{
		Fields: []FieldDoc{
			{
				Field: 			pfnMonster,
				Type: 			"int (id: monster)",
				Required: 		true,
				MinVal:     	h.GetInt32Ptr(1),
				MaxVal:     	h.GetInt32Ptr(int32(len(cfg.l.Monsters))),
				Description: 	"The monster that drops the equipment.",
			},
			{
				Field: 			pfnCharacter,
				Type:			"int (id: character)",
				MinVal:     	h.GetInt32Ptr(1),
				MaxVal:     	h.GetInt32Ptr(int32(len(cfg.l.Characters)-1)),
				Description: 	"Specify a specific character that should get the equipment. Character-specific auto-abilities will be factored in, when calculating the drop chance.",
			},
			{
				Field: 			pfnPartyMembers,
				Type:			"array[int (id: character)]",
				DefaultVal: 	[]int32{1,2,3,4,5,6,7},
				ChildMinVal:   	h.GetInt32Ptr(1),
				ChildMaxVal:    h.GetInt32Ptr(int32(len(cfg.l.Characters)-1)),
				Description: 	"Specify the characters that have permanently joined the party. If a character has not permanently joined the party, they can not get equipment drops. I couldn't find definitive evidence on whether a character's temporary absence (like Yuna in Bikanel/Bevelle; Via Purifico; when underwater) means they cannot get equipment drops, but if you're of that opinion, you can manually leave those characters out.",
			},
			{
				Field: 			pfnAutoAbilities,
				Type:			"array[int (id: auto-ability)]",
				ChildMinVal:   	h.GetInt32Ptr(1),
				ChildMaxVal:    h.GetInt32Ptr(int32(len(cfg.l.AutoAbilities))),
				Description: 	"The auto-abilities the equipment must have. By default, the calculator is looking for exact matches of the given auto-abilities (order doesn't matter), while respecting, if a character can get those auto-abilities. Use an empty array, if it shouldn't have any auto-abilities. Use null, if you don't care about auto-abilities at all. Throws an error, if the monster doesn't drop one of the stated auto-abilities, or if weapon- and armor-abilities are combined.",
			},
			{
				Field: 			pfnLenientAbilities,
				Type: 			"bool",
				Description: 	"If this field is set to true, the calculator treats any equipment that contains the given auto-abilities as a match, no matter what other auto-abilities are on the equipment as well. Example: 'auto_abilities' = [sensor, piercing] => if lenient_abilities is true, an equipment with [sensor, piercing, darkstrike] will be a match, if the field is false, it won't, because it's not an exact match. If auto-abilities is null, this field will have no effect on the calculation.",
			},
			{
				Field: 			pfnMinEmptySlots,
				Type: 			"int",
				MinVal:     	h.GetInt32Ptr(0),
				MaxVal:     	h.GetInt32Ptr(4),
				Description: 	"The minimum amount of empty slots an equipment should have. The calculator will treat equipment with N or more slots as a match. If this field is null, the calculator will treat the difference between total_slots and the amount of auto-abilities as the required empty slot value. If total_slots is also null, the amount of slots has no effect on the calculation. Throws an error, if the amount of auto-abilities and the number of empty slots combined exceed the number of total slots, if given, or 4 of not.",
			},
			{
				Field: 			pfnTotalSlots,
				Type: 			"int",
				MinVal:     	h.GetInt32Ptr(1),
				MaxVal:     	h.GetInt32Ptr(4),
				Description: 	"The amount of total slots the equipment should have. If this field is null, any equipment that matches the auto-ability requirements and has at least N empty slots, will be a match. If empty_slots is also null, the amount of slots has no effect on the calculation. Throws an error, if the given amount of total_slots is lower than the amount of auto-abilities and the number of empty slots combined.",
			},
		},
	}
}