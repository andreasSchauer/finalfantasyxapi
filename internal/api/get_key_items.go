package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getKeyItem(r *http.Request, i handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource], id int32) (KeyItem, error) {
	keyItem, err := i.queryParamsVerifyGet(r, id)
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

func (cfg *Config) retrieveKeyItems(r *http.Request, i handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, enums.T.KeyItemCategory, ids, typestrings.QpnCategory, cfg.db.GetKeyItemIDsCategory),
		i.queryValueList(cfg, r, ids, typestrings.QpnMethods, cfg.db.GetKeyItemIDsByMethods),
		i.queryID(r, ids, typestrings.QpnLocation, cfg.E.Locations.objLookup, cfg.db.GetKeyItemIDsByLocation),
		i.queryID(r, ids, typestrings.QpnSublocation, cfg.E.Sublocations.objLookup, cfg.db.GetKeyItemIDsBySublocation),
		i.queryID(r, ids, typestrings.QpnArea, cfg.E.Areas.objLookup, cfg.db.GetKeyItemIDsByArea),
	})
}
