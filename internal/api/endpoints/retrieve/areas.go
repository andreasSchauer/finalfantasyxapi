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

func RetrieveAreas(r *http.Request, i handlerinput.Main[seedtypes.Area, apitypes.Area, apiresource.Area]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationAreaIDs),
		i.QueryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetSublocationAreaIDs),
		i.QueryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetAreaIDsWithMonster),
		i.QueryIDWrapper(r, ids, ts.QpnItem, lookup.L.Items, GetAreasByItem),
		i.QueryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, config.Cfg.DB.GetAreaIDsWithKeyItem),
		i.QueryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetAreasByMonster),
		i.QueryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, config.Cfg.DB.GetAreaIDsWithAutoAbility),
		i.QueryBool(r, ids, ts.QpnSaveSphere, config.Cfg.DB.GetAreaIDsWithSaveSphere),
		i.QueryBool(r, ids, ts.QpnCompSphere, config.Cfg.DB.GetAreaIDsWithCompSphere),
		i.QueryBool(r, ids, ts.QpnAirship, config.Cfg.DB.GetAreaIDsWithDropOff),
		i.QueryBool(r, ids, ts.QpnChocobo, config.Cfg.DB.GetAreaIDsChocobo),
		i.QueryBool2(r, ids, ts.QpnCharacters, config.Cfg.DB.GetAreaIDsWithCharacters),
		i.QueryBool2(r, ids, ts.QpnAeons, config.Cfg.DB.GetAreaIDsWithAeons),
		i.QueryBool2(r, ids, ts.QpnMonsters, config.Cfg.DB.GetAreaIDsWithMonsters),
		i.QueryBool2(r, ids, ts.QpnBossFights, config.Cfg.DB.GetAreaIDsWithBosses),
		i.QueryBool2(r, ids, ts.QpnShops, config.Cfg.DB.GetAreaIDsWithShops),
		i.QueryBool2(r, ids, ts.QpnTreasures, config.Cfg.DB.GetAreaIDsWithTreasures),
		i.QueryBool2(r, ids, ts.QpnSidequests, config.Cfg.DB.GetAreaIDsWithSidequests),
		i.QueryBool2(r, ids, ts.QpnFMVs, config.Cfg.DB.GetAreaIDsWithFMVs),
	})
}
