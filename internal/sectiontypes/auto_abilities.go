package sectiontypes

type AutoAbilitySimple struct {
	ID                  int32    `json:"id"`
	URL                 string   `json:"url"`
	Name                string   `json:"name"`
	RequiredItem        *string  `json:"required_item"`
	ActivationCondition string   `json:"activation_condition"`
	Counter             *string  `json:"counter,omitempty"`
	GradualRecovery     *string  `json:"gradual_recovery,omitempty"`
	AutoItemUse         []string `json:"auto_item_use,omitempty"`
	OnHitElement        *string  `json:"on_hit_element,omitempty"`
	AddedElemResist     *string  `json:"added_elem_resist,omitempty"`
	OnHitStatus         *string  `json:"on_hit_status,omitempty"`
	AddedStatusses      []string `json:"added_statusses,omitempty"`
	AddedStatusResists  []string `json:"added_status_resists,omitempty"`
	AddedProperty       *string  `json:"added_property,omitempty"`
	ConversionTo        *string  `json:"conversion_to,omitempty"`
	StatChanges         []string `json:"stat_changes,omitempty"`
	ModifierChanges     []string `json:"modifier_changes,omitempty"`
}

func (a AutoAbilitySimple) GetURL() string {
	return a.URL
}
