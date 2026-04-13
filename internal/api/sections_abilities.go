package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AbilitySimple struct {
	ID                 int32                     `json:"id"`
	URL                string                    `json:"url"`
	Name               string                    `json:"name"`
	Version            *int32                    `json:"version,omitempty"`
	Specification      *string                   `json:"specification,omitempty"`
	Type               *string                   `json:"type,omitempty"`
	Rank               *int32                    `json:"rank"`
	BattleInteractions []BattleInteractionSimple `json:"battle_interactions"`
}

func (a AbilitySimple) GetURL() string {
	return a.URL
}

type PlayerAbilitySimple struct {
	ID                 int32                     `json:"id"`
	URL                string                    `json:"url"`
	Name               string                    `json:"name"`
	Version            *int32                    `json:"version,omitempty"`
	Specification      *string                   `json:"specification,omitempty"`
	Type               *string                   `json:"type,omitempty"`
	Rank               *int32                    `json:"rank"`
	MpCost             int32                     `json:"mp_cost"`
	BattleInteractions []BattleInteractionSimple `json:"battle_interactions"`
}

func (a PlayerAbilitySimple) GetURL() string {
	return a.URL
}

func createAbilitySimple(cfg *Config, r *http.Request, id int32, subsection Subsection) (SimpleResource, error) {
	i := cfg.e.abilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)
	typeStr := string(ability.Type)
	rank := ability.Rank

	if ability.Type == database.AbilityTypeOverdriveAbility {
		fetchedRank := subsection.relations[id][RelationRanks][0]
		rank = &fetchedRank
	}

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                createResourceURL(cfg, i.endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               rank,
		Type:               &typeStr,
		BattleInteractions: getAbilityBattleInteractionsSimple(cfg, ability),
	}

	return abilitySimple, nil
}

func getAbilitySectionRelations(cfg *Config, r *http.Request, abilityIDs []int32) (map[int32]map[Relation][]int32, error) {
	i := cfg.e.abilities
	relations := make(map[int32]map[Relation][]int32)

	abilityJunctions, err := getDbJunctions(r, abilityIDs, i.resourceType, "rank", cfg.db.GetAbilityIdRankPairs, juncAbilityRank)
	if err != nil {
		return nil, err
	}

	for _, abilityID := range abilityIDs {
		relationMap := make(map[Relation][]int32)

		relationMap[RelationRanks], abilityJunctions = getJunctionIDs(abilityID, abilityJunctions)

		relations[abilityID] = relationMap
	}

	return relations, nil
}

func createEnemyAbilitySimple(cfg *Config, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := cfg.e.enemyAbilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                createResourceURL(cfg, i.endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createItemAbilitySimple(cfg *Config, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := cfg.e.itemAbilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                createResourceURL(cfg, i.endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createOverdriveAbilitySimple(cfg *Config, r *http.Request, id int32, subsection Subsection) (SimpleResource, error) {
	i := cfg.e.overdriveAbilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	rank := subsection.relations[id][RelationRanks][0]

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                createResourceURL(cfg, i.endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               &rank,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func getOverdriveAbilitySectionRelations(cfg *Config, r *http.Request, abilityIDs []int32) (map[int32]map[Relation][]int32, error) {
	i := cfg.e.overdriveAbilities
	relations := make(map[int32]map[Relation][]int32)

	abilityJunctions, err := getDbJunctions(r, abilityIDs, i.resourceType, "rank", cfg.db.GetOverdriveAbilityIdRankPairs, juncOverdriveAbilityRank)
	if err != nil {
		return nil, err
	}

	for _, abilityID := range abilityIDs {
		relationMap := make(map[Relation][]int32)

		relationMap[RelationRanks], abilityJunctions = getJunctionIDs(abilityID, abilityJunctions)

		relations[abilityID] = relationMap
	}

	return relations, nil
}

func createPlayerAbilitySimple(cfg *Config, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := cfg.e.playerAbilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	playerAbilitySimple := PlayerAbilitySimple{
		ID:                 ability.ID,
		URL:                createResourceURL(cfg, i.endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		MpCost:             ability.MPCost,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return playerAbilitySimple, nil
}

func createTriggerCommandSimple(cfg *Config, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := cfg.e.triggerCommands
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                createResourceURL(cfg, i.endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createUnspecifiedAbilitySimple(cfg *Config, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := cfg.e.unspecifiedAbilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                createResourceURL(cfg, i.endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}
