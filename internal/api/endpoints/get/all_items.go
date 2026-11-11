package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func MasterItem(r *http.Request, i handlerinput.Main[seedtypes.MasterItem, apitypes.MasterItem, apiresource.Typed], id int32) (apitypes.MasterItem, error) {
	masterItem, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.MasterItem{}, err
	}

	rel, err := getMasterItemRelationships(r, masterItem)
	if err != nil {
		return apitypes.MasterItem{}, err
	}

	response := apitypes.MasterItem{
		ID:             masterItem.ID,
		Name:           masterItem.Name,
		Type:           masterItem.Type,
		ObtainableFrom: rel.ObtainableFrom,
	}

	return completeMasterItemResponse(r, response), nil
}

func getMasterItemRelationships(r *http.Request, masterItem seedtypes.MasterItem) (apitypes.MasterItem, error) {
	var rel apitypes.MasterItem
	var err error

	rel.ObtainableFrom, err = getMasterItemObtainable(r, masterItem)
	if err != nil {
		return apitypes.MasterItem{}, err
	}

	return rel, nil
}

func getMasterItemObtainable(r *http.Request, masterItem seedtypes.MasterItem) (apitypes.ObtainableFrom, error) {
	availabilityParams, err := apilookups.E.AllItems.QueryAvlRelGetParams(r, masterItem.ID)
	if err != nil {
		return apitypes.ObtainableFrom{}, err
	}

	bools, err := config.Cfg.DB.GetMasterItemObtainableBools(r.Context(), database.GetMasterItemObtainableBoolsParams{
		MasterItemID: availabilityParams.ParentID,
		Availability: availabilityParams.Availability,
		Repeatable:   availabilityParams.Repeatable,
	})
	if err != nil {
		return apitypes.ObtainableFrom{}, err
	}

	obtainable := apitypes.ObtainableFrom{
		Monsters:        bools.Monsters,
		Treasures:       bools.Treasures,
		Shops:           bools.Shops,
		Quests:          bools.Quests,
		BlitzballPrizes: bools.Blitzball,
	}

	return obtainable, nil
}

func completeMasterItemResponse(r *http.Request, masterItem apitypes.MasterItem) apitypes.MasterItem {
	switch masterItem.Type {
	case database.ItemTypeItem:
		item, _ := lookup.GetResource(masterItem.Name, lookup.L.Items)

		masterItem.Description = item.Description
		masterItem.Effect = item.Effect
		masterItem.TypedItem = handlerinput.NameToNamedAPIResource(apilookups.E.Items, item.Name, nil)

	case database.ItemTypeKeyItem:
		keyItem, _ := lookup.GetResource(masterItem.Name, lookup.L.KeyItems)

		masterItem.Description = keyItem.Description
		masterItem.Effect = keyItem.Effect
		masterItem.TypedItem = handlerinput.NameToNamedAPIResource(apilookups.E.KeyItems, keyItem.Name, nil)
	}

	return masterItem
}
