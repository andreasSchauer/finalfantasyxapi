package api

import (
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func applyPossessedAeonStats(cfg *Config, r *http.Request, mon Monster) ([]BaseStat, error) {
	paramYunaStats := cfg.q.monsters[qpnYunaStats]
	if !slices.Contains(paramYunaStats.AllowedIDs, mon.ID) {
		return mon.BaseStats, nil
	}

	allowedStatIDs := []int32{1, 3, 4, 5, 6, 7, 9, 10}

	aeonLookup, _ := seeding.GetResource(mon.Name, cfg.l.Aeons)
	aeon := Aeon{
		ID:        aeonLookup.ID,
		Name:      aeonLookup.Name,
		BaseStats: getAeonBaseStats(cfg, aeonLookup),
	}

	aeon, err := applyAeonStats(cfg, r, aeon, nil)
	if err != nil {
		return nil, err
	}

	aeonStatMap := getResAmtTypeMap(aeon.BaseStats)
	newBaseStats := replaceBaseStats(mon.BaseStats, aeonStatMap, allowedStatIDs)

	return newBaseStats, nil
}
