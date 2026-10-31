package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getTriggerCommand(r *http.Request, i handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedAPIResourceList], id int32) (TriggerCommand, error) {
	command, err := verifyParamsAndGet(r, i, id)
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

	battleInteractions, err := applyUser(cfg, r, i, response, QpnAbilityUser)
	if err != nil {
		return TriggerCommand{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrieveTriggerCommands(r *http.Request, i handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		nameIDQuery(r, i, ids, QpnUser, cfg.E.CharacterClasses.resTypeSingle, cfg.L.CharClasses, cfg.db.GetTriggerCommandIDsByCharClass),
		nameIDQuery(r, i, ids, QpnRelatedStat, cfg.E.Stats.resTypeSingle, cfg.L.Stats, cfg.db.GetTriggerCommandIDsByRelatedStat),
	})
}
