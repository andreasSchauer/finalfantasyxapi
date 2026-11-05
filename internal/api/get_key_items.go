package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getKeyItem(r *http.Request, i handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource], id int32) (KeyItem, error) {
	keyItem, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return KeyItem{}, err
	}

	rel, err := getKeyItemRelationships(r, keyItem)
	if err != nil {
		return KeyItem{}, err
	}

	response := KeyItem{
		ID:              keyItem.ID,
		Name:            keyItem.Name,
		UntypedItem:     idToTypedAPIResource(E.AllItems, keyItem.MasterItem.ID),
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

func retrieveKeyItems(r *http.Request, i handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnumList(r, enums.T.KeyItemCategory, ids, ts.QpnCategory, config.Cfg.DB.GetKeyItemIDsCategory),
		i.queryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetKeyItemIDsByMethods),
		i.queryID(r, ids, ts.QpnLocation, E.Locations.objLookup, config.Cfg.DB.GetKeyItemIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, E.Sublocations.objLookup, config.Cfg.DB.GetKeyItemIDsBySublocation),
		i.queryID(r, ids, ts.QpnArea, E.Areas.objLookup, config.Cfg.DB.GetKeyItemIDsByArea),
	})
}
