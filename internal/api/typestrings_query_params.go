package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type NamedParam string

type QueryParamName NamedParam

const (
	QpnAbilityUser      QueryParamName = "ability_user"
	QpnAeons            QueryParamName = "aeons"
	QpnAgility          QueryParamName = "agility"
	QpnAirship          QueryParamName = "airship"
	QpnAlteredState     QueryParamName = "altered_state"
	QpnAmbush           QueryParamName = "ambush"
	QpnAnima            QueryParamName = "anima"
	QpnArea             QueryParamName = "area"
	QpnArranger         QueryParamName = "arranger"
	QpnAttackType       QueryParamName = "attack_type"
	QpnAutoAbilities    QueryParamName = "auto_abilities"
	QpnAutoAbility      QueryParamName = "auto_ability"
	QpnAvailability     QueryParamName = "availability"
	QpnBattles          QueryParamName = "battles"
	QpnBDL              QueryParamName = "bdl"
	QpnBest             QueryParamName = "best"
	QpnBombWpn          QueryParamName = "bomb_wpn"
	QpnBossFights       QueryParamName = "boss_fights"
	QpnCanCrit          QueryParamName = "can_crit"
	QpnCapture          QueryParamName = "capture"
	QpnCategory         QueryParamName = "category"
	QpnCelestialWeapon  QueryParamName = "celestial_weapon"
	QpnChangesOnly      QueryParamName = "changes_only"
	QpnCharacter        QueryParamName = "character"
	QpnCharacters       QueryParamName = "characters"
	QpnChocobo          QueryParamName = "chocobo"
	QpnColor            QueryParamName = "color"
	QpnComposer         QueryParamName = "composer"
	QpnCompSphere       QueryParamName = "comp_sphere"
	QpnContainsItem     QueryParamName = "contains_item"
	QpnCopycat          QueryParamName = "copycat"
	QpnCreationArea     QueryParamName = "creation_area"
	QpnCustomize        QueryParamName = "customize"
	QpnDamageFormula    QueryParamName = "damage_formula"
	QpnDamageType       QueryParamName = "damage_type"
	QpnDarkable         QueryParamName = "darkable"
	QpnDelay            QueryParamName = "delay"
	QpnDistance         QueryParamName = "distance"
	QpnElement          QueryParamName = "element"
	QpnElementalResists QueryParamName = "elemental_resists"
	QpnEmptySlots       QueryParamName = "empty_slots"
	QpnEscape           QueryParamName = "escape"
	QpnEquipment        QueryParamName = "equipment"
	QpnExpSg            QueryParamName = "exp_sg"
	QpnFlip             QueryParamName = "flip"
	QpnFMVs             QueryParamName = "fmvs"
	QpnFormula          QueryParamName = "formula"
	QpnHasAbility       QueryParamName = "has_ability"
	QpnHasOverdrive     QueryParamName = "has_overdrive"
	QpnHelpBar          QueryParamName = "help_bar"
	QpnIDs              QueryParamName = "ids"
	QpnInflictMax       QueryParamName = "inflict_max"
	QpnInflictMin       QueryParamName = "inflict_min"
	QpnIsForced         QueryParamName = "is_forced"
	QpnItem             QueryParamName = "item"
	QpnItems            QueryParamName = "items"
	QpnKeyItem          QueryParamName = "key_item"
	QpnKimahriStats     QueryParamName = "kimahri_stats"
	QpnLearnItem        QueryParamName = "learn_item"
	QpnLimit            QueryParamName = "limit"
	QpnLocation         QueryParamName = "location"
	QpnLootType         QueryParamName = "loot_type"
	QpnMethods          QueryParamName = "methods"
	QpnModChanges       QueryParamName = "mod_changes"
	QpnMonster          QueryParamName = "monster"
	QpnMonsters         QueryParamName = "monsters"
	QpnMonsterItems     QueryParamName = "monster_items"
	QpnMp               QueryParamName = "mp"
	QpnMpMin            QueryParamName = "mp_min"
	QpnMpMax            QueryParamName = "mp_max"
	QpnOffset           QueryParamName = "offset"
	QpnOmnisElements    QueryParamName = "omnis_elements"
	QpnOptional         QueryParamName = "optional"
	QpnOsgStats         QueryParamName = "osg_stats"
	QpnOutsideBattle    QueryParamName = "outside_battle"
	QpnPreAirship       QueryParamName = "pre_airship"
	QpnRank             QueryParamName = "rank"
	QpnReflectable      QueryParamName = "reflectable"
	QpnRelatedStat      QueryParamName = "related_stat"
	QpnRelAvailability  QueryParamName = "rel_availability"
	QpnRelRepeatable    QueryParamName = "rel_repeatable"
	QpnRepeatable       QueryParamName = "repeatable"
	QpnReqItem          QueryParamName = "req_item"
	QpnResistance       QueryParamName = "resistance"
	QpnRonsoRage        QueryParamName = "ronso_rage"
	QpnSaveSphere       QueryParamName = "save_sphere"
	QpnSecondItem       QueryParamName = "second_item"
	QpnShop             QueryParamName = "shop"
	QpnShops            QueryParamName = "shops"
	QpnSidequests       QueryParamName = "sidequests"
	QpnSilenceable      QueryParamName = "silenceable"
	QpnSpecialUse       QueryParamName = "special_use"
	QpnSpecies          QueryParamName = "species"
	QpnStatChanges      QueryParamName = "stat_changes"
	QpnStatusInflict    QueryParamName = "status_inflict"
	QpnStatusRemove     QueryParamName = "status_remove"
	QpnStatusResists    QueryParamName = "status_resists"
	QpnStoryBased       QueryParamName = "story_based"
	QpnSublocation      QueryParamName = "sublocation"
	QpnStdSg            QueryParamName = "std_sg"
	QpnTable            QueryParamName = "table"
	QpnTargetType       QueryParamName = "target_type"
	QpnTopmenu          QueryParamName = "topmenu"
	QpnTreasures        QueryParamName = "treasures"
	QpnTreasureType     QueryParamName = "treasure_type"
	QpnType             QueryParamName = "type"
	QpnUnderwater       QueryParamName = "underwater"
	QpnUser             QueryParamName = "user"
	QpnUserAtk          QueryParamName = "user_atk"
	QpnYunaStats        QueryParamName = "yuna_stats"
	QpnZombie           QueryParamName = "zombie"
)

func QpnsToNamedParams(qpns []QueryParamName) []NamedParam {
	if qpns == nil {
		return nil
	}

	nps := make([]NamedParam, len(qpns))

	for i, sn := range qpns {
		nps[i] = NamedParam(sn)
	}

	return nps
}

func formatQpnSlice(qpns []QueryParamName) string {
	if qpns == nil {
		return ""
	}

	strings := []string{}

	for _, qpn := range qpns {
		strings = append(strings, string(qpn))
	}

	return h.FormatStringSlice(strings)
}

type QueryParamType string

const (
	qptBool          QueryParamType = "bool"
	qptEnum          QueryParamType = "enum"
	qptEnumList      QueryParamType = "enum-list"
	qptID            QueryParamType = "id"
	qptIDNul         QueryParamType = "id-nul"
	qptIDList        QueryParamType = "id-list"
	qptInt           QueryParamType = "int"
	qptIntList       QueryParamType = "int-list"
	qptNameID        QueryParamType = "name/id"
	qptNameIDList    QueryParamType = "name/id-list"
	qptNameIDListNul QueryParamType = "name/id-list-nul"
	qptOther         QueryParamType = "other"
	qptStat          QueryParamType = "stat"
	qptValue         QueryParamType = "value"
	qptValueList     QueryParamType = "value-list"
)

type QueryValue string

const (
	qvBlitzball QueryValue = "blitzball"
	qvMonster   QueryValue = "monster"
	qvQuest     QueryValue = "quest"
	qvShop      QueryValue = "shop"
	qvTreasure  QueryValue = "treasure"
	qvSteal     QueryValue = "steal"
	qvDrop      QueryValue = "drop"
	qvBribe     QueryValue = "bribe"
	qvOther     QueryValue = "other"
	qvF         QueryValue = "f"
	qvL         QueryValue = "l"
	qvW         QueryValue = "w"
	qvI         QueryValue = "i"
)

func qvsToStrings(qvs []QueryValue) []string {
	if qvs == nil {
		return nil
	}

	strings := []string{}

	for _, qv := range qvs {
		strings = append(strings, string(qv))
	}

	return strings
}

func formatQvSlice(qvs []QueryValue) string {
	if qvs == nil {
		return ""
	}
	strings := qvsToStrings(qvs)
	return h.FormatStringSlice(strings)
}

type QuerySpecialVal string

const (
	qsvMax      QuerySpecialVal = "max"
	qsvImmune   QuerySpecialVal = "immune"
	qsvInfinite QuerySpecialVal = "infinite"
	qsvAlways   QuerySpecialVal = "always"
)
