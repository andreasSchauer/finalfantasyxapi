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

func (cfg *Config) retrieveKeyItems(r *http.Request, i handlerInput[seeding.KeyItem, KeyItem, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(enumListQuery(cfg, r, i, cfg.t.KeyItemCategory, resources, "category", cfg.db.GetKeyItemIDsCategory)),
		frl(valueQuery(cfg, r, i, resources, "method", cfg.db.GetKeyItemIDsByMethod)),
		frl(idQuery(cfg, r, i, resources, "location", len(cfg.e.locations.objLookup), cfg.db.GetKeyItemIDsByLocation)),
		frl(idQuery(cfg, r, i, resources, "sublocation", len(cfg.e.sublocations.objLookup), cfg.db.GetKeyItemIDsBySublocation)),
		frl(idQuery(cfg, r, i, resources, "area", len(cfg.e.areas.objLookup), cfg.db.GetKeyItemIDsByArea)),
	})
}
