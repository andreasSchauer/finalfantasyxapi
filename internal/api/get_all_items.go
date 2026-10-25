package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMasterItem(r *http.Request, i handlerInput[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList], id int32) (MasterItem, error) {
	masterItem, err := verifyParamsAndGet(r, i, id)
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
		Type:           masterItem.Type,
		ObtainableFrom: obtainable,
	}

	switch masterItem.Type {
	case database.ItemTypeItem:
		item, _ := seeding.GetResource(masterItem, cfg.L.Items)

		response.Description = item.Description
		response.Effect = item.Effect
		response.TypedItem = nameToNamedAPIResource(cfg, cfg.E.Items, item.Name, nil)

	case database.ItemTypeKeyItem:
		keyItem, _ := seeding.GetResource(masterItem, cfg.L.KeyItems)

		response.Description = keyItem.Description
		response.Effect = keyItem.Effect
		response.TypedItem = nameToNamedAPIResource(cfg, cfg.E.KeyItems, keyItem.Name, nil)
	}

	return response, nil
}

func (cfg *Config) retrieveMasterItems(r *http.Request, i handlerInput[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		enumListQuery(cfg, r, i, cfg.T.ItemType, ids, QpnType, cfg.db.GetMasterItemIDsByType),
		valueListQuery(cfg, r, i, ids, QpnMethods, cfg.db.GetMasterItemIDsByMethods),
		idQueryWrapper(cfg, r, i, ids, QpnLocation, cfg.E.Locations.objLookup, getMasterItemsByLocation),
		idQueryWrapper(cfg, r, i, ids, QpnSublocation, cfg.E.Sublocations.objLookup, getMasterItemsBySublocation),
		idQueryWrapper(cfg, r, i, ids, QpnArea, cfg.E.Areas.objLookup, getMasterItemsByArea),
	})
}
