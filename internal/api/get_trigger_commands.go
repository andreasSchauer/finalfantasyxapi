package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getTriggerCommand(r *http.Request, i handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList], id int32) (TriggerCommand, error) {
	ability, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return TriggerCommand{}, err
	}

	monsterFormations, err := getResourcesDbItem(cfg, r, cfg.e.monsterFormations, ability, cfg.db.GetTriggerCommandMonsterFormationIDs)
	if err != nil {
		return TriggerCommand{}, err
	}

	users, err := getResourcesDbItem(cfg, r, cfg.e.characterClasses, ability, cfg.db.GetTriggerCommandCharClassIDs)
	if err != nil {
		return TriggerCommand{}, err
	}

	response := TriggerCommand{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.e.abilities, ability.Ability.ID),
		Description:        ability.Description,
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		UsedBy:             users,
		RelatedStats:       namesToNamedAPIResources(cfg, cfg.e.stats, ability.RelatedStats),
		Topmenu:            namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, ability.Topmenu, nil),
		Cursor:             ability.Cursor,
		MonsterFormations:  monsterFormations,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(cfg, r, i, response, qpnAbilityUser)
	if err != nil {
		return TriggerCommand{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrieveTriggerCommands(r *http.Request, i handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(nameIdQuery(r, i, ids, qpnUser, cfg.e.characterClasses.resTypeSing, cfg.l.CharClasses, cfg.db.GetTriggerCommandIDsByCharClass)),
		fidl(nameIdQuery(r, i, ids, qpnRelatedStat, cfg.e.stats.resTypeSing, cfg.l.Stats, cfg.db.GetTriggerCommandIDsByRelatedStat)),
	})
}
