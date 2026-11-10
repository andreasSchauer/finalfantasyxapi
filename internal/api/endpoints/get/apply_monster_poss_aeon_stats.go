package get

import (
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func applyPossessedAeonStats(r *http.Request, mon apitypes.Monster) ([]apitypes.BaseStat, error) {
	paramYunaStats := apilookups.Q.Monsters[ts.QpnYunaStats]
	if !slices.Contains(paramYunaStats.AllowedIDs, mon.ID) {
		return mon.BaseStats, nil
	}

	allowedStatIDs := []int32{1, 3, 4, 5, 6, 7, 9, 10}

	aeonLookup, _ := lookup.GetResource(mon.Name, lookup.L.Aeons)
	aeon := apitypes.Aeon{
		ID:        aeonLookup.ID,
		Name:      aeonLookup.Name,
		BaseStats: th.GetAeonBaseStats(aeonLookup),
	}

	aeon, err := applyAeonStats(r, aeon, nil)
	if err != nil {
		return nil, err
	}

	aeonStatMap := handlerinput.GetResAmtTypeMap(aeon.BaseStats)
	newBaseStats := th.ReplaceBaseStats(mon.BaseStats, aeonStatMap, allowedStatIDs)

	return newBaseStats, nil
}
