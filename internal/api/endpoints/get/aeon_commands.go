package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func AeonCommand(r *http.Request, i handlerinput.Main[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.Named], id int32) (apitypes.AeonCommand, error) {
	command, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.AeonCommand{}, err
	}

	
	response := apitypes.AeonCommand{
		ID:                command.ID,
		Name:              command.Name,
		Description:       command.Description,
		Effect:            command.Effect,
		Cursor:            command.Cursor,
		Topmenu:           handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Topmenus, command.Topmenu, nil),
		OpenSubmenu:       handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Submenus, command.OpenSubmenu, nil),
		PossibleAbilities: ah.ConvertObjSlice(command.PossibleAbilities, apiconvert.PossibleAbilityList),
	}
	
	getAeonCommandRelationships(command, &response)

	return response, nil
}

func getAeonCommandRelationships(command seedtypes.AeonCommand, response *apitypes.AeonCommand) {
	userName := command.PossibleAbilities[0].User

	if len(command.PossibleAbilities) == 3 {
		userName = "magus-sisters"
	}

	response.User = handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, userName, nil)
}
