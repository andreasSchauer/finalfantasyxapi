package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func RetrieveMiscAbilities(r *http.Request, i handlerinput.Main[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeMiscAbility

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryIntList(r, ids, ts.QpnRank, getTypedAbilityIDsByRank(abilityType)),
		i.QueryNameID(r, ids, ts.QpnUser, apilookups.E.CharacterClasses.ResTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetMiscAbilityIDsByCharClass),
		i.QueryBool(r, ids, ts.QpnCopycat, getTypedAbilityIDsByCanCopycat(abilityType)),
		i.QueryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(abilityType)),
		i.QueryBool2(r, ids, ts.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(abilityType)),
	})
}
