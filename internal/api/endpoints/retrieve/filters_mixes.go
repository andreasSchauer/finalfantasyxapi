package retrieve

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func GetMixesByItem(r *http.Request, ctx context.Context, firstItemID int32) ([]int32, error) {
	i := apilookups.E.Mixes

	secondItemIDPtr, err := handlerinput.NullID(r, apilookups.E.Items, ts.QpnSecondItem, i.QueryLookup)
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
