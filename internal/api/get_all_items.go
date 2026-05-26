package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMasterItem(r *http.Request, i handlerInput[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList], id int32) (MasterItem, error) {
	masterItem, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return MasterItem{}, err
	}

	obtainable, err := getMasterItemObtainable(cfg, r, masterItem)
	if err != nil {
		return MasterItem{}, err
	}

	response := MasterItem{
		ID:             masterItem.ID,
		Name:           masterItem.Name,
		Type:           enumToNamedAPIResource(cfg, cfg.e.itemType.endpoint, string(masterItem.Type), cfg.t.ItemType),
		ObtainableFrom: obtainable,
	}

	switch masterItem.Type {
	case database.ItemTypeItem:
		item, _ := seeding.GetResource(masterItem, cfg.l.Items)

		response.Description = item.Description
		response.Effect = item.Effect
		response.TypedItem = nameToNamedAPIResource(cfg, cfg.e.items, item.Name, nil)

	case database.ItemTypeKeyItem:
		keyItem, _ := seeding.GetResource(masterItem, cfg.l.KeyItems)

		response.Description = keyItem.Description
		response.Effect = keyItem.Effect
		response.TypedItem = nameToNamedAPIResource(cfg, cfg.e.keyItems, keyItem.Name, nil)
	}

	return response, nil
}

func (cfg *Config) retrieveMasterItems(r *http.Request, i handlerInput[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList]) (TypedAPIResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return TypedAPIResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[TypedAPIResource]{
		frl(enumListQuery(cfg, r, i, cfg.t.ItemType, resources, "type", cfg.db.GetMasterItemIDsByType)),
		frl(valueQuery(cfg, r, i, resources, "method", cfg.db.GetMasterItemIDsByMethod)),
		frl(idQueryWrapper(cfg, r, i, resources, "location", len(cfg.e.locations.objLookup), getMasterItemsByLocation)),
		frl(idQueryWrapper(cfg, r, i, resources, "sublocation", len(cfg.e.sublocations.objLookup), getMasterItemsBySublocation)),
		frl(idQueryWrapper(cfg, r, i, resources, "area", len(cfg.e.areas.objLookup), getMasterItemsByArea)),
	})
}
