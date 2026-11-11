package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Overdrive(r *http.Request, i handlerinput.Main[seedtypes.Overdrive, apitypes.Overdrive, apiresource.Named], id int32) (apitypes.Overdrive, error) {
	overdrive, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Overdrive{}, err
	}

	response := apitypes.Overdrive{
		ID:                 overdrive.ID,
		Name:               overdrive.Name,
		Version:            overdrive.Version,
		Specification:      overdrive.Specification,
		Description:        overdrive.Description,
		Effect:             overdrive.Effect,
		Rank:               overdrive.Rank,
		AppearsInHelpBar:   overdrive.AppearsInHelpBar,
		CanCopycat:         overdrive.CanCopycat,
		UnlockCondition:    overdrive.UnlockCondition,
		CountdownInSec:     overdrive.CountdownInSec,
		Cursor:             overdrive.Cursor,
		User:               handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, overdrive.User, nil),
		OverdriveCommand:   handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.OverdriveCommands, overdrive.OverdriveCommand, nil),
		OverdriveAbilities: apiconvert.RefsToNamedAPIResources(overdrive.OverdriveAbilities),
	}

	return response, nil
}
