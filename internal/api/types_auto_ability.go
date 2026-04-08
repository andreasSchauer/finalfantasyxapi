package api

type AutoAbility struct {
	ID                     int32                             `json:"id"`
	Name                   string                            `json:"name"`
	Category               NamedAPIResource                  `json:"category"`
	Description            *string                           `json:"description"`
	Effect                 string                            `json:"effect"`
	EquipType              *string                           `json:"equip_type"`
	AbilityValue           *int32                            `json:"ability_value"`
	RequiredItem           *ResourceAmount[TypedAPIResource] `json:"required_item"`
	RelatedStats           []NamedAPIResource                `json:"related_stats"`
	LockedOutAutoAbilities []NamedAPIResource                `json:"locked_out_auto_abilities"`
	ActivationCondition    *string                           `json:"activation_condition"`
	Counter                *string                           `json:"counter"`
	GradualRecovery        *NamedAPIResource                 `json:"gradual_recovery"`
	AutoItemUse            []NamedAPIResource                `json:"auto_item_use"`
	OnHitElement           *NamedAPIResource                 `json:"on_hit_element"`
	AddedElemResist        *ElementalResist                  `json:"added_elem_resist"`
	OnHitStatus            *InflictedStatus                  `json:"on_hit_status"`
	AddedStatusses         []NamedAPIResource                `json:"added_statusses"`
	AddedStatusResists     []StatusResist                    `json:"added_status_resists"`
	AddedProperty          *NamedAPIResource                 `json:"added_property"`
	ConversionTo           *NamedAPIResource                 `json:"conversion_to"`
	StatChanges            []StatChange                      `json:"stat_changes"`
	ModifierChanges        []ModifierChange                  `json:"modifier_changes"`
	MonstersDrop 		   []NamedAPIResource 				 `json:"monsters_drop"`
	MonstersItems          []MonItemAmts      				 `json:"monsters_items"`
	ShopsPreAirship  	   []UnnamedAPIResource 			 `json:"shops_pre_airship"`
	ShopsPostAirship 	   []UnnamedAPIResource 			 `json:"shops_post_airship"`
	Treasures			   []UnnamedAPIResource				 `json:"treasures"`
	EquipmentTables        []UnnamedAPIResource              `json:"equipment_tables"`
}