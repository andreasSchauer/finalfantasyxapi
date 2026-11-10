package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func RetrieveSublocations(r *http.Request, i handlerinput.Main[seedtypes.Sublocation, apitypes.Sublocation, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationSublocationIDs),
		i.QueryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetSublocationIDsWithMonster),
		i.QueryIDWrapper(r, ids, ts.QpnItem, lookup.L.Items, GetSublocationsByItem),
		i.QueryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, config.Cfg.DB.GetSublocationIDsWithKeyItem),
		i.QueryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, config.Cfg.DB.GetSublocationIDsWithAutoAbility),
		i.QueryBool2(r, ids, ts.QpnCharacters, config.Cfg.DB.GetSublocationIDsWithCharacters),
		i.QueryBool2(r, ids, ts.QpnAeons, config.Cfg.DB.GetSublocationIDsWithAeons),
		i.QueryBool2(r, ids, ts.QpnMonsters, config.Cfg.DB.GetSublocationIDsWithMonsters),
		i.QueryBool2(r, ids, ts.QpnBossFights, config.Cfg.DB.GetSublocationIDsWithBosses),
		i.QueryBool2(r, ids, ts.QpnShops, config.Cfg.DB.GetSublocationIDsWithShops),
		i.QueryBool2(r, ids, ts.QpnTreasures, config.Cfg.DB.GetSublocationIDsWithTreasures),
		i.QueryBool2(r, ids, ts.QpnSidequests, config.Cfg.DB.GetSublocationIDsWithSidequests),
		i.QueryBool2(r, ids, ts.QpnFMVs, config.Cfg.DB.GetSublocationIDsWithFMVs),
	})
}
