package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type FieldName string

const (
	pfnAgility          FieldName = "agility"
	pfnAglOverride      FieldName = "agl_override"
	pfnAltState         FieldName = "alt_state"
	pfnAmt              FieldName = "amt"
	pfnAutoAbilities    FieldName = "auto_abilities"
	pfnBattles          FieldName = "battles"
	pfnBattleStart      FieldName = "battle_start"
	pfnCharacter        FieldName = "character"
	pfnDecimals         FieldName = "decimals"
	pfnDirection        FieldName = "direction"
	pfnEquipType        FieldName = "equip_type"
	pfnFirstStrike      FieldName = "first_strike"
	pfnFormation        FieldName = "formation"
	pfnID               FieldName = "id"
	pfnIgnFirstTurn     FieldName = "ign_first_turn"
	pfnIgnImmunities    FieldName = "ign_immunities"
	pfnLenientAbilities FieldName = "lenient_abilities"
	pfnMinEmptySlots    FieldName = "min_empty_slots"
	pfnMons             FieldName = "mons"
	pfnMonster          FieldName = "monster"
	pfnName             FieldName = "name"
	pfnParty            FieldName = "party"
	pfnPartyMembers     FieldName = "party_members"
	pfnRNG              FieldName = "rng"
	pfnStatus           FieldName = "status"
	pfnText             FieldName = "text"
	pfnTotalSlots       FieldName = "total_slots"
	pfnTurnsAmt         FieldName = "turns_amt"
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
