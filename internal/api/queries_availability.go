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

func (i handlerInput[T, R, A]) queryAvlGetParams(r *http.Request) (avlParams, error) {
	queryParamAvl := i.queryLookup[ts.QpnAvailability]

	availabilities, err := queryParamAvl.ParseEnumList(r, enums.T.AvailabilityType)
	if errExceptEmptyQuery(err) {
		return avlParams{}, err
	}
	if queryIsEmpty(err) {
		return avlParams{}, errs.ErrEmptyQuery
	}

	var preAirship bool

	_, ok := i.queryLookup[ts.QpnPreAirship]
	if ok {
		preAirship, err = i.queryLookup[ts.QpnPreAirship].ParseBool(r)
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

func (i handlerInput[T, R, A]) queryAvlRepGetParams(r *http.Request) (avlParams, error) {
	params, errAvl := i.queryAvlGetParams(r)
	if errExceptEmptyQuery(errAvl) {
		return avlParams{}, errAvl
	}

	isRepeatable, errRepl := nullBool(r, ts.QpnRepeatable, i.queryLookup)
	if errExceptEmptyQuery(errRepl) {
		return avlParams{}, errRepl
	}

	if queryIsEmpty(errAvl) && queryIsEmpty(errRepl) {
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
