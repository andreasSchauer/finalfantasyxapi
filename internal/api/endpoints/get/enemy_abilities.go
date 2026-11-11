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

func EnemyAbility(r *http.Request, i handlerinput.Main[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.Named], id int32) (apitypes.EnemyAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.EnemyAbility{}, err
	}

	rel, err := getEnemyAbilityRelationships(r, ability)
	if err != nil {
		return apitypes.EnemyAbility{}, err
	}

	response := apitypes.EnemyAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     handlerinput.IDToTypedAPIResource(apilookups.E.Abilities, ability.Ability.ID),
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Monsters:           rel.Monsters,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, apiconvert.BattleInteraction),
	}

	return response, nil
}

func getEnemyAbilityRelationships(r *http.Request, ability seedtypes.EnemyAbility) (apitypes.EnemyAbility, error) {
	var err error
	var rel apitypes.EnemyAbility

	rel.Monsters, err = apilookups.E.Monsters.DBGetRelations(r.Context(), ability, config.Cfg.DB.GetEnemyAbilityMonsterIDs)
	if err != nil {
		return apitypes.EnemyAbility{}, err
	}

	return rel, nil
}
