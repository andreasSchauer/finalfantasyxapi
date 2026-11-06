package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertQuestCompletion(qc seedtypes.QuestCompletion) apitypes.QuestCompletion {
	return apitypes.QuestCompletion{
		Condition: qc.Condition,
		Areas:     ah.ConvertObjSlice(qc.Areas, convertCompletionArea),
		Reward:    E.AllItems.toResAmtFromName(qc.Reward),
	}
}

func convertCompletionArea(cl seedtypes.CompletionArea) apitypes.CompletionArea {
	return apitypes.CompletionArea{
		Area:  locAreaToAreaAPIResource(E.Areas, cl.LocationArea),
		Notes: cl.Notes,
	}
}
