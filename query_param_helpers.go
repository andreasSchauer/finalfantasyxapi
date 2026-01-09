package main

import (
	"cmp"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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

		err := verifyRequiredParams(q, queryType, param)
		if err != nil {
			return err
		}

		err = verifyQueryUsage(queryType, param, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func verifyRequiredParams(q url.Values, queryType QueryType, param string) error {
	if queryType.RequiredWith == nil {
		return nil
	}

	for _, reqParam := range queryType.RequiredWith {
		reqParamVal := q.Get(reqParam)
		if reqParamVal == "" {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used in combination with parameter(s): %s.", param, param, strings.Join(queryType.RequiredWith, ", ")), nil)
		}
	}

	return nil
}


func verifyQueryUsage(queryType QueryType, param string, id *int32) error {
	if queryType.ForSingle {
		if id == nil {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used with single-resource-endpoints.", param, param), nil)
		}

		err := verifyAllowedIDs(queryType, param, *id)
		if err != nil {
			return err
		}
	}

	if queryType.ForList && id != nil {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used with list-endpoints.", param, param), nil)
	}

	return nil
}


func verifyAllowedIDs(queryType QueryType, param string, id int32) error {
	if queryType.AllowedIDs == nil {
		return nil
	}

	allowedIDPresent := false
	
	for _, reqID := range queryType.AllowedIDs {
		if id == reqID {
			allowedIDPresent = true
		}
	}
	if !allowedIDPresent {
		idsString := getIDsString(queryType.AllowedIDs)
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid id %d. Parameter %s can only be used with ids %s.", id, param, idsString), nil)
	}

	return nil
}

func getIDsString(IDs []int32) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(IDs), " "), ", "), "[]")
}


func getQueryParamList[T h.HasID, R any, L IsAPIResourceList] (cfg *Config, r *http.Request, i handlerInput[T, R, L]) (QueryParameterList, error) {
	section := r.URL.Query().Get("section")
	queryParams := queryMapToSlice(i.queryLookup)
	queryParams, err := filterParamsOnSection(queryParams, section, i.endpoint)
	if err != nil {
		return QueryParameterList{}, err
	}

	listParams, shownResources, err := createPaginatedList(cfg, r, i, queryParams)
	if err != nil {
		return QueryParameterList{}, err
	}

	list := QueryParameterList{
		ListParams: listParams,
		Results: shownResources,
	}

	return list, nil
}

func filterParamsOnSection(params []QueryType, section, endpoint string) ([]QueryType, error) {
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


func queryIDsToSlice(query string, queryParam QueryType, maxID int) ([]int32, error) {
	idStrs := strings.Split(query, ",")
	ids := []int32{}

	for _, idStr := range idStrs {
		id, err := parseQueryIdVal(idStr, queryParam, maxID)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	ids = removeDuplicateIDs(ids)

	return ids, nil
}


func removeDuplicateIDs(ids []int32) []int32 {
	idMap := make(map[int32]bool)
	idsNew := []int32{}

	for _, item := range ids {
		idMap[item] = true
	}

	for id := range idMap {
		idsNew = append(idsNew, id)
	}

	return idsNew
}