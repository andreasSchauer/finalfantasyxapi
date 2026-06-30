package api

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getMultipleAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], name string) (L, error) {
	var zeroType L

	dbIDs, err := i.getMultipleQuery(r.Context(), name)
	if err != nil {
		return zeroType, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get multiple %s with name '%s'.", i.resTypeSingle, name), err)
	}

	return idsToAPIResourceList(cfg, r, i, dbIDs)
}

func filterIDs[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], IDs []int32, queryFns []IdFilter) ([]int32, error) {
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

	flip, err := parseBooleanQuery(r, i.queryLookup[qpnFlip])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	if flip {
		filteredIDs = removeIDs(IDs, filteredIDs)
	}

	return filteredIDs, nil
}
