package api

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

// checks for emptiness of id-list-queryParam and converts its input into a slice of valid ids.
func (q QueryParam) parseIDList[F seedtypes.Lookupable](server *Server, r *http.Request, fLookup map[string]F) ([]int32, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return nil, err
	}

	return q.sliceIDs(server, r, query, fLookup)
}

// converts a list of unique query ids into a slice of valid ids.
func (q QueryParam) sliceIDs[F seedtypes.Lookupable](server *Server, r *http.Request, query string, fLookup map[string]F) ([]int32, error) {
	idSegments, err := queryListSplit(server, query)
	if err != nil {
		return nil, err
	}
	ids := []int32{}

	for _, segment := range idSegments {
		idsNew, err := q.checkIDRange(segment, fLookup)
		if err != nil {
			return nil, err
		}
		ids = slices.Concat(ids, idsNew)

		if len(ids) > server.Cfg.FetchLimit {
			return nil, newHTTPErrorFetchLimit(server.Cfg.FetchLimit)
		}
	}

	ids, err = cleanUpIntList(server, r, ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (q QueryParam) checkIDRange[F seedtypes.Lookupable](segment string, fLookup map[string]F) ([]int32, error) {
	idStrs := strings.Split(segment, "-")
	ids := []int32{}

	switch len(idStrs) {
	case 0:
		return nil, nil

	case 1:
		id, err := q.parseIDVal(idStrs[0], len(fLookup))
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)

	case 2:
		newIDs, err := q.sliceIDRange(idStrs, fLookup)
		if err != nil {
			return nil, err
		}
		ids = slices.Concat(ids, newIDs)

	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'", q.Name, segment, q.Usage), nil)
	}

	return ids, nil
}

func (q QueryParam) sliceIDRange[F seedtypes.Lookupable](idStrs []string, fLookup map[string]F) ([]int32, error) {
	minID, err := q.parseIDVal(idStrs[0], len(fLookup))
	if err != nil {
		return nil, err
	}

	maxID, err := q.parseIDVal(idStrs[1], len(fLookup))
	if err != nil {
		return nil, err
	}

	ids := sliceFromIntRange(minID, maxID)

	return ids, nil
}
