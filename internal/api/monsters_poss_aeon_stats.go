package api

import (
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func applyPossessedAeonStats(srv *Server, r *http.Request, mon Monster) ([]BaseStat, error) {
	paramYunaStats := Q.monsters[ts.QpnYunaStats]
	if !slices.Contains(paramYunaStats.AllowedIDs, mon.ID) {
		return mon.BaseStats, nil
	}

	allowedStatIDs := []int32{1, 3, 4, 5, 6, 7, 9, 10}

	aeonLookup, _ := lookup.GetResource(mon.Name, lookup.L.Aeons)
	aeon := Aeon{
		ID:        aeonLookup.ID,
		Name:      aeonLookup.Name,
		BaseStats: getAeonBaseStats(srv, aeonLookup),
	}

	aeon, err := applyAeonStats(srv, r, aeon, nil)
	if err != nil {
		return nil, err
	}

	aeonStatMap := GetResAmtTypeMap(aeon.BaseStats)
	newBaseStats := replaceBaseStats(mon.BaseStats, aeonStatMap, allowedStatIDs)

	return newBaseStats, nil
}
