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
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, apiconvert.BattleInteraction),
	}

	err = getEnemyAbilityRelationships(r, ability, &response)
	if err != nil {
		return apitypes.EnemyAbility{}, err
	}
	
	return response, nil
}

func getEnemyAbilityRelationships(r *http.Request, ability seedtypes.EnemyAbility, response *apitypes.EnemyAbility) error {
	var err error

	response.Monsters, err = apilookups.E.Monsters.DBGetRelations(r.Context(), ability, config.Cfg.DB.GetEnemyAbilityMonsterIDs)
	if err != nil {
		return err
	}

	return nil
}
