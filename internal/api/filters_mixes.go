package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getMixesByItem(cfg *Config, r *http.Request, ctx context.Context, firstItemId int32) ([]int32, error) {
	i := cfg.e.mixes

	secondItemIdPtr, err := getQueryIdPtr(r, cfg.e.items, qpnSecondItem, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMixIDsByItems(ctx, database.GetMixIDsByItemsParams{
		FirstItemID:  firstItemId,
		SecondItemID: h.GetNullInt32(secondItemIdPtr),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by items.", i.resTypePlural), err)
	}

	return dbIDs, nil
}
