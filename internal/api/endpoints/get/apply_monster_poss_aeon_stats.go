package get

import (
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func applyPossessedAeonStats(r *http.Request, mon *apitypes.Monster) error {
	qpYunaStats := apilookups.Q.Monsters[ts.QpnYunaStats]
	if !slices.Contains(qpYunaStats.AllowedIDs, mon.ID) {
		return nil
	}

	allowedStatIDs := []int32{1, 3, 4, 5, 6, 7, 9, 10}

	aeonLookup, _ := lookup.GetResource(mon.Name, lookup.L.Aeons)
	aeon := apitypes.Aeon{
		ID:        aeonLookup.ID,
		Name:      aeonLookup.Name,
		BaseStats: th.GetAeonBaseStats(aeonLookup),
	}

	err := applyAeonStats(r, &aeon, nil)
	if err != nil {
		return err
	}

	aeonStatMap := apiresource.GetResAmtTypeMap(aeon.BaseStats)
	mon.BaseStats = th.ReplaceBaseStats(mon.BaseStats, aeonStatMap, allowedStatIDs)

	return nil
}
