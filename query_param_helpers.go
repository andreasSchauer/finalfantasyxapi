package main

import (
	"cmp"
	"fmt"
	"net/http"
	"slices"
	"strings"
)

type QueryParameterList struct {
	ListParams
	Results []QueryType	`json:"results"`
}


func verifyQueryParams(r *http.Request, endpoint string, id *int32, lookup map[string]QueryType) error {
	q := r.URL.Query()

	for param := range q {
		queryType, ok := lookup[param]
		if !ok {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Parameter %s does not exist for endpoint %s.", param, endpoint), nil)
		}

		if queryType.RequiredWith != nil {
			for _, reqParam := range queryType.RequiredWith {
				reqParamVal := q.Get(reqParam)
				if reqParamVal == "" {
					return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used in combination with parameter(s): %s.", param, param, strings.Join(queryType.RequiredWith, ", ")), nil)
				}
			}
		}

		if queryType.ForSingle {
			if id == nil {
				return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used with single-resource-endpoints.", param, param), nil)
			}

			if queryType.AllowedIDs != nil {
				allowedIDPresent := false
				
				for _, reqID := range queryType.AllowedIDs {
					if *id == reqID {
						allowedIDPresent = true
					}
				}
				if !allowedIDPresent {
					idsString := strings.Trim(strings.Join(strings.Split(fmt.Sprint(queryType.AllowedIDs), " "), ", "), "[]")
					return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid id %d. Parameter %s can only be used with ids %s.", *id, param, idsString), nil)
				}
			}
		}

		if queryType.ForList && id != nil {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used with list-endpoints.", param, param), nil)
		}
	}

	return nil
}


func (cfg *Config) getQueryParamList (r *http.Request, lookup map[string]QueryType, endpoint string) (QueryParameterList, error) {
	queryParams := queryMapToSlice(lookup)
	queryParams, err := filterParamsOnSection(r, queryParams, endpoint)
	if err != nil {
		return QueryParameterList{}, err
	}

	listParams, shownResources, err := createPaginatedList(cfg, r, queryParams)
	if err != nil {
		return QueryParameterList{}, err
	}

	list := QueryParameterList{
		ListParams: listParams,
		Results: shownResources,
	}

	return list, nil
}

func filterParamsOnSection(r *http.Request, params []QueryType, endpoint string) ([]QueryType, error) {
	section := r.URL.Query().Get("section")

	if section == "" {
		return params, nil
	}

	filteredParams := []QueryType{}

	for _, param := range params {
		if len(param.ForSections) == 0 {
			filteredParams = append(filteredParams, param)
			continue
		}
		for _, sctn := range param.ForSections {
			if section == sctn {
				filteredParams = append(filteredParams, param)
			}
		}
	}

	if len(filteredParams) == 0 {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("section %s does not exist for endpoint %s.", section, endpoint), nil)
	}

	return filteredParams, nil
}


func queryMapToSlice(lookup map[string]QueryType) []QueryType {
	queryParams := []QueryType{}

	for key := range lookup {
		queryParams = append(queryParams, lookup[key])
	}

	slices.SortStableFunc(queryParams, func(a, b QueryType) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return queryParams
}