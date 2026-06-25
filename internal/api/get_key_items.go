package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getKeyItem(r *http.Request, i handlerInput[seeding.KeyItem, KeyItem, NamedAPIResource, NamedApiResourceList], id int32) (KeyItem, error) {
	keyItem, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return KeyItem{}, err
	}

	rel, err := getKeyItemRelationships(cfg, r, keyItem)
	if err != nil {
		return KeyItem{}, err
	}

	response := KeyItem{
		ID:              keyItem.ID,
		Name:            keyItem.Name,
		UntypedItem:     idToTypedAPIResource(cfg, cfg.e.allItems, keyItem.MasterItem.ID),
		Category:        enumToNamedAPIResource(cfg, cfg.e.keyItemCategory.endpoint, keyItem.Category, cfg.t.KeyItemCategory),
		Description:     keyItem.Description,
		Effect:          keyItem.Effect,
		Primer:          rel.Primer,
		CelestialWeapon: rel.CelestialWeapon,
		Treasures:       rel.Treasures,
		Quests:          rel.Quests,
		Areas:           rel.Areas,
	}

	return response, nil
}

func (cfg *Config) retrieveKeyItems(r *http.Request, i handlerInput[seeding.KeyItem, KeyItem, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.KeyItemCategory, ids, qpnCategory, cfg.db.GetKeyItemIDsCategory)),
		fidl(valueListQuery(cfg, r, i, ids, qpnMethods, cfg.db.GetKeyItemIDsByMethods)),
		fidl(idQuery(r, i, ids, qpnLocation, cfg.e.locations.objLookup, cfg.db.GetKeyItemIDsByLocation)),
		fidl(idQuery(r, i, ids, qpnSublocation, cfg.e.sublocations.objLookup, cfg.db.GetKeyItemIDsBySublocation)),
		fidl(idQuery(r, i, ids, qpnArea, cfg.e.areas.objLookup, cfg.db.GetKeyItemIDsByArea)),
	})
}
