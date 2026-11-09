package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetAeonCommandRelationships(r *http.Request, command seedtypes.AeonCommand) apitypes.AeonCommand {
	var rel apitypes.AeonCommand
	userName := command.PossibleAbilities[0].User

	if len(command.PossibleAbilities) == 3 {
		userName = "magus-sisters"
	}

	rel.User = handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, userName, nil)

	return rel
}
