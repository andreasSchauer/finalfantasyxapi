package api

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func (i handlerInput[T, R, A]) handleMultiple(cfg *Config, r *http.Request, name string) (APIResourceList[A], error) {
	dbIDs, err := i.getMultipleQuery(r.Context(), name)
	if err != nil {
		return APIResourceList[A]{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get multiple %s with name '%s'", i.resTypeSingle, name), err)
	}

	return i.toResList(cfg, r, dbIDs)
}

func (i handlerInput[T, R, A]) filterIDs(cfg *Config, r *http.Request, IDs []int32, queryFns []IDFilter) ([]int32, error) {
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
			filteredIDs = getSharedIDs(filteredIDs, dbIDs)
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
		filteredIDs, err = i.avlFunc(cfg, r, filteredIDs)
		if err != nil {
			return nil, err
		}
	}

	flip, err := i.queryLookup[typestrings.QpnFlip].parseBool(r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	if flip {
		filteredIDs = removeIDs(IDs, filteredIDs)
	}

	return filteredIDs, nil
}
