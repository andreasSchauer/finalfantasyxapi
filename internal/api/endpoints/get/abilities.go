// Package get ajsasdijadsij
package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Ability(r *http.Request, i handlerinput.Main[seedtypes.Ability, apitypes.Ability, apiresource.Typed], id int32) (apitypes.Ability, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Ability{}, err
	}

	rel, err := getAbilityRelationships(r, ability)
	if err != nil {
		return apitypes.Ability{}, err
	}

	response := apitypes.Ability{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Type:               ability.Type,
		Rank:               ability.Rank,
		CanCopycat:         ability.CanCopycat,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		TypedAbility:       apiconvert.RefToNamedAPIResource(ability.GetAbilityRef()),
		Monsters:           rel.Monsters,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, apiconvert.BattleInteraction),
	}

	return response, nil
}

func getAbilityRelationships(r *http.Request, ability seedtypes.Ability) (apitypes.Ability, error) {
	var err error
	var rel apitypes.Ability

	rel.Monsters, err = apilookups.E.Monsters.DBGetRelations(r.Context(), ability, config.Cfg.DB.GetAbilityMonsterIDs)
	if err != nil {
		return apitypes.Ability{}, err
	}

	return rel, nil
}
