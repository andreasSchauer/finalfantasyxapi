package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getTriggerCommand(r *http.Request, i handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList], id int32) (TriggerCommand, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return TriggerCommand{}, err
	}
	
	monsterFormations, err := getResourcesDB(cfg, r, cfg.e.monsterFormations, ability, cfg.db.GetTriggerCommandMonsterFormationIDs)
	if err != nil {
		return TriggerCommand{}, err
	}

	users, err := getResourcesDB(cfg, r, cfg.e.characterClasses, ability, cfg.db.GetTriggerCommandCharClassIDs)
	if err != nil {
		return TriggerCommand{}, err
	}

	response := TriggerCommand{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Version:               ability.Version,
		Specification: 		   ability.Specification,
		Description:           ability.Description,
		Effect:                ability.Effect,
		Rank:                  ability.Rank,
		AppearsInHelpBar:      ability.AppearsInHelpBar,
		CanCopycat:            ability.CanCopycat,
		UsedBy:                users,
		RelatedStats:          namesToNamedAPIResources(cfg, cfg.e.stats, ability.RelatedStats),
		Topmenu:               namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, ability.Topmenu, nil),
		Cursor:                ability.Cursor,
		MonsterFormations:     monsterFormations,
		BattleInteractions:    convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}



func (cfg *Config) retrieveTriggerCommands(r *http.Request, i handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(intQueryNullable(cfg, r, i, resources, "rank", cfg.db.GetTriggerCommandIDsByRank)),
		frl(nameOrIdQuery(cfg, r, i, resources, "char_class", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetTriggerCommandIDsByCharClass)),
		frl(nameOrIdQuery(cfg, r, i, resources, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetTriggerCommandIDsByRelatedStat)),
		frl(boolQuery2(cfg, r, i, resources, "stat_changes", cfg.db.GetTriggerCommandIDsWithStatChanges)),
		frl(boolQuery2(cfg, r, i, resources, "mod_changes", cfg.db.GetTriggerCommandIDsWithModifierChanges)),
	})
}
