package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMixesByItem(cfg *Config, r *http.Request, ctx context.Context, firstItemID int32) ([]int32, error) {
	i := cfg.E.Mixes

	secondItemIDPtr, err := nullID(r, cfg.E.Items, ts.QpnSecondItem, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMixIDsByItems(ctx, database.GetMixIDsByItemsParams{
		FirstItemID:  firstItemID,
		SecondItemID: secondItemIDPtr,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by items", i.resTypePlural), err)
	}

	return dbIDs, nil
}
