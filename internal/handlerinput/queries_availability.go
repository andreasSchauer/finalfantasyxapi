package handlerinput

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
	Availabilities []int32
	IsRepeatable   sql.NullBool
	PreAirship     bool
}

func (i Main[T, R, A]) QueryAvlGetParams(r *http.Request) (avlParams, error) {
	queryParamAvl := i.QueryLookup[ts.QpnAvailability]

	availabilities, err := queryParamAvl.ParseEnumList(r, enums.T.AvailabilityType)
	if errs.ErrExceptEmptyQuery(err) {
		return avlParams{}, err
	}
	if errs.QueryIsEmpty(err) {
		return avlParams{}, errs.ErrEmptyQuery
	}

	var preAirship bool

	_, ok := i.QueryLookup[ts.QpnPreAirship]
	if ok {
		preAirship, err = i.QueryLookup[ts.QpnPreAirship].ParseBool(r)
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
		Availabilities: h.SliceOrNil(avlRanks),
		PreAirship:     preAirship,
	}

	return params, nil
}

func (i Main[T, R, A]) QueryAvlRepGetParams(r *http.Request) (avlParams, error) {
	params, errAvl := i.QueryAvlGetParams(r)
	if errs.ErrExceptEmptyQuery(errAvl) {
		return avlParams{}, errAvl
	}

	isRepeatable, errRepl := NullBool(r, ts.QpnRepeatable, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(errRepl) {
		return avlParams{}, errRepl
	}

	if errs.QueryIsEmpty(errAvl) && errs.QueryIsEmpty(errRepl) {
		return avlParams{}, errs.ErrEmptyQuery
	}

	params.IsRepeatable = isRepeatable

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
