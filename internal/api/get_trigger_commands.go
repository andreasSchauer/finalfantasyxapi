package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getTriggerCommand(r *http.Request, i handlerinput.Main[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.Named], id int32) (apitypes.TriggerCommand, error) {
	command, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.TriggerCommand{}, err
	}

	rel, err := getTriggerCommandRelationships(r, command)
	if err != nil {
		return apitypes.TriggerCommand{}, err
	}

	response := apitypes.TriggerCommand{
		ID:                 command.ID,
		Name:               command.Name,
		Version:            command.Version,
		Specification:      command.Specification,
		UntypedAbility:     handlerinput.IDToTypedAPIResource(apilookups.E.Abilities, command.Ability.ID),
		Description:        command.Description,
		Effect:             command.Effect,
		Rank:               command.Rank,
		AppearsInHelpBar:   command.AppearsInHelpBar,
		CanCopycat:         command.CanCopycat,
		UsedBy:             rel.UsedBy,
		RelatedStats:       handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, command.RelatedStats),
		Topmenu:            handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Topmenus, command.Topmenu, nil),
		Cursor:             command.Cursor,
		MonsterFormations:  rel.MonsterFormations,
		BattleInteractions: ah.ConvertObjSlice(command.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return apitypes.TriggerCommand{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func retrieveTriggerCommands(r *http.Request, i handlerinput.Main[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryNameID(r, ids, ts.QpnUser, apilookups.E.CharacterClasses.ResTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetTriggerCommandIDsByCharClass),
		i.QueryNameID(r, ids, ts.QpnRelatedStat, apilookups.E.Stats.ResTypeSingle, lookup.L.Stats, config.Cfg.DB.GetTriggerCommandIDsByRelatedStat),
	})
}
