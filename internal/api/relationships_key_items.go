package api

import (
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getKeyItemRelationships(cfg *Config, r *http.Request, keyItem seeding.KeyItem) (KeyItem, error) {
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.keyItems, keyItem.ID)
	if err != nil {
		return KeyItem{}, err
	}

	treasures, err := runRelAvailabilityQuery(cfg, r, cfg.e.treasures, keyItem, availabilityParams, getKeyItemSourceIDs(cfg, ViewSourceTypeTreasure))
	if err != nil {
		return KeyItem{}, err
	}

	quests, err := runRelAvailabilityQuery(cfg, r, cfg.e.quests, keyItem, availabilityParams, getKeyItemSourceIDs(cfg, ViewSourceTypeQuest))
	if err != nil {
		return KeyItem{}, err
	}

	areas, err := runRelAvailabilityQuery(cfg, r, cfg.e.areas, keyItem, availabilityParams, getKeyItemAreaIDs(cfg))
	if err != nil {
		return KeyItem{}, err
	}

	var primer *NamedAPIResource
	if keyItem.Category == string(database.KeyItemCategoryPrimer) {
		primerRes := nameToNamedAPIResource(cfg, cfg.e.primers, keyItem.Name, nil)
		primer = &primerRes
	}

	celestialWeapon, err := getKeyItemCelestialWeapon(cfg, r, keyItem)
	if err != nil {
		return KeyItem{}, err
	}

	rel := KeyItem{
		Primer:          primer,
		CelestialWeapon: celestialWeapon,
		Treasures:       treasures,
		Quests:          quests,
		Areas:           areas,
	}

	return rel, nil
}

func getKeyItemCelestialWeapon(cfg *Config, r *http.Request, keyItem seeding.KeyItem) (*NamedAPIResource, error) {
	if !(strings.HasSuffix(keyItem.Name, "crest") || strings.HasSuffix(keyItem.Name, "sigil")) {
		return nil, nil
	}

	keyItemBase := strings.Split(keyItem.Name, " ")[0]

	celestialID, err := cfg.db.GetKeyItemCelestialWeapon(r.Context(), database.KeyItemBase(keyItemBase))
	if err != nil {
		return nil, newHTTPErrorDbOne(cfg.e.celestialWeapons.resourceType, keyItem, err)
	}

	celestialRes := idToNamedAPIResource(cfg, cfg.e.celestialWeapons, celestialID)
	return &celestialRes, nil
}
