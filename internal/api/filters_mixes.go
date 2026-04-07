package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func getMixesByItem(cfg *Config, r *http.Request, firstItemId int32) ([]NamedAPIResource, error) {
	i := cfg.e.mixes

	secondItemIdPtr, err := getQueryIdPtr(r, cfg.e.items, "second_item", i.queryLookup)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMixIDsByItems(r.Context(), database.GetMixIDsByItemsParams{
		FirstItemID:  firstItemId,
		SecondItemID: h.GetNullInt32(secondItemIdPtr),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by items.", i.resourceType), err)
	}
	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}