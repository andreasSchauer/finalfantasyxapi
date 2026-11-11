package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func BackgroundMusic(bm seedtypes.BackgroundMusic) apitypes.BackgroundMusic {
	return apitypes.BackgroundMusic{
		Condition:              bm.Condition,
		ReplacesEncounterMusic: bm.ReplacesEncounterMusic,
		Areas:                  handlerinput.LocAreasToAreaAPIResources(apilookups.E.Areas, bm.LocationAreas),
	}
}

func Cue(cue seedtypes.Cue) apitypes.Cue {
	var triggerAreaPtr *apiresource.Area

	if cue.TriggerLocationArea != nil {
		triggerArea := handlerinput.LocAreaToAreaAPIResource(apilookups.E.Areas, *cue.TriggerLocationArea)
		triggerAreaPtr = &triggerArea
	}

	return apitypes.Cue{
		SceneDescription:       cue.SceneDescription,
		TriggerArea:            triggerAreaPtr,
		IncludedAreas:          handlerinput.LocAreasToAreaAPIResources(apilookups.E.Areas, cue.IncludedAreas),
		ReplacesEncounterMusic: cue.ReplacesEncounterMusic,
		ReplacesBGMusic:        cue.ReplacesBGMusic,
		EndTrigger:             cue.EndTrigger,
	}
}
