package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetOverdriveAbility(r *http.Request, i handlerinput.Main[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.Named], id int32) (apitypes.OverdriveAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.OverdriveAbility{}, err
	}

	rel, err := getOverdriveAbilityRelationships(r, ability)
	if err != nil {
		return apitypes.OverdriveAbility{}, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive", err)
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
		Overdrives:         rel.Overdrives,
		OverdriveCommand:   rel.OverdriveCommand,
		User:               rel.User,
		RelatedStats:       handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, ability.RelatedStats),
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, apiconv.ConvertBattleInteraction),
	}

	return response, nil
}

func getOverdriveAbilityRelationships(r *http.Request, ability seedtypes.OverdriveAbility) (apitypes.OverdriveAbility, error) {
	overdriveIDs, err := config.Cfg.DB.GetOverdriveAbilityOverdriveIDs(r.Context(), ability.ID)
	if err != nil {
		return apitypes.OverdriveAbility{}, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive", err)
	}
	overdrive, _ := lookup.GetResourceByID(overdriveIDs[0], lookup.L.OverdrivesID)

	rel := apitypes.OverdriveAbility{
		Overdrives:       apilookups.E.Overdrives.ToAPIResources(overdriveIDs),
		OverdriveCommand: handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.OverdriveCommands, overdrive.OverdriveCommand, nil),
		User:             handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, overdrive.User, nil),
	}

	return rel, nil
}
