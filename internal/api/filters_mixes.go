package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMixesByItem(r *http.Request, ctx context.Context, firstItemID int32) ([]int32, error) {
	i := E.Mixes

	secondItemIDPtr, err := NullID(r, E.Items, ts.QpnSecondItem, i.queryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetMixIDsByItems(ctx, database.GetMixIDsByItemsParams{
		FirstItemID:  firstItemID,
		SecondItemID: secondItemIDPtr,
	})
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by items", i.resTypePlural), err)
	}

	return dbIDs, nil
}
