package api

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getRawQueryLimit(r *http.Request) (int, error) {
	queryParamLimit := Q.DefaultParams[ts.QpnLimit]
	query, err := queryParamLimit.CheckEmpty(r)
	if err != nil {
		return 0, err
	}

	limit, err := queryParamLimit.CheckInt(query)
	if err != nil {
		return 0, err
	}
	if limit == 0 {
		limit = *queryParamLimit.DefaultVal
	}
	if limit > config.Cfg.FetchLimit {
		return config.Cfg.FetchLimit, nil
	}

	return limit, nil
}

func getLimitMax() string {
	queryParamLimit := Q.DefaultParams[ts.QpnLimit]

	for _, input := range queryParamLimit.SpecialInputs {
		if input.Key == ts.QsvMax {
			return strconv.Itoa(input.Val)
		}
	}

	return ""
}

func setLimitMax(r *http.Request, q url.Values) {
	limit := q.Get(string(ts.QpnLimit))
	if limit != "" {
		return
	}

	q.Set(string(ts.QpnLimit), getLimitMax())
	r.URL.RawQuery = q.Encode()
}
