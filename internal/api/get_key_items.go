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
		ID:                 keyItem.ID,
		Name:               keyItem.Name,
		UntypedItem: 		idToTypedAPIResource(cfg, cfg.e.masterItems, keyItem.MasterItem.ID),
		Category:           newNamedAPIResourceFromEnum(cfg, cfg.e.keyItemCategory.endpoint, keyItem.Category, cfg.t.KeyItemCategory),
		Description:        keyItem.Description,
		Effect:             keyItem.Effect,
		Primer: 			rel.Primer,
		CelestialWeapon: 	rel.CelestialWeapon,
		Treasures:          rel.Treasures,
		Quests:             rel.Quests,
		Areas: 				rel.Areas,
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
		frl(boolQuery2(cfg, r, i, resources, "treasure", cfg.db.GetKeyItemIDsTreasure)),
		frl(boolQuery2(cfg, r, i, resources, "quest", cfg.db.GetKeyItemIDsQuest)),
	})
}
