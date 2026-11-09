package api

import (
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
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
		BaseStats: getAeonBaseStats(aeonLookup),
	}

	aeon, err := applyAeonStats(r, aeon, nil)
	if err != nil {
		return nil, err
	}

	aeonStatMap := handlerinput.GetResAmtTypeMap(aeon.BaseStats)
	newBaseStats := replaceBaseStats(mon.BaseStats, aeonStatMap, allowedStatIDs)

	return newBaseStats, nil
}
