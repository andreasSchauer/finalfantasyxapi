package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertEncounterArea(ea seedtypes.EncounterArea) apitypes.EncounterArea {
	return apitypes.EncounterArea{
		Area:          handlerinput.LocAreaToAreaAPIResource(E.Areas, ea.LocationArea),
		Specification: ea.Specification,
		Availability:  ea.Availability,
	}
}

func convertFormationTriggerCommand(tc seedtypes.FormationTriggerCommand) apitypes.FormationTriggerCommand {
	return apitypes.FormationTriggerCommand{
		Ability:   handlerinput.NameToNamedAPIResource(E.TriggerCommands, tc.Name, tc.Version),
		Condition: tc.Condition,
		UseAmount: tc.UseAmount,
		Users:     handlerinput.NamesToNamedAPIResources(E.CharacterClasses, tc.Users),
	}
}

func convertFormationBossSong(bossMusic seedtypes.FormationBossSong) apitypes.FormationBossSong {
	return apitypes.FormationBossSong{
		Song:             handlerinput.NameToNamedAPIResource(E.Songs, bossMusic.Song, nil),
		CelebrateVictory: bossMusic.CelebrateVictory,
	}
}
