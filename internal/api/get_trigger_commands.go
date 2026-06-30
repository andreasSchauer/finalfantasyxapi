package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getTriggerCommand(r *http.Request, i handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList], id int32) (TriggerCommand, error) {
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
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.e.abilities, command.Ability.ID),
		Description:        command.Description,
		Effect:             command.Effect,
		Rank:               command.Rank,
		AppearsInHelpBar:   command.AppearsInHelpBar,
		CanCopycat:         command.CanCopycat,
		UsedBy:             rel.UsedBy,
		RelatedStats:       namesToNamedAPIResources(cfg, cfg.e.stats, command.RelatedStats),
		Topmenu:            namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, command.Topmenu, nil),
		Cursor:             command.Cursor,
		MonsterFormations:  rel.MonsterFormations,
		BattleInteractions: convertObjSlice(cfg, command.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(cfg, r, i, response, qpnAbilityUser)
	if err != nil {
		return TriggerCommand{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrieveTriggerCommands(r *http.Request, i handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		nameIdQuery(r, i, ids, qpnUser, cfg.e.characterClasses.resTypeSingle, cfg.l.CharClasses, cfg.db.GetTriggerCommandIDsByCharClass),
		nameIdQuery(r, i, ids, qpnRelatedStat, cfg.e.stats.resTypeSingle, cfg.l.Stats, cfg.db.GetTriggerCommandIDsByRelatedStat),
	})
}
