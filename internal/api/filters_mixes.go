package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMixesByItem(r *http.Request, ctx context.Context, firstItemID int32) ([]int32, error) {
	i := E.Mixes

	secondItemIDPtr, err := handlerinput.NullID(r, E.Items, ts.QpnSecondItem, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetMixIDsByItems(ctx, database.GetMixIDsByItemsParams{
		FirstItemID:  firstItemID,
		SecondItemID: secondItemIDPtr,
	})
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by items", i.ResTypePlural), err)
	}

	return dbIDs, nil
}
