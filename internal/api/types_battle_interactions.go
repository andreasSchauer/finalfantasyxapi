package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"



type BattleInteraction struct {
	Target            			string				`json:"target"`
	BasedOnPhysAttack 			bool				`json:"based_on_phys_attack"`
	Range             			*int32				`json:"range"`
	Damage			  			*Damage				`json:"damage"`
	ShatterRate       			*int32				`json:"shatter_rate"`
	Accuracy					Accuracy			`json:"accuracy"`
	Darkable          			bool				`json:"darkable"`
	Silenceable       			bool				`json:"silenceable"`
	Reflectable       			bool				`json:"reflectable"`
	HitAmount         			int32				`json:"hit_amount"`
	SpecialAction     			*string				`json:"special_action"`
	InflictedDelay            	[]InflictedDelay  	`json:"inflicted_delay"`
	InflictedStatusConditions 	[]InflictedStatus 	`json:"inflicted_status_conditions"`
	RemovedStatusConditions		[]NamedAPIResource	`json:"removed_status_conditions"`
	CopiedStatusConditions		[]InflictedStatus	`json:"copied_status_conditions"`
	StatChanges					[]StatChange		`json:"stat_changes"`
	ModifierChanges				[]ModifierChange	`json:"modifier_changes"`
}

func convertBattleInteraction(cfg *Config, ba seeding.BattleInteraction) BattleInteraction {
	battleInteraction := BattleInteraction{
		Target: ba.Target,
		BasedOnPhysAttack: ba.BasedOnPhysAttack,
		Range: ba.Range,
		Damage: convertObjPtr(cfg, ba.Damage, convertDamage),
		ShatterRate: ba.ShatterRate,
		Accuracy: convertAccuracy(cfg, ba.Accuracy),
		HitAmount: ba.HitAmount,
		SpecialAction: ba.SpecialAction,
		InflictedDelay: convertObjSlice(cfg, ba.InflictedDelay, convertInflictedDelay),
		InflictedStatusConditions: convertObjSlice(cfg, ba.InflictedStatusConditions, convertInflictedStatus),
		RemovedStatusConditions: namesToNamedAPIResources(cfg, cfg.e.statusConditions, ba.RemovedStatusConditions),
		CopiedStatusConditions: convertObjSlice(cfg, ba.CopiedStatusConditions, convertInflictedStatus),
		StatChanges: convertObjSlice(cfg, ba.StatChanges, convertStatChange),
		ModifierChanges: convertObjSlice(cfg, ba.ModifierChanges, convertModifierChange),
	}

	for _, status := range ba.AffectedBy {
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
	DamageCalc      []AbilityDamage 	`json:"damage_calc"`
	Critical        *string         	`json:"critical"`
	CriticalPlusVal *int32          	`json:"critical_plus_val"`
	IsPiercing      bool            	`json:"is_piercing"`
	BreakDmgLimit   *string         	`json:"break_dmg_lmt"`
	Element         *NamedAPIResource 	`json:"element"`		
}

func convertDamage(cfg *Config, d seeding.Damage) Damage {
	return Damage{
		DamageCalc: 		convertObjSlice(cfg, d.DamageCalc, convertAbilityDamage),
		Critical: 			d.Critical,
		CriticalPlusVal: 	d.CriticalPlusVal,
		IsPiercing: 		d.IsPiercing,
		BreakDmgLimit: 		d.BreakDmgLimit,
		Element: 			namePtrToNamedAPIResPtr(cfg, cfg.e.elements, d.Element, nil),
	}
}




type AbilityDamage struct {
	Condition      *string 				`json:"condition"`
	AttackType     NamedAPIResource  	`json:"attack_type"`
	TargetStat     NamedAPIResource 	`json:"target_stat"`
	DamageType     NamedAPIResource 	`json:"damage_type"`
	DamageFormula  NamedAPIResource 	`json:"damage_formula"`
	DamageConstant int32  				`json:"damage_constant"`
}

func convertAbilityDamage(cfg *Config, ad seeding.AbilityDamage) AbilityDamage {
	attackType, _ := newNamedAPIResourceFromType(cfg, cfg.e.attackType.endpoint, ad.AttackType, cfg.t.AttackType)
	damageType, _ := newNamedAPIResourceFromType(cfg, cfg.e.damageType.endpoint, ad.DamageType, cfg.t.DamageType)
	damageFormula, _ := newNamedAPIResourceFromType(cfg, cfg.e.damageFormula.endpoint, ad.DamageFormula, cfg.t.DamageFormula)

	return AbilityDamage{
		Condition: 		ad.Condition,
		AttackType: 	attackType,
		TargetStat:		nameToNamedAPIResource(cfg, cfg.e.stats, ad.TargetStat, nil),
		DamageType: 	damageType,
		DamageFormula: 	damageFormula,
		DamageConstant: ad.DamageConstant,
	}
}




type Accuracy struct {
	AccSource   string   `json:"acc_source"`
	HitChance   *int32   `json:"hit_chance"`
	AccModifier *float32 `json:"acc_modifier"`
}

func convertAccuracy(_ *Config, a seeding.Accuracy) Accuracy {
	return Accuracy{
		AccSource: 		a.AccSource,
		HitChance: 		a.HitChance,
		AccModifier: 	a.AccModifier,
	}
}




type InflictedDelay struct {
	Condition      *string `json:"condition"`
	CTBAttackType  string  `json:"ctb_attack_type"`
	DelayType      string  `json:"delay_type"`
	DamageConstant int32   `json:"damage_constant"`
}

func convertInflictedDelay(_ *Config, id seeding.InflictedDelay) InflictedDelay {
	return InflictedDelay{
		Condition: 		id.Condition,
		CTBAttackType: 	id.CTBAttackType,
		DelayType: 		id.DelayType,
		DamageConstant: id.DamageConstant,
	}
}




type InflictedStatus struct {
	StatusCondition NamedAPIResource `json:"status_condition"`
	Probability     int32            `json:"probability,omitempty"`
	DurationType    string           `json:"duration_type,omitempty"`
	Amount          *int32           `json:"amount,omitempty"`
}

func (is InflictedStatus) GetAPIResource() APIResource {
	return is.StatusCondition
}

func (is InflictedStatus) IsZero() bool {
	return is.StatusCondition.Name == ""
}

func convertInflictedStatus(cfg *Config, status seeding.InflictedStatus) InflictedStatus {
	return InflictedStatus{
		StatusCondition: nameToNamedAPIResource(cfg, cfg.e.statusConditions, status.StatusCondition, nil),
		Probability:     status.Probability,
		DurationType:    status.DurationType,
		Amount:          status.Amount,
	}
}



type StatChange struct {
	Stat			NamedAPIResource	`json:"stat"`
	CalculationType	string				`json:"calculation_type"`
	Value			float32				`json:"value"`
}

func (sc StatChange) GetAPIResource() APIResource {
	return sc.Stat
}

func (sc StatChange) IsZero() bool {
	return sc.Stat.Name == ""
}

func convertStatChange(cfg *Config, sc seeding.StatChange) StatChange {
	return StatChange{
		Stat: 				nameToNamedAPIResource(cfg, cfg.e.stats, sc.StatName, nil),
		CalculationType: 	sc.CalculationType,
		Value: 				sc.Value,
	}
}



type ModifierChange struct {
	Modifier		NamedAPIResource	`json:"modifier"`
	CalculationType	string				`json:"calculation_type"`
	Value			float32				`json:"value"`
}

func (mc ModifierChange) GetAPIResource() APIResource {
	return mc.Modifier
}

func (mc ModifierChange) IsZero() bool {
	return mc.Modifier.Name == ""
}

func convertModifierChange(cfg *Config, mc seeding.ModifierChange) ModifierChange {
	return ModifierChange{
		Modifier: 			nameToNamedAPIResource(cfg, cfg.e.modifiers, mc.ModifierName, nil),
		CalculationType: 	mc.CalculationType,
		Value: 				mc.Value,
	}
}