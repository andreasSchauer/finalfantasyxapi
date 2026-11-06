package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertEncounterArea(ea seedtypes.EncounterArea) apitypes.EncounterArea {
	return apitypes.EncounterArea{
		Area:          locAreaToAreaAPIResource(E.Areas, ea.LocationArea),
		Specification: ea.Specification,
		Availability:  ea.Availability,
	}
}

func convertFormationTriggerCommand(tc seedtypes.FormationTriggerCommand) apitypes.FormationTriggerCommand {
	return apitypes.FormationTriggerCommand{
		Ability:   nameToNamedAPIResource(E.TriggerCommands, tc.Name, tc.Version),
		Condition: tc.Condition,
		UseAmount: tc.UseAmount,
		Users:     namesToNamedAPIResources(E.CharacterClasses, tc.Users),
	}
}

func convertFormationBossSong(bossMusic seedtypes.FormationBossSong) apitypes.FormationBossSong {
	return apitypes.FormationBossSong{
		Song:             nameToNamedAPIResource(E.Songs, bossMusic.Song, nil),
		CelebrateVictory: bossMusic.CelebrateVictory,
	}
}
