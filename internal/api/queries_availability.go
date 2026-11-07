package api

import (
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type avlParams struct {
	availabilities []int32
	isRepeatable   sql.NullBool
	preAirship     bool
}

func (i HandlerInput[T, R, A]) QueryAvlGetParams(r *http.Request) (avlParams, error) {
	queryParamAvl := i.queryLookup[ts.QpnAvailability]

	availabilities, err := queryParamAvl.ParseEnumList(r, enums.T.AvailabilityType)
	if errs.ErrExceptEmptyQuery(err) {
		return avlParams{}, err
	}
	if errs.QueryIsEmpty(err) {
		return avlParams{}, errs.ErrEmptyQuery
	}

	var preAirship bool

	_, ok := i.queryLookup[ts.QpnPreAirship]
	if ok {
		preAirship, err = i.queryLookup[ts.QpnPreAirship].ParseBool(r)
		if errs.ErrExceptEmptyQuery(err) {
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

// what's the difference?

func (i HandlerInput[T, R, A]) QueryAvlRepGetParams(r *http.Request) (avlParams, error) {
	params, errAvl := i.QueryAvlGetParams(r)
	if errs.ErrExceptEmptyQuery(errAvl) {
		return avlParams{}, errAvl
	}

	isRepeatable, errRepl := nullBool(r, ts.QpnRepeatable, i.queryLookup)
	if errs.ErrExceptEmptyQuery(errRepl) {
		return avlParams{}, errRepl
	}

	if errs.QueryIsEmpty(errAvl) && errs.QueryIsEmpty(errRepl) {
		return avlParams{}, errs.ErrEmptyQuery
	}

	params.isRepeatable = isRepeatable

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
