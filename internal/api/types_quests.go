package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func ConvertQuestCompletion(qc seedtypes.QuestCompletion) apitypes.QuestCompletion {
	return apitypes.QuestCompletion{
		Condition: qc.Condition,
		Areas:     ah.ConvertObjSlice(qc.Areas, ConvertCompletionArea),
		Reward:    apilookups.E.AllItems.ToResAmtFromName(qc.Reward),
	}
}

func ConvertCompletionArea(cl seedtypes.CompletionArea) apitypes.CompletionArea {
	return apitypes.CompletionArea{
		Area:  handlerinput.LocAreaToAreaAPIResource(apilookups.E.Areas, cl.LocationArea),
		Notes: cl.Notes,
	}
}
