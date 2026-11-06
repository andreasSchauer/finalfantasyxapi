package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertBackgroundMusic(bm seedtypes.BackgroundMusic) apitypes.BackgroundMusic {
	return apitypes.BackgroundMusic{
		Condition:              bm.Condition,
		ReplacesEncounterMusic: bm.ReplacesEncounterMusic,
		Areas:                  locAreasToAreaAPIResources(E.Areas, bm.LocationAreas),
	}
}

func convertCue(cue seedtypes.Cue) apitypes.Cue {
	var triggerAreaPtr *apiresource.AreaAPIResource

	if cue.TriggerLocationArea != nil {
		triggerArea := locAreaToAreaAPIResource(E.Areas, *cue.TriggerLocationArea)
		triggerAreaPtr = &triggerArea
	}

	return apitypes.Cue{
		SceneDescription:       cue.SceneDescription,
		TriggerArea:            triggerAreaPtr,
		IncludedAreas:          locAreasToAreaAPIResources(E.Areas, cue.IncludedAreas),
		ReplacesEncounterMusic: cue.ReplacesEncounterMusic,
		ReplacesBGMusic:        cue.ReplacesBGMusic,
		EndTrigger:             cue.EndTrigger,
	}
}
