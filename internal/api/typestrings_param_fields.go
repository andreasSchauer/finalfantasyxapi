package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type FieldName string

const (
	pfnAgl			FieldName = "agl"
	pfnAglOverride	FieldName = "agl_override"
	pfnAltState		FieldName = "alt_state"
	pfnAmt			FieldName = "amt"
	pfnBattles		FieldName = "battles"
	pfnBattleStart	FieldName = "battle_start"
	pfnDirection 	FieldName = "direction"
	pfnFormation	FieldName = "formation"
	pfnFS			FieldName = "fs"
	pfnID			FieldName = "id"
	pfnIgnFirstTurn	FieldName = "ign_first_turn"
	pfnMons			FieldName = "mons"
	pfnMonsCustom	FieldName = "mons_custom"
	pfnName			FieldName = "name"
	pfnOffset		FieldName = "offset"
	pfnParty		FieldName = "party"
	pfnRNG			FieldName = "rng"
	pfnStatus		FieldName = "status"
	pfnText      	FieldName = "text"
	pfnTurnsAmt		FieldName = "turns_amt"
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