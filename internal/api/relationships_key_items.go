package api

import (
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getKeyItemRelationships(cfg *Config, r *http.Request, keyItem seeding.KeyItem) (KeyItem, error) {
	availabilityParams, err := getAvailabilityParams(cfg, r, cfg.e.keyItems, keyItem)
	if err != nil {
		return KeyItem{}, err
	}

	treasures, err := runAvailabilityQuery(cfg, r, cfg.e.treasures, keyItem, availabilityParams, convGetKeyItemTreasureIDs(cfg))
	if err != nil {
		return KeyItem{}, err
	}

	quests, err := runAvailabilityQuery(cfg, r, cfg.e.quests, keyItem, availabilityParams, convGetKeyItemQuestIDs(cfg))
	if err != nil {
		return KeyItem{}, err
	}

	areas, err := getResourcesDbItem(cfg, r, cfg.e.areas, keyItem, cfg.db.GetKeyItemAreaIDs)
	if err != nil {
		return KeyItem{}, err
	}

	rel := KeyItem{
		Treasures:  treasures,
		Quests:  	quests,
		Areas:		areas,
	}

	if keyItem.Category == string(database.KeyItemCategoryPrimer) {
		primerRes := nameToNamedAPIResource(cfg, cfg.e.primers, keyItem.Name, nil)
		rel.Primer = &primerRes
	}

	if strings.HasSuffix(keyItem.Name, "crest") || strings.HasSuffix(keyItem.Name, "sigil") {
		keyItemBase := strings.Split(keyItem.Name, " ")[0]

		celestialID, err := cfg.db.GetKeyItemCelestialWeapon(r.Context(), database.KeyItemBase(keyItemBase))
		if err != nil {
			return KeyItem{}, newHTTPErrorDbOne(cfg.e.celestialWeapons.resourceType, keyItem, err)
		}

		celestialRes := idToNamedAPIResource(cfg, cfg.e.celestialWeapons, celestialID)
		rel.CelestialWeapon = &celestialRes
	}

	return rel, nil
}
