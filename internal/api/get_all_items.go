package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getMasterItem(r *http.Request, i handlerInput[seedtypes.MasterItem, MasterItem, TypedAPIResource], id int32) (MasterItem, error) {
	masterItem, err := i.queryParamsVerifyGet(r, id)
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
		item, _ := lookup.GetResource(masterItem, lookup.L.Items)

		response.Description = item.Description
		response.Effect = item.Effect
		response.TypedItem = nameToNamedAPIResource(cfg, cfg.E.Items, item.Name, nil)

	case database.ItemTypeKeyItem:
		keyItem, _ := lookup.GetResource(masterItem, lookup.L.KeyItems)

		response.Description = keyItem.Description
		response.Effect = keyItem.Effect
		response.TypedItem = nameToNamedAPIResource(cfg, cfg.E.KeyItems, keyItem.Name, nil)
	}

	return response, nil
}

func (cfg *Config) retrieveMasterItems(r *http.Request, i handlerInput[seedtypes.MasterItem, MasterItem, TypedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, enums.T.ItemType, ids, typestrings.QpnType, cfg.db.GetMasterItemIDsByType),
		i.queryValueList(cfg, r, ids, typestrings.QpnMethods, cfg.db.GetMasterItemIDsByMethods),
		i.queryIDWrapper(cfg, r, ids, typestrings.QpnLocation, cfg.E.Locations.objLookup, getMasterItemsByLocation),
		i.queryIDWrapper(cfg, r, ids, typestrings.QpnSublocation, cfg.E.Sublocations.objLookup, getMasterItemsBySublocation),
		i.queryIDWrapper(cfg, r, ids, typestrings.QpnArea, cfg.E.Areas.objLookup, getMasterItemsByArea),
	})
}
