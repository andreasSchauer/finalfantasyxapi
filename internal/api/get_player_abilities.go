package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func (cfg *Config) getPlayerAbility(r *http.Request, i handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList], id int32) (PlayerAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return PlayerAbility{}, err
	}

	category, _ := newNamedAPIResourceFromType(cfg, cfg.e.playerAbilityCategory.endpoint, ability.Category, cfg.t.PlayerAbilityCategory)

	monsters, err := getResourcesDB(cfg, r, cfg.e.monsters, ability, cfg.db.GetPlayerAbilityMonsterIDs)
	if err != nil {
		return PlayerAbility{}, err
	}

	response := PlayerAbility{
		ID:          ability.ID,
		Name:        ability.Name,
		Version: 	ability.Version,	
		Description: ability.Description,
		Effect:      ability.Effect,
		Rank: ability.Rank,
		AppearsInHelpBar: ability.AppearsInHelpBar,
		CanCopycat: ability.CanCopycat,
		CanUseOutsideBattle: ability.CanUseOutsideBattle,
		MpCost: ability.MPCost,
		Category: category,
		AeonLearnItem: convertObjPtr(cfg, ability.AeonLearnItem, convertItemAmount),
		LearnedBy: namesToNamedAPIResources(cfg, cfg.e.characterClasses, ability.LearnedBy),
		RelatedStats: namesToNamedAPIResources(cfg, cfg.e.stats, ability.RelatedStats),
		StandardGridCharacter: namePtrToNamedAPIResPtr(cfg, cfg.e.characters, ability.StandardGridPos, nil),
		ExpertGridCharacter: namePtrToNamedAPIResPtr(cfg, cfg.e.characters, ability.ExpertGridPos, nil),
		Topmenu: namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, ability.Topmenu, nil),
		Submenu: namePtrToNamedAPIResPtr(cfg, cfg.e.submenus, ability.Submenu, nil),
		OpenSubmenu: namePtrToNamedAPIResPtr(cfg, cfg.e.submenus, ability.OpenSubmenu, nil),
		Cursor: ability.Cursor,
		Monsters: monsters,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}


func (cfg *Config) retrievePlayerAbilities(r *http.Request, i handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.PlayerAbilityCategory, resources, "category", cfg.db.GetPlayerAbilityIDsByCategory)),
		frl(boolQuery(cfg, r, i, resources, "phys_atk", cfg.db.GetPlayerAbilityIDsBasedOnPhysAttack)),
		frl(boolQuery2(cfg, r, i, resources, "darkable", cfg.db.GetPlayerAbilityIDsDarkable)),
		frl(idQuery(cfg, r, i, resources, "element", len(cfg.l.Elements), cfg.db.GetPlayerAbilityIDsByElement)),
		frl(typeQuery(cfg, r, i, cfg.t.DamageType, resources, "damage_type", cfg.db.GetPlayerAbilityIDsByDamageType)),
	})
}