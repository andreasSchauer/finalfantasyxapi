package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func RetrieveTreasures(r *http.Request, i handlerinput.Main[seedtypes.Treasure, apitypes.Treasure, apiresource.Unnamed]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationTreasureIDs),
		i.QueryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetSublocationTreasureIDs),
		i.QueryID(r, ids, ts.QpnArea, lookup.L.Areas, config.Cfg.DB.GetAreaTreasureIDs),
		i.QueryID(r, ids, ts.QpnItem, lookup.L.Items, config.Cfg.DB.GetTreasureIDsByItem),
		i.QueryJoined(r, ids, []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnEmptySlots, ts.QpnCharacter}, FilterTreasuresEquipment),
		i.QueryBool(r, ids, ts.QpnAnima, config.Cfg.DB.GetTreasureIDsByIsAnimaTreasure),
		i.QueryEnum(r, enums.T.LootType, ids, ts.QpnLootType, config.Cfg.DB.GetTreasureIDsByLootType),
		i.QueryEnum(r, enums.T.TreasureType, ids, ts.QpnTreasureType, config.Cfg.DB.GetTreasureIDsByTreasureType),
	})
}
