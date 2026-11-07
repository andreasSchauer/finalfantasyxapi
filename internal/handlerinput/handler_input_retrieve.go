package handlerinput

import (
	"net/http"
	"sync"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

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
