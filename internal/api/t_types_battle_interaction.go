package api

type expBattleInteraction struct {
	target                    string
	abilityRange              *int32
	hitAmount                 int32
	shatterRate               int32
	specialAction             *string
	basedOnUserAtk            bool
	darkable                  bool
	silenceable               bool
	reflectable               bool
	accuracy                  expAccuracy
	damage                    *expDamage
	inflictedDelay            *expInflictedDelay
	inflictedStatusConditions []int32
	removedStatusConditions   []int32
	copiedStatusConditions    []int32
	statChanges               []expStatChange
	modifierChanges           []expModChange
}

func compareBattleInteractions(test test, fieldName string, exp expBattleInteraction, got BattleInteraction) {
	compare(test, fieldName+" target", exp.target, got.Target)
	compare(test, fieldName+" range", exp.abilityRange, got.Range)
	compare(test, fieldName+" hit amount", exp.hitAmount, got.HitAmount)
	compare(test, fieldName+" shatter rate", exp.shatterRate, got.ShatterRate)
	compare(test, fieldName+" special action", exp.specialAction, got.SpecialAction)
	compare(test, fieldName+" based on phys atk", exp.basedOnUserAtk, got.BasedOnUserAttack)
	compare(test, fieldName+" darkable", exp.darkable, got.Darkable)
	compare(test, fieldName+" silenceable", exp.silenceable, got.Silenceable)
	compare(test, fieldName+" reflectable", exp.reflectable, got.Reflectable)
	compareAccuracies(test, fieldName+" accuracy", exp.accuracy, got.Accuracy)
	compTestStructPtrs(test, fieldName+" damage", exp.damage, got.Damage, compareDamages)
	compTestStructPtrs(test, fieldName+" inflicted delay", exp.inflictedDelay, got.InflictedDelay, compareInflictedDelays)
	checkResIDsInSlice(test, fieldName+" inflicted status conditions", test.cfg.e.statusConditions.endpoint, exp.inflictedStatusConditions, got.InflictedStatusConditions)
	checkResIDsInSlice(test, fieldName+" removed status conditions", test.cfg.e.statusConditions.endpoint, exp.removedStatusConditions, got.RemovedStatusConditions)
	checkResIDsInSlice(test, fieldName+" copied status conditions", test.cfg.e.statusConditions.endpoint, exp.copiedStatusConditions, got.CopiedStatusConditions)
	compTestStructSlices(test, fieldName+" stat changes", exp.statChanges, got.StatChanges, compareStatChanges)
	compTestStructSlices(test, fieldName+" mod changes", exp.modifierChanges, got.ModifierChanges, compareModChanges)
}

type expAccuracy struct {
	accSource   string
	hitChance   *int32
	accModifier *float32
}

func compareAccuracies(test test, fieldName string, exp expAccuracy, got Accuracy) {
	compare(test, fieldName+" acc source", exp.accSource, got.AccSource)
	compare(test, fieldName+" hit chance", exp.hitChance, got.HitChance)
	compare(test, fieldName+" acc modifier", exp.accModifier, got.AccModifier)
}

type expDamage struct {
	damageCalc      []expAbilityDamage
	critical        *string
	criticalPlusVal *int32
	isPiercing      bool
	breakDmgLmt     *string
	element         *int32
}

func compareDamages(test test, fieldName string, exp expDamage, got Damage) {
	compTestStructSlices(test, "damage calc", exp.damageCalc, got.DamageCalc, compareAbilityDamages)
	compare(test, fieldName+" critical", exp.critical, got.Critical)
	compare(test, fieldName+" criticalPlusVal", exp.criticalPlusVal, got.CriticalPlusVal)
	compare(test, fieldName+" is piercing", exp.isPiercing, got.IsPiercing)
	compare(test, fieldName+" break dmg lmt", exp.breakDmgLmt, got.BreakDmgLimit)
	compIdApiResourcePtrs(test, fieldName+" element", test.cfg.e.elements.endpoint, exp.element, got.Element)
}

type expAbilityDamage struct {
	attackType     int32
	targetStat     int32
	damageType     int32
	damageFormula  int32
	damageConstant int32
}

func compareAbilityDamages(test test, fieldName string, exp expAbilityDamage, got AbilityDamage) {
	compIdApiResource(test, fieldName+" - ad attack type", test.cfg.e.attackType.endpoint, exp.attackType, got.AttackType)
	compIdApiResource(test, fieldName+" - ad target stat", test.cfg.e.stats.endpoint, exp.targetStat, got.TargetStat)
	compIdApiResource(test, fieldName+" - ad damage type", test.cfg.e.damageType.endpoint, exp.damageType, got.DamageType)
	compIdApiResource(test, fieldName+" - ad damage formula", test.cfg.e.damageFormula.endpoint, exp.damageFormula, got.DamageFormula)
	compare(test, fieldName+" - ad damage constant", exp.damageConstant, got.DamageConstant)
}

type expInflictedDelay struct {
	ctbAttackType  string
	delayType      string
	damageConstant int32
	delayStrength  string
}

func compareInflictedDelays(test test, fieldName string, exp expInflictedDelay, got InflictedDelay) {
	compare(test, fieldName+" ctb attack type", exp.ctbAttackType, got.CTBAttackType)
	compare(test, fieldName+" value", exp.delayType, got.DelayType)
	compare(test, fieldName+" delay damage constant", exp.damageConstant, got.DamageConstant)
	compare(test, fieldName+" delay strength", exp.delayStrength, got.DelayStrength)
}

type expStatChange struct {
	stat            int32
	calculationType string
	value           float32
}

func compareStatChanges(test test, fieldName string, exp expStatChange, got StatChange) {
	compIdApiResource(test, fieldName+" - sc stat", test.cfg.e.stats.endpoint, exp.stat, got.Stat)
	compare(test, fieldName+" - calculation type", exp.calculationType, got.CalculationType)
	compare(test, fieldName+" - value", exp.value, got.Value)
}

type expModChange struct {
	modifier        int32
	calculationType string
	value           float32
}

func compareModChanges(test test, fieldName string, exp expModChange, got ModifierChange) {
	compIdApiResource(test, fieldName+" - mc modifier", test.cfg.e.modifiers.endpoint, exp.modifier, got.Modifier)
	compare(test, fieldName+" - calculation type", exp.calculationType, got.CalculationType)
	compare(test, fieldName+" - value", exp.value, got.Value)
}
