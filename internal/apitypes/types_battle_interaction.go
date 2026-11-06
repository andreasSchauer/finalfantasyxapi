package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type BattleInteraction struct {
	Target                    string                         `json:"target"`
	Range                     *int32                         `json:"range"`
	HitAmount                 int32                          `json:"hit_amount"`
	ShatterRate               int32                          `json:"shatter_rate"`
	SpecialAction             *string                        `json:"special_action,omitempty"`
	BasedOnUserAttack         bool                           `json:"based_on_user_attack"`
	Darkable                  bool                           `json:"darkable"`
	Silenceable               bool                           `json:"silenceable"`
	Reflectable               bool                           `json:"reflectable"`
	Accuracy                  Accuracy                       `json:"accuracy"`
	Damage                    *Damage                        `json:"damage"`
	InflictedDelay            *InflictedDelay                `json:"inflicted_delay"`
	InflictedStatusConditions []InflictedStatus              `json:"inflicted_status_conditions"`
	RemovedStatusConditions   []apiresource.NamedAPIResource `json:"removed_status_conditions"`
	CopiedStatusConditions    []InflictedStatus              `json:"copied_status_conditions"`
	StatChanges               []StatChange                   `json:"stat_changes"`
	ModifierChanges           []ModifierChange               `json:"modifier_changes"`
}

type Damage struct {
	AttackType      string                        `json:"attack_type"`
	DamageType      string                        `json:"damage_type"`
	HP              *ClassDamage                  `json:"hp"`
	MP              *ClassDamage                  `json:"mp"`
	CTB             *ClassDamage                  `json:"ctb"`
	Critical        *string                       `json:"critical"`
	CriticalPlusVal *int32                        `json:"critical_plus_val,omitempty"`
	IsPiercing      bool                          `json:"is_piercing"`
	BreakDmgLimit   *string                       `json:"break_dmg_lmt"`
	Element         *apiresource.NamedAPIResource `json:"element"`
}

type ClassDamage struct {
	Condition      *string `json:"condition,omitempty"`
	DamageFormula  string  `json:"damage_formula"`
	DamageConstant int32   `json:"damage_constant"`
}

type Accuracy struct {
	AccSource   string   `json:"acc_source"`
	HitChance   *int32   `json:"hit_chance,omitempty"`
	AccModifier *float32 `json:"acc_modifier,omitempty"`
}

type InflictedDelay struct {
	Condition     *string `json:"condition,omitempty"`
	DelayStrength string  `json:"delay_strength"`
}

type InflictedStatus struct {
	StatusCondition apiresource.NamedAPIResource `json:"status_condition"`
	Probability     int32                        `json:"probability"`
	DurationType    string                       `json:"duration_type"`
	Amount          *int32                       `json:"amount,omitempty"`
}

func (is InflictedStatus) GetAPIResource() apiresource.APIResource {
	return is.StatusCondition
}

func (is InflictedStatus) IsZero() bool {
	return is.StatusCondition.Name == ""
}

type StatChange struct {
	Stat            apiresource.NamedAPIResource `json:"stat"`
	CalculationType string                       `json:"calculation_type"`
	Value           float32                      `json:"value"`
}

func (sc StatChange) GetAPIResource() apiresource.APIResource {
	return sc.Stat
}

func (sc StatChange) IsZero() bool {
	return sc.Stat.Name == ""
}

type ModifierChange struct {
	Modifier        apiresource.NamedAPIResource `json:"modifier"`
	CalculationType string                       `json:"calculation_type"`
	Value           float32                      `json:"value"`
}

func (mc ModifierChange) GetAPIResource() apiresource.APIResource {
	return mc.Modifier
}

func (mc ModifierChange) IsZero() bool {
	return mc.Modifier.Name == ""
}
