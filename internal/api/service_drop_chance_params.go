package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type DropChanceParams struct {
	Monster           int32   `json:"monster_id,omitempty"`
	CharacterID       *int32  `json:"character_id,omitempty"`
	PartyMembers      []int32 `json:"party_members,omitempty"`
	AutoAbilities     []int32 `json:"auto_abilities,omitempty"`
	OptionalAbilities bool    `json:"optional_abilities,omitempty"`
	EquipType         *string `json:"equip_type,omitempty"`
	MinEmptySlots     *int32  `json:"min_empty_slots,omitempty"`
	TotalSlots        *int32  `json:"total_slots,omitempty"`
	Decimals          bool    `json:"decimals,omitempty"`
}

func (p DropChanceParams) GetRequiredSlots() int32 {
	return h.Len32(p.AutoAbilities) + h.DerefInt32Ptr(p.MinEmptySlots)
}

func getDropChanceParamsDoc(cfg *Config) ParamsDoc {
	return ParamsDoc{
		Fields: []FieldDoc{
			{
				Field:       PfnMonsterID,
				Type:        "int (id: monster)",
				Required:    true,
				MinVal:      h.GetInt32Ptr(1),
				MaxVal:      h.GetInt32Ptr(int32(len(cfg.L.Monsters))),
				Description: "The monster that drops the equipment.",
			},
			{
				Field:       PfnCharacterID,
				Type:        "int (id: character)",
				MinVal:      h.GetInt32Ptr(1),
				MaxVal:      h.GetInt32Ptr(int32(len(cfg.L.Characters) - 1)),
				Description: "Specify a character to calculate the chance of them receiving equipment. Character-specific auto-abilities like 'Piercing' for Kimahri and Auron will be factored in, when calculating the drop chance.",
			},
			{
				Field:       PfnPartyMembers,
				Type:        "array[int (id: character)]",
				DefaultVal:  []int32{1, 2, 3, 4, 5, 6, 7},
				ChildMinVal: h.GetInt32Ptr(1),
				ChildMaxVal: h.GetInt32Ptr(int32(len(cfg.L.Characters) - 1)),
				MaxArrayLen: h.GetIntPtr(7),
				Description: "Specify the characters that are currently accessible in the menu. Only these characters can receive equipment. For example, in the Gagazet underwater section, only Tidus, Wakka and Rikku can get equipment. If you fight an underwater enemy in the Monster Arena however, everybody is eligible, even though you can only fight with these three.",
			},
			{
				Field:       PfnAutoAbilities,
				Type:        "array[int (id: auto-ability)]",
				ChildMinVal: h.GetInt32Ptr(1),
				ChildMaxVal: h.GetInt32Ptr(int32(len(cfg.L.AutoAbilities))),
				MaxArrayLen: h.GetIntPtr(4),
				RequiredOr:  []FieldName{PfnMinEmptySlots, PfnTotalSlots},
				Description: "The auto-abilities the equipment must have. By default, the calculator is looking for exact matches of the given auto-abilities (order doesn't matter), while respecting, if a character can get those auto-abilities. Use null, if you don't care about auto-abilities at all. Throws an error, if the monster doesn't drop one of the stated auto-abilities, or if weapon- and armor-abilities are combined.",
			},
			{
				Field:       PfnOptionalAbilities,
				Type:        "bool",
				Description: "A searched equipment consists of required slots (amount of wanted auto-abilities + minEmptySlots) and optional slots (total slots - required slots). By default, the calculator checks, if all required slots are present, and if all optional slots are empty. This means, only equipment with only the specified abilities and at least the given amount of empty slots will be a match. If this field is set to true, the calculator will also treat equipment with unspecified abilities on optional slots as a match. Example: 'auto_abilities' = [sensor, piercing] => if optional_abilities is true, an equipment with [sensor, piercing, darkstrike] will be a match; if the field is false, it won't, because 'darkstrike' wasn't specified. If 'auto_abilities' is null: If this field is true, any ability + empty slot combination will be a match. If it's false, only equipment with all empty slots will match.",
			},
			{
				Field:       PfnEquipType,
				Type:        "string (enum: equipType)",
				EnumValues:  cfg.T.EquipType.Strings(),
				Description: "The type of the desired equipment. Will be assigned automatically by looking at the first auto-ability's type, if left empty. If a value is given, all given auto-abilities need to be of that type. If 'auto_abilities' is null, this field is required to discern the equipment's type.",
			},
			{
				Field:       PfnMinEmptySlots,
				Type:        "int",
				MinVal:      h.GetInt32Ptr(0),
				MaxVal:      h.GetInt32Ptr(4),
				RequiredOr:  []FieldName{PfnAutoAbilities, PfnTotalSlots},
				Description: "The minimum amount of empty slots an equipment should have. The calculator will treat equipment with N or more empty slots as a match, since more empty slots are objectively better and thus more desirable. If this field is null, the calculator will treat the difference between total_slots and the amount of auto-abilities as the required empty slot value. If total_slots is also null, the amount of slots has no effect on the calculation. Throws an error, if the amount of auto-abilities and the number of empty slots combined exceed the number of total slots, if given, or 4 if not.",
			},
			{
				Field:       PfnTotalSlots,
				Type:        "int",
				MinVal:      h.GetInt32Ptr(0),
				MaxVal:      h.GetInt32Ptr(4),
				RequiredOr:  []FieldName{PfnAutoAbilities, PfnMinEmptySlots},
				Description: "The amount of total slots the equipment should have. If this field is null, any equipment that matches the auto-ability requirements and has at least N empty slots, will be a match. If empty_slots is also null, the amount of slots has no effect on the calculation. Throws an error, if the given amount of total_slots is lower than the amount of auto-abilities and the number of empty slots combined.",
			},
			{
				Field:       PfnDecimals,
				Type:        "bool",
				Description: "If this field is set to true, the chances in the response are returned as decimals instead of percentages.",
			},
		},
	}
}
