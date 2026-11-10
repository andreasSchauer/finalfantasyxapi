package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconv"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func GetAeonCommand(r *http.Request, i handlerinput.Main[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.Named], id int32) (apitypes.AeonCommand, error) {
	command, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.AeonCommand{}, err
	}

	rel := getAeonCommandRelationships(command)

	response := apitypes.AeonCommand{
		ID:                command.ID,
		Name:              command.Name,
		Description:       command.Description,
		Effect:            command.Effect,
		Cursor:            command.Cursor,
		User:              rel.User,
		Topmenu:           handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Topmenus, command.Topmenu, nil),
		OpenSubmenu:       handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Submenus, command.OpenSubmenu, nil),
		PossibleAbilities: ah.ConvertObjSlice(command.PossibleAbilities, apiconv.ConvertPossibleAbilityList),
	}

	return response, nil
}

func getAeonCommandRelationships(command seedtypes.AeonCommand) apitypes.AeonCommand {
	var rel apitypes.AeonCommand
	userName := command.PossibleAbilities[0].User

	if len(command.PossibleAbilities) == 3 {
		userName = "magus-sisters"
	}

	rel.User = handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, userName, nil)

	return rel
}
