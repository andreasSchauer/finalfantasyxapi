package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getTriggerCommand(r *http.Request, i handlerInput[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource], id int32) (TriggerCommand, error) {
	command, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return TriggerCommand{}, err
	}

	rel, err := getTriggerCommandRelationships(cfg, r, command)
	if err != nil {
		return TriggerCommand{}, err
	}

	response := TriggerCommand{
		ID:                 command.ID,
		Name:               command.Name,
		Version:            command.Version,
		Specification:      command.Specification,
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.E.Abilities, command.Ability.ID),
		Description:        command.Description,
		Effect:             command.Effect,
		Rank:               command.Rank,
		AppearsInHelpBar:   command.AppearsInHelpBar,
		CanCopycat:         command.CanCopycat,
		UsedBy:             rel.UsedBy,
		RelatedStats:       namesToNamedAPIResources(cfg, cfg.E.Stats, command.RelatedStats),
		Topmenu:            namePtrToNamedAPIResPtr(cfg, cfg.E.Topmenus, command.Topmenu, nil),
		Cursor:             command.Cursor,
		MonsterFormations:  rel.MonsterFormations,
		BattleInteractions: convertObjSlice(cfg, command.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(cfg, r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return TriggerCommand{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrieveTriggerCommands(r *http.Request, i handlerInput[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryNameID(r, ids, ts.QpnUser, cfg.E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, cfg.db.GetTriggerCommandIDsByCharClass),
		i.queryNameID(r, ids, ts.QpnRelatedStat, cfg.E.Stats.resTypeSingle, lookup.L.Stats, cfg.db.GetTriggerCommandIDsByRelatedStat),
	})
}
