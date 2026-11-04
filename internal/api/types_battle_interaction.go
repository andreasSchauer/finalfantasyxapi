package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type BattleInteraction struct {
	Target                    string             `json:"target"`
	Range                     *int32             `json:"range"`
	HitAmount                 int32              `json:"hit_amount"`
	ShatterRate               int32              `json:"shatter_rate"`
	SpecialAction             *string            `json:"special_action,omitempty"`
	BasedOnUserAttack         bool               `json:"based_on_user_attack"`
	Darkable                  bool               `json:"darkable"`
	Silenceable               bool               `json:"silenceable"`
	Reflectable               bool               `json:"reflectable"`
	Accuracy                  Accuracy           `json:"accuracy"`
	Damage                    *Damage            `json:"damage"`
	InflictedDelay            *InflictedDelay    `json:"inflicted_delay"`
	InflictedStatusConditions []InflictedStatus  `json:"inflicted_status_conditions"`
	RemovedStatusConditions   []NamedAPIResource `json:"removed_status_conditions"`
	CopiedStatusConditions    []InflictedStatus  `json:"copied_status_conditions"`
	StatChanges               []StatChange       `json:"stat_changes"`
	ModifierChanges           []ModifierChange   `json:"modifier_changes"`
}

func convertBattleInteraction(srv *Server, bi seedtypes.BattleInteraction) BattleInteraction {
	battleInteraction := BattleInteraction{
		Target:                    bi.Target,
		BasedOnUserAttack:         bi.BasedOnUserAttack,
		Range:                     bi.Range,
		Damage:                    convertObjPtr(srv, bi.Damage, convertDamage),
		ShatterRate:               bi.ShatterRate,
		Accuracy:                  convertAccuracy(srv, bi.Accuracy),
		HitAmount:                 bi.HitAmount,
		SpecialAction:             bi.SpecialAction,
		InflictedDelay:            convertObjPtr(srv, bi.InflictedDelay, convertInflictedDelay),
		InflictedStatusConditions: convertObjSlice(srv, bi.InflictedStatusConditions, convertInflictedStatus),
		RemovedStatusConditions:   namesToNamedAPIResources(srv, E.StatusConditions, bi.RemovedStatusConditions),
		CopiedStatusConditions:    convertObjSlice(srv, bi.CopiedStatusConditions, convertInflictedStatus),
		StatChanges:               convertObjSlice(srv, bi.StatChanges, convertStatChange),
		ModifierChanges:           convertObjSlice(srv, bi.ModifierChanges, convertModifierChange),
	}

	for _, status := range bi.AffectedBy {
		switch status {
		case "darkness":
			battleInteraction.Darkable = true

		case "silence":
			battleInteraction.Silenceable = true

		case "reflect":
			battleInteraction.Reflectable = true
		}
	}

	return battleInteraction
}

type Damage struct {
	AttackType      string            `json:"attack_type"`
	DamageType      string            `json:"damage_type"`
	HP              *ClassDamage      `json:"hp"`
	MP              *ClassDamage      `json:"mp"`
	CTB             *ClassDamage      `json:"ctb"`
	Critical        *string           `json:"critical"`
	CriticalPlusVal *int32            `json:"critical_plus_val,omitempty"`
	IsPiercing      bool              `json:"is_piercing"`
	BreakDmgLimit   *string           `json:"break_dmg_lmt"`
	Element         *NamedAPIResource `json:"element"`
}

func convertDamage(srv *Server, d seedtypes.Damage) Damage {
	return Damage{
		AttackType:      d.AttackType,
		DamageType:      d.DamageType,
		HP:              convertObjPtr(srv, d.HP, convertClassDamage),
		MP:              convertObjPtr(srv, d.MP, convertClassDamage),
		CTB:             convertObjPtr(srv, d.CTB, convertClassDamage),
		Critical:        d.Critical,
		CriticalPlusVal: d.CriticalPlusVal,
		IsPiercing:      d.IsPiercing,
		BreakDmgLimit:   d.BreakDmgLimit,
		Element:         namePtrToNamedAPIResPtr(srv, E.Elements, d.Element, nil),
	}
}

type ClassDamage struct {
	Condition      *string `json:"condition,omitempty"`
	DamageFormula  string  `json:"damage_formula"`
	DamageConstant int32   `json:"damage_constant"`
}

func convertClassDamage(srv *Server, ad seedtypes.ClassDamage) ClassDamage {
	return ClassDamage{
		Condition:      ad.Condition,
		DamageFormula:  ad.DamageFormula,
		DamageConstant: ad.DamageConstant,
	}
}

type Accuracy struct {
	AccSource   string   `json:"acc_source"`
	HitChance   *int32   `json:"hit_chance,omitempty"`
	AccModifier *float32 `json:"acc_modifier,omitempty"`
}

func convertAccuracy(_ *Server, a seedtypes.Accuracy) Accuracy {
	return Accuracy{
		AccSource:   a.AccSource,
		HitChance:   a.HitChance,
		AccModifier: a.AccModifier,
	}
}

type InflictedDelay struct {
	Condition     *string `json:"condition,omitempty"`
	DelayStrength string  `json:"delay_strength"`
}

func convertInflictedDelay(_ *Server, id seedtypes.InflictedDelay) InflictedDelay {
	return InflictedDelay{
		Condition:     id.Condition,
		DelayStrength: id.DelayStrength,
	}
}

type InflictedStatus struct {
	StatusCondition NamedAPIResource `json:"status_condition"`
	Probability     int32            `json:"probability"`
	DurationType    string           `json:"duration_type"`
	Amount          *int32           `json:"amount,omitempty"`
}

func (is InflictedStatus) GetAPIResource() APIResource {
	return is.StatusCondition
}

func (is InflictedStatus) IsZero() bool {
	return is.StatusCondition.Name == ""
}

func convertInflictedStatus(srv *Server, status seedtypes.InflictedStatus) InflictedStatus {
	return InflictedStatus{
		StatusCondition: nameToNamedAPIResource(srv, E.StatusConditions, status.StatusCondition, nil),
		Probability:     status.Probability,
		DurationType:    status.DurationType,
		Amount:          status.Amount,
	}
}

type StatChange struct {
	Stat            NamedAPIResource `json:"stat"`
	CalculationType string           `json:"calculation_type"`
	Value           float32          `json:"value"`
}

func (sc StatChange) GetAPIResource() APIResource {
	return sc.Stat
}

func (sc StatChange) IsZero() bool {
	return sc.Stat.Name == ""
}

func convertStatChange(srv *Server, sc seedtypes.StatChange) StatChange {
	return StatChange{
		Stat:            nameToNamedAPIResource(srv, E.Stats, sc.StatName, nil),
		CalculationType: sc.CalculationType,
		Value:           sc.Value,
	}
}

type ModifierChange struct {
	Modifier        NamedAPIResource `json:"modifier"`
	CalculationType string           `json:"calculation_type"`
	Value           float32          `json:"value"`
}

func (mc ModifierChange) GetAPIResource() APIResource {
	return mc.Modifier
}

func (mc ModifierChange) IsZero() bool {
	return mc.Modifier.Name == ""
}

func convertModifierChange(srv *Server, mc seedtypes.ModifierChange) ModifierChange {
	return ModifierChange{
		Modifier:        nameToNamedAPIResource(srv, E.Modifiers, mc.ModifierName, nil),
		CalculationType: mc.CalculationType,
		Value:           mc.Value,
	}
}
