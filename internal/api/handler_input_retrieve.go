package api

import (
	"fmt"
	"net/http"
	"sync"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

// is this in the right file?

func (i HandlerInput[T, R, A]) HandleMultiple(r *http.Request, name string) (apiresource.APIResourceList[A], error) {
	dbIDs, err := i.getMultipleQuery(r.Context(), name)
	if err != nil {
		return apiresource.APIResourceList[A]{}, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get multiple %s with name '%s'", i.resTypeSingle, name), err)
	}

	return i.ToAPIResList(r, dbIDs)
}

func (i HandlerInput[T, R, A]) FilterIDs(r *http.Request, IDs []int32, queryFns []IDFilter) ([]int32, error) {
	filteredIDs := IDs
	g, ctx := errgroup.WithContext(r.Context())
	var mu sync.Mutex

	for _, fn := range queryFns {
		g.Go(func() error {
			dbIDs, err := fn(ctx)
			if err != nil {
				return err
			}

			mu.Lock()
			filteredIDs = ah.GetSharedIDs(filteredIDs, dbIDs)
			mu.Unlock()

			return nil
		})
	}

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	if i.avlFunc != nil {
		var err error
		filteredIDs, err = i.avlFunc(r, filteredIDs)
		if err != nil {
			return nil, err
		}
	}

	flip, err := i.queryLookup[ts.QpnFlip].ParseBool(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	if flip {
		filteredIDs = ah.RemoveIDs(IDs, filteredIDs)
	}

	return filteredIDs, nil
}
