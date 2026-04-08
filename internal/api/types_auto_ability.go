package api



type AutoAbility struct {
	ID						int32								`json:"id"`
	Name               	 	string           					`json:"name"`
	Description         	*string          					`json:"description"`
	Effect              	string           					`json:"effect"`
	EquipType				string								`json:"equip_type"`
	Category				NamedAPIResource					`json:"category"`
	RelatedStats			[]NamedAPIResource					`json:"related_stats"`
	AbilityValue			*int32								`json:"ability_value"`
	RequiredItem			*ResourceAmount[NamedAPIResource] 	`json:"required_item"`
	LockedOutAutoAbilities	[]NamedAPIResource					`json:"locked_out_auto_abilities"`
	ActivationCondition		*string								`json:"activation_condition"`
	Counter					*string								`json:"counter"`
	GradualRecovery 		*NamedAPIResource					`json:"gradual_recovery"`
	AutoItemUse				[]NamedAPIResource 					`json:"auto_item_use"`
	OnHitElement			*NamedAPIResource 					`json:"on_hit_element"`
	AddedElemResist			*ElementalResist					`json:"added_elem_resist"`
	OnHitStatus				*InflictedStatus 					`json:"on_hit_status"`
	AddedStatusses			[]NamedAPIResource					`json:"added_statusses"`
	AddedProperty			*NamedAPIResource					`json:"added_property"`
	ConversionTo			*NamedAPIResource					`json:"conversion_to"`
	StatChanges				[]StatChange 						`json:"stat_changes"`
	ModifierChanges			[]ModifierChange 					`json:"modifier_changes"`
	MonsterDrops			*AutoAbilityMonsterDrop 			`json:"monster_drops"`
	Shops					*AutoAbilityShop					`json:"shops"`
	EquipmentTables			[]UnnamedAPIResource 				`json:"equipment_tables"`
}

type AutoAbilityMonsterDrop struct {
	AutoAbility	[]NamedAPIResource	`json:"auto_ability"`
	Item		[]MonItemAmts		`json:"item"`
}

type AutoAbilityShop struct {
	PreAirship	[]UnnamedAPIResource	`json:"pre_airship"`
	PostAirship	[]UnnamedAPIResource	`json:"post_airship"`
}