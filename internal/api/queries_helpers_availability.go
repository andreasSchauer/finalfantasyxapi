package api

import (
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type avlParams struct {
	availabilities []int32
	isRepeatable   sql.NullBool
	preAirship     bool
}

func checkAvl(cfg *Config, r *http.Request, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam) (avlParams, error) {
	queryParamAvl := queryLookup[QpnAvailability]

	availabilities, err := parseEnumListQuery(cfg, r, endpoint, queryParamAvl, cfg.T.AvailabilityType)
	if errExceptEmptyQuery(err) {
		return avlParams{}, err
	}
	if queryIsEmpty(err) {
		return avlParams{}, errEmptyQuery
	}

	var preAirship bool

	_, ok := queryLookup[QpnPreAirship]
	if ok {
		preAirship, err = parseBooleanQuery(r, queryLookup[QpnPreAirship])
		if errExceptEmptyQuery(err) {
			return avlParams{}, err
		}

		avlQuery := r.URL.Query().Get(string(queryParamAvl.Name))
		if avlQuery == string(database.AvailabilityTypePreAirship) {
			preAirship = true
		}
	}

	avlRanks := avlToRanks(availabilities, preAirship)

	params := avlParams{
		availabilities: h.SliceOrNil(avlRanks),
		preAirship:     preAirship,
	}

	return params, nil
}

func checkAvlAndRep(cfg *Config, r *http.Request, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam) (avlParams, error) {
	params, errAvl := checkAvl(cfg, r, endpoint, queryLookup)
	if errExceptEmptyQuery(errAvl) {
		return avlParams{}, errAvl
	}

	isRepeatable, errRepl := getQueryBoolPtr(r, QpnRepeatable, queryLookup)
	if errExceptEmptyQuery(errRepl) {
		return avlParams{}, errRepl
	}

	if queryIsEmpty(errAvl) && queryIsEmpty(errRepl) {
		return avlParams{}, errEmptyQuery
	}

	params.isRepeatable = h.GetNullBool(isRepeatable)

	return params, nil
}

func avlToRanks(avls []database.AvailabilityType, preAirship bool) []int32 {
	rankMap := map[database.AvailabilityType]int32{
		database.AvailabilityTypeAlways:    1,
		database.AvailabilityTypePost:      2,
		database.AvailabilityTypePreStory:  3,
		database.AvailabilityTypePostStory: 4,
	}
	if preAirship {
		rankMap[database.AvailabilityTypePreStory] = 2
		rankMap[database.AvailabilityTypePost] = 3
	}

	ranks := []int32{}

	for _, avl := range avls {
		ranks = append(ranks, rankMap[avl])
	}

	return ranks
}
