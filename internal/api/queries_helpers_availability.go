package api

import (
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type avlParams struct {
	availabilities 	[]int32
	isRepeatable	sql.NullBool
	preAirship		bool
}


func checkAvl[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (avlParams, error) {
	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errExceptEmptyQuery(err) {
		return avlParams{}, err
	}
	if queryIsEmpty(err) {
		return avlParams{}, errEmptyQuery
	}

	var preAirship bool

	_, ok := i.queryLookup["pre_airship"]
	if ok {
		preAirship, err = parseBooleanQuery(r, i.queryLookup["pre_airship"])
		if errExceptEmptyQuery(err) {
			return avlParams{}, err
		}
	}
	avlRanks := avlToRanks(availabilities, preAirship)

	params := avlParams{
		availabilities: h.SliceOrNil(avlRanks),
		preAirship: 	preAirship,
	}

	return params, nil
}


func checkAvlAndRep[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (avlParams, error) {
	params, errAvl := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(errAvl) {
		return avlParams{}, errAvl
	}

	isRepeatable, errRepl := getQueryBoolPtr(r, "repeatable", i.queryLookup)
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
		database.AvailabilityTypeAlways: 1,
		database.AvailabilityTypePost: 2,
		database.AvailabilityTypePreStory: 3,
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