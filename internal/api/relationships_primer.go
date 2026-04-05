package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getPrimerRelationships(cfg *Config, r *http.Request, primer seeding.Primer) (Primer, error) {
	treasures, err := getResourcesDbItem(cfg, r, cfg.e.treasures, primer, cfg.db.GetPrimerTreasureIDs)
	if err != nil {
		return Primer{}, err
	}

	areas, err := getResourcesDbItem(cfg, r, cfg.e.areas, primer, cfg.db.GetPrimerAreaIDs)
	if err != nil {
		return Primer{}, err
	}

	rel := Primer{
		Treasures:  treasures,
		Areas:		areas,
	}

	return rel, nil
}
