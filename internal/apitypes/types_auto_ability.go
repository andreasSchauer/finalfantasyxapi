package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type AutoAbility struct {
	ID                     int32                                                     `json:"id"`
	Name                   string                                                    `json:"name"`
	Category               string                                                    `json:"category"`
	Description            *string                                                   `json:"description"`
	Effect                 string                                                    `json:"effect"`
	EquipType              string                                                    `json:"equip_type"`
	AbilityValue           *int32                                                    `json:"ability_value"`
	RequiredItem           *apiresource.ResourceAmount[apiresource.TypedAPIResource] `json:"required_item"`
	RelatedStats           []apiresource.NamedAPIResource                            `json:"related_stats"`
	LockedOutAutoAbilities []apiresource.NamedAPIResource                            `json:"locked_out_auto_abilities"`
	ActivationCondition    string                                                    `json:"activation_condition"`
	Counter                *string                                                   `json:"counter"`
	GradualRecovery        *apiresource.NamedAPIResource                             `json:"gradual_recovery"`
	AutoItemUse            []apiresource.NamedAPIResource                            `json:"auto_item_use"`
	OnHitElement           *apiresource.NamedAPIResource                             `json:"on_hit_element"`
	AddedElemResist        *ElementalResist                                          `json:"added_elem_resist"`
	OnHitStatus            *InflictedStatus                                          `json:"on_hit_status"`
	AddedStatusses         []apiresource.NamedAPIResource                            `json:"added_statusses"`
	AddedStatusResists     []StatusResist                                            `json:"added_status_resists"`
	AddedProperty          *apiresource.NamedAPIResource                             `json:"added_property"`
	ConversionTo           *apiresource.NamedAPIResource                             `json:"conversion_to"`
	StatChanges            []StatChange                                              `json:"stat_changes"`
	ModifierChanges        []ModifierChange                                          `json:"modifier_changes"`
	MonstersDrop           []apiresource.NamedAPIResource                            `json:"monsters_drop"`
	MonstersItems          []MonItemAmts                                             `json:"monsters_items"`
	ShopsPreAirship        []apiresource.UnnamedAPIResource                          `json:"shops_pre_airship"`
	ShopsPostAirship       []apiresource.UnnamedAPIResource                          `json:"shops_post_airship"`
	Treasures              []apiresource.UnnamedAPIResource                          `json:"treasures"`
	EquipmentTables        []apiresource.UnnamedAPIResource                          `json:"equipment_tables"`
}
