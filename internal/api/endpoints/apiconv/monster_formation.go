package apiconv

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
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
