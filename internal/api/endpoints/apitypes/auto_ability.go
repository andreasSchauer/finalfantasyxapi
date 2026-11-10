package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type AutoAbility struct {
	ID                     int32                                          `json:"id"`
	Name                   string                                         `json:"name"`
	Category               string                                         `json:"category"`
	Description            *string                                        `json:"description"`
	Effect                 string                                         `json:"effect"`
	EquipType              string                                         `json:"equip_type"`
	AbilityValue           *int32                                         `json:"ability_value"`
	RequiredItem           *apiresource.ResourceAmount[apiresource.Typed] `json:"required_item"`
	RelatedStats           []apiresource.Named                            `json:"related_stats"`
	LockedOutAutoAbilities []apiresource.Named                            `json:"locked_out_auto_abilities"`
	ActivationCondition    string                                         `json:"activation_condition"`
	Counter                *string                                        `json:"counter"`
	GradualRecovery        *apiresource.Named                             `json:"gradual_recovery"`
	AutoItemUse            []apiresource.Named                            `json:"auto_item_use"`
	OnHitElement           *apiresource.Named                             `json:"on_hit_element"`
	AddedElemResist        *ElementalResist                               `json:"added_elem_resist"`
	OnHitStatus            *InflictedStatus                               `json:"on_hit_status"`
	AddedStatusses         []apiresource.Named                            `json:"added_statusses"`
	AddedStatusResists     []StatusResist                                 `json:"added_status_resists"`
	AddedProperty          *apiresource.Named                             `json:"added_property"`
	ConversionTo           *apiresource.Named                             `json:"conversion_to"`
	StatChanges            []StatChange                                   `json:"stat_changes"`
	ModifierChanges        []ModifierChange                               `json:"modifier_changes"`
	MonstersDrop           []apiresource.Named                            `json:"monsters_drop"`
	MonstersItems          []MonItemAmts                                  `json:"monsters_items"`
	ShopsPreAirship        []apiresource.Unnamed                          `json:"shops_pre_airship"`
	ShopsPostAirship       []apiresource.Unnamed                          `json:"shops_post_airship"`
	Treasures              []apiresource.Unnamed                          `json:"treasures"`
	EquipmentTables        []apiresource.Unnamed                          `json:"equipment_tables"`
}
