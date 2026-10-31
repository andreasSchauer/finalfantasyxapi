package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getKeyItem(r *http.Request, i handlerInput[seeding.KeyItem, KeyItem, NamedAPIResource, NamedAPIResourceList], id int32) (KeyItem, error) {
	keyItem, err := verifyParamsAndGet(r, i, id)
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
		UntypedItem:     idToTypedAPIResource(cfg, cfg.E.AllItems, keyItem.MasterItem.ID),
		Category:        keyItem.Category,
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

func (cfg *Config) retrieveKeyItems(r *http.Request, i handlerInput[seeding.KeyItem, KeyItem, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		enumListQuery(cfg, r, i, cfg.T.KeyItemCategory, ids, QpnCategory, cfg.db.GetKeyItemIDsCategory),
		valueListQuery(cfg, r, i, ids, QpnMethods, cfg.db.GetKeyItemIDsByMethods),
		idQuery(r, i, ids, QpnLocation, cfg.E.Locations.objLookup, cfg.db.GetKeyItemIDsByLocation),
		idQuery(r, i, ids, QpnSublocation, cfg.E.Sublocations.objLookup, cfg.db.GetKeyItemIDsBySublocation),
		idQuery(r, i, ids, QpnArea, cfg.E.Areas.objLookup, cfg.db.GetKeyItemIDsByArea),
	})
}
