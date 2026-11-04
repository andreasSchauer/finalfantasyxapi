package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
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

func createAbilitySimple(server *Server, r *http.Request, id int32, subsection Subsection) (SimpleResource, error) {
	i := E.Abilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)
	typeStr := string(ability.Type)

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                CreateResourceURL(server, i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		Type:               &typeStr,
		BattleInteractions: convertObjSlice(server, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createEnemyAbilitySimple(server *Server, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := E.EnemyAbilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                CreateResourceURL(server, i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: convertObjSlice(server, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createItemAbilitySimple(server *Server, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := E.ItemAbilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                CreateResourceURL(server, i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: convertObjSlice(server, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createOverdriveAbilitySimple(server *Server, r *http.Request, id int32, subsection Subsection) (SimpleResource, error) {
	i := E.OverdriveAbilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                CreateResourceURL(server, i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: convertObjSlice(server, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createPlayerAbilitySimple(server *Server, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := E.PlayerAbilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	playerAbilitySimple := PlayerAbilitySimple{
		ID:                 ability.ID,
		URL:                CreateResourceURL(server, i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		MpCost:             ability.MPCost,
		BattleInteractions: convertObjSlice(server, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return playerAbilitySimple, nil
}

func createTriggerCommandSimple(server *Server, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := E.TriggerCommands
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                CreateResourceURL(server, i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: convertObjSlice(server, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createMiscAbilitySimple(server *Server, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := E.MiscAbilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID:                 ability.ID,
		URL:                CreateResourceURL(server, i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: convertObjSlice(server, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}
