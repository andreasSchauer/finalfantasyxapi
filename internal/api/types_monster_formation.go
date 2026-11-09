package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func ConvertEncounterArea(ea seedtypes.EncounterArea) apitypes.EncounterArea {
	return apitypes.EncounterArea{
		Area:          handlerinput.LocAreaToAreaAPIResource(apilookups.E.Areas, ea.LocationArea),
		Specification: ea.Specification,
		Availability:  ea.Availability,
	}
}

func ConvertFormationTriggerCommand(tc seedtypes.FormationTriggerCommand) apitypes.FormationTriggerCommand {
	return apitypes.FormationTriggerCommand{
		Ability:   handlerinput.NameToNamedAPIResource(apilookups.E.TriggerCommands, tc.Name, tc.Version),
		Condition: tc.Condition,
		UseAmount: tc.UseAmount,
		Users:     handlerinput.NamesToNamedAPIResources(apilookups.E.CharacterClasses, tc.Users),
	}
}

func ConvertFormationBossSong(bossMusic seedtypes.FormationBossSong) apitypes.FormationBossSong {
	return apitypes.FormationBossSong{
		Song:             handlerinput.NameToNamedAPIResource(apilookups.E.Songs, bossMusic.Song, nil),
		CelebrateVictory: bossMusic.CelebrateVictory,
	}
}
