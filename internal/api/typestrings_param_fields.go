package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type FieldName string

const (
	PfnAbility           FieldName = "ability"
	PfnAbilityID         FieldName = "ability_id"
	PfnAgility           FieldName = "agility"
	PfnAglOverride       FieldName = "agl_override"
	PfnAltState          FieldName = "alt_state"
	PfnAmt               FieldName = "amt"
	PfnAttackType        FieldName = "attack_type"
	PfnAutoAbilities     FieldName = "auto_abilities"
	PfnBattles           FieldName = "battles"
	PfnBattleStart       FieldName = "battle_start"
	PfnCharacterID       FieldName = "character_id"
	PfnCritical			 FieldName = "critical"
	PfnCriticalPlusChar	 FieldName = "critical_plus_char"
	PfnCriticalPlusMon	 FieldName = "critical_plus_mon"
	PfnCriticalPlusVal	 FieldName = "critical_plus_val"
	PfnCustom		 	 FieldName = "custom"
	PfnDecimals          FieldName = "decimals"
	PfnDelay             FieldName = "delay"
	PfnDelayConstant     FieldName = "delay_constant"
	PfnDelayType         FieldName = "delay_type"
	PfnDirection         FieldName = "direction"
	PfnEquipType         FieldName = "equip_type"
	PfnFirstStrike       FieldName = "first_strike"
	PfnFormationID       FieldName = "formation_id"
	PfnIgnFirstTurn      FieldName = "ign_first_turn"
	PfnIgnImmunities     FieldName = "ign_immunities"
	PfnInteraction       FieldName = "interaction"
	PfnJinxStacks		 FieldName = "jinx_stacks"
	PfnLookup			 FieldName = "lookup"
	PfnLuck				 FieldName = "luck"
	PfnLuckStacks		 FieldName = "luck_stacks"
	PfnMinEmptySlots     FieldName = "min_empty_slots"
	PfnMons              FieldName = "mons"
	PfnMonsterID         FieldName = "monster_id"
	PfnName              FieldName = "name"
	PfnOptionalAbilities FieldName = "optional_abilities"
	PfnOverrides 		 FieldName = "overrides"
	PfnParty             FieldName = "party"
	PfnPartyMembers      FieldName = "party_members"
	PfnRemainingTicks    FieldName = "remaining_ticks"
	PfnRNG               FieldName = "rng"
	PfnStatus            FieldName = "status"
	PfnStrength          FieldName = "strength"
	PfnTarget            FieldName = "target"
	PfnText              FieldName = "text"
	PfnTotalSlots        FieldName = "total_slots"
	PfnTurnsAmt          FieldName = "turns_amt"
	PfnUnitID            FieldName = "unit_id"
	PfnUser				 FieldName = "user"
)

func formatPfnSlice(pfns []FieldName) string {
	if pfns == nil {
		return ""
	}

	strings := []string{}

	for _, qpn := range pfns {
		strings = append(strings, string(qpn))
	}

	return h.FormatStringSlice(strings)
}

func pfnPtr(fieldName FieldName) *FieldName {
	return &fieldName
}
