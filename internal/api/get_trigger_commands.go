package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getTriggerCommand(r *http.Request, i handlerInput[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource], id int32) (TriggerCommand, error) {
	command, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return TriggerCommand{}, err
	}

	rel, err := getTriggerCommandRelationships(srv, r, command)
	if err != nil {
		return TriggerCommand{}, err
	}

	response := TriggerCommand{
		ID:                 command.ID,
		Name:               command.Name,
		Version:            command.Version,
		Specification:      command.Specification,
		UntypedAbility:     idToTypedAPIResource(srv, E.Abilities, command.Ability.ID),
		Description:        command.Description,
		Effect:             command.Effect,
		Rank:               command.Rank,
		AppearsInHelpBar:   command.AppearsInHelpBar,
		CanCopycat:         command.CanCopycat,
		UsedBy:             rel.UsedBy,
		RelatedStats:       namesToNamedAPIResources(srv, E.Stats, command.RelatedStats),
		Topmenu:            namePtrToNamedAPIResPtr(srv, E.Topmenus, command.Topmenu, nil),
		Cursor:             command.Cursor,
		MonsterFormations:  rel.MonsterFormations,
		BattleInteractions: convertObjSlice(srv, command.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(srv, r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return TriggerCommand{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (srv *Server) retrieveTriggerCommands(r *http.Request, i handlerInput[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetTriggerCommandIDsByCharClass),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, config.Cfg.DB.GetTriggerCommandIDsByRelatedStat),
	})
}
