package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func OverdriveAbility(r *http.Request, i handlerinput.Main[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.Named], id int32) (apitypes.OverdriveAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.OverdriveAbility{}, err
	}
	
	response := apitypes.OverdriveAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     handlerinput.IDToTypedAPIResource(apilookups.E.Abilities, ability.Ability.ID),
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		RelatedStats:       handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, ability.RelatedStats),
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, apiconvert.BattleInteraction),
	}

	err = getOverdriveAbilityRelationships(r, ability, &response)
	if err != nil {
		return apitypes.OverdriveAbility{}, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive", err)
	}
	
	return response, nil
}

func getOverdriveAbilityRelationships(r *http.Request, ability seedtypes.OverdriveAbility, response *apitypes.OverdriveAbility) error {
	overdriveIDs, err := config.Cfg.DB.GetOverdriveAbilityOverdriveIDs(r.Context(), ability.ID)
	if err != nil {
		return errs.NewHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive", err)
	}
	overdrive, _ := lookup.GetResourceByID(overdriveIDs[0], lookup.L.OverdrivesID)

	response.Overdrives = apilookups.E.Overdrives.ToAPIResources(overdriveIDs)
	response.OverdriveCommand = handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.OverdriveCommands, overdrive.OverdriveCommand, nil)
	response.User = handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, overdrive.User, nil)

	return nil
}
