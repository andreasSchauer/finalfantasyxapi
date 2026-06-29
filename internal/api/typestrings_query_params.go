package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type NamedParam string

type QueryParamName NamedParam

const (
	qpnAbilityUser      QueryParamName = "ability_user"
	qpnAeons            QueryParamName = "aeons"
	qpnAeonStats        QueryParamName = "aeon_stats"
	qpnAgility          QueryParamName = "agility"
	qpnAirship          QueryParamName = "airship"
	qpnAlteredState     QueryParamName = "altered_state"
	qpnAmbush           QueryParamName = "ambush"
	qpnAnima            QueryParamName = "anima"
	qpnArea             QueryParamName = "area"
	qpnArranger         QueryParamName = "arranger"
	qpnAttackType       QueryParamName = "attack_type"
	qpnAutoAbilities    QueryParamName = "auto_abilities"
	qpnAutoAbility      QueryParamName = "auto_ability"
	qpnAvailability     QueryParamName = "availability"
	qpnBattles          QueryParamName = "battles"
	qpnBDL              QueryParamName = "bdl"
	qpnBest             QueryParamName = "best"
	qpnBombWpn          QueryParamName = "bomb_wpn"
	qpnBossFights       QueryParamName = "boss_fights"
	qpnCanCrit          QueryParamName = "can_crit"
	qpnCapture          QueryParamName = "capture"
	qpnCategory         QueryParamName = "category"
	qpnCelestialWeapon  QueryParamName = "celestial_weapon"
	qpnChangesOnly      QueryParamName = "changes_only"
	qpnCharacter        QueryParamName = "character"
	qpnCharacters       QueryParamName = "characters"
	qpnChocobo          QueryParamName = "chocobo"
	qpnColor            QueryParamName = "color"
	qpnComposer         QueryParamName = "composer"
	qpnCompSphere       QueryParamName = "comp_sphere"
	qpnContainsItem     QueryParamName = "contains_item"
	qpnCopycat          QueryParamName = "copycat"
	qpnCreationArea     QueryParamName = "creation_area"
	qpnCustomize		QueryParamName = "customize"
	qpnDamageFormula    QueryParamName = "damage_formula"
	qpnDamageType       QueryParamName = "damage_type"
	qpnDarkable         QueryParamName = "darkable"
	qpnDelay            QueryParamName = "delay"
	qpnDistance         QueryParamName = "distance"
	qpnElement          QueryParamName = "element"
	qpnElementalResists QueryParamName = "elemental_resists"
	qpnEmptySlots       QueryParamName = "empty_slots"
	qpnEquipment        QueryParamName = "equipment"
	qpnExpSg            QueryParamName = "exp_sg"
	qpnFlip             QueryParamName = "flip"
	qpnFMVs             QueryParamName = "fmvs"
	qpnFormula          QueryParamName = "formula"
	qpnHasAbility       QueryParamName = "has_ability"
	qpnHasOverdrive     QueryParamName = "has_overdrive"
	qpnHelpBar          QueryParamName = "help_bar"
	qpnIDs              QueryParamName = "ids"
	qpnInflictMax       QueryParamName = "inflict_max"
	qpnInflictMin       QueryParamName = "inflict_min"
	qpnIsForced         QueryParamName = "is_forced"
	qpnItem             QueryParamName = "item"
	qpnItems            QueryParamName = "items"
	qpnKeyItem          QueryParamName = "key_item"
	qpnKimahriStats     QueryParamName = "kimahri_stats"
	qpnLearnItem        QueryParamName = "learn_item"
	qpnLimit            QueryParamName = "limit"
	qpnLocation         QueryParamName = "location"
	qpnLootType         QueryParamName = "loot_type"
	qpnMethods          QueryParamName = "methods"
	qpnModChanges       QueryParamName = "mod_changes"
	qpnMonster          QueryParamName = "monster"
	qpnMonsters         QueryParamName = "monsters"
	qpnMonsterItems     QueryParamName = "monster_items"
	qpnMp               QueryParamName = "mp"
	qpnMpMin            QueryParamName = "mp_min"
	qpnMpMax            QueryParamName = "mp_max"
	qpnOffset           QueryParamName = "offset"
	qpnOmnisElements    QueryParamName = "omnis_elements"
	qpnOptional         QueryParamName = "optional"
	qpnOsgStats			QueryParamName = "osg_stats"
	qpnOutsideBattle    QueryParamName = "outside_battle"
	qpnPreAirship       QueryParamName = "pre_airship"
	qpnRank             QueryParamName = "rank"
	qpnReflectable      QueryParamName = "reflectable"
	qpnRelatedStat      QueryParamName = "related_stat"
	qpnRelAvailability  QueryParamName = "rel_availability"
	qpnRelRepeatable    QueryParamName = "rel_repeatable"
	qpnRepeatable       QueryParamName = "repeatable"
	qpnReqItem          QueryParamName = "req_item"
	qpnResistance       QueryParamName = "resistance"
	qpnRonsoRage        QueryParamName = "ronso_rage"
	qpnSaveSphere       QueryParamName = "save_sphere"
	qpnSecondItem       QueryParamName = "second_item"
	qpnShop             QueryParamName = "shop"
	qpnShops            QueryParamName = "shops"
	qpnSidequests       QueryParamName = "sidequests"
	qpnSilenceable      QueryParamName = "silenceable"
	qpnSpecialUse       QueryParamName = "special_use"
	qpnSpecies          QueryParamName = "species"
	qpnStatChanges      QueryParamName = "stat_changes"
	qpnStatusInflict    QueryParamName = "status_inflict"
	qpnStatusRemove     QueryParamName = "status_remove"
	qpnStatusResists    QueryParamName = "status_resists"
	qpnStoryBased       QueryParamName = "story_based"
	qpnSublocation      QueryParamName = "sublocation"
	qpnStdSg            QueryParamName = "std_sg"
	qpnTable            QueryParamName = "table"
	qpnTargetType       QueryParamName = "target_type"
	qpnTopmenu          QueryParamName = "topmenu"
	qpnTreasures        QueryParamName = "treasures"
	qpnTreasureType     QueryParamName = "treasure_type"
	qpnType             QueryParamName = "type"
	qpnUnderwater       QueryParamName = "underwater"
	qpnUser             QueryParamName = "user"
	qpnUserAtk          QueryParamName = "user_atk"
	qpnYunaStats        QueryParamName = "yuna_stats"
	qpnZombie           QueryParamName = "zombie"
)

func qpnsToNamedParams(qpns []QueryParamName) []NamedParam {
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
	qptId            QueryParamType = "id"
	qptIdNul         QueryParamType = "id-nul"
	qptIdList        QueryParamType = "id-list"
	qptInt           QueryParamType = "int"
	qptIntList       QueryParamType = "int-list"
	qptNameId        QueryParamType = "name/id"
	qptNameIdList    QueryParamType = "name/id-list"
	qptNameIdListNul QueryParamType = "name/id-list-nul"
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
	qsvMax			QuerySpecialVal = "max"
	qsvImmune		QuerySpecialVal = "immune"
	qsvInfinite		QuerySpecialVal = "infinite"
	qsvAlways		QuerySpecialVal = "always"
)