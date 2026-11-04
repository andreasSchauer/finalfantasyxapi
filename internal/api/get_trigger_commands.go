package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getTriggerCommand(r *http.Request, i handlerInput[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource], id int32) (TriggerCommand, error) {
	command, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return TriggerCommand{}, err
	}

	rel, err := getTriggerCommandRelationships(s, r, command)
	if err != nil {
		return TriggerCommand{}, err
	}

	response := TriggerCommand{
		ID:                 command.ID,
		Name:               command.Name,
		Version:            command.Version,
		Specification:      command.Specification,
		UntypedAbility:     idToTypedAPIResource(s, E.Abilities, command.Ability.ID),
		Description:        command.Description,
		Effect:             command.Effect,
		Rank:               command.Rank,
		AppearsInHelpBar:   command.AppearsInHelpBar,
		CanCopycat:         command.CanCopycat,
		UsedBy:             rel.UsedBy,
		RelatedStats:       namesToNamedAPIResources(s, E.Stats, command.RelatedStats),
		Topmenu:            namePtrToNamedAPIResPtr(s, E.Topmenus, command.Topmenu, nil),
		Cursor:             command.Cursor,
		MonsterFormations:  rel.MonsterFormations,
		BattleInteractions: convertObjSlice(s, command.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(s, r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return TriggerCommand{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (s *Server) retrieveTriggerCommands(r *http.Request, i handlerInput[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, s.Cfg.DB.GetTriggerCommandIDsByCharClass),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, s.Cfg.DB.GetTriggerCommandIDsByRelatedStat),
	})
}
