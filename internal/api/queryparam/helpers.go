package queryparam

import (
	"cmp"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
)

func querySplit(query, sep string) []string {
	queryTrimmed := strings.TrimSuffix(query, sep)
	return strings.Split(queryTrimmed, sep)
}

func QueryListSplit(query string) ([]string, error) {
	segments := querySplit(query, ",")

	if len(segments) > config.Cfg.FetchLimit {
		return nil, errs.NewHTTPErrorFetchLimit(config.Cfg.FetchLimit)
	}

	return segments, nil
}

func queryIntMapToSlice(m map[string]int32) []int32 {
	items := []int32{}

	for _, item := range m {
		items = append(items, item)
	}

	slices.SortStableFunc(items, func(a, b int32) int {
		return cmp.Compare(a, b)
	})

	return items
}

// checks, if "none" was used as input and returns errQueryNone, if it was
func checkNoneQuery(query string) error {
	if query == "none" {
		return errs.ErrQueryNone
	}

	return nil
}
