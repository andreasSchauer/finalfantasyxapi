package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


type QueryParameterList struct {
	ListParams
	Results []QueryType `json:"results"`
}


func getQueryParamList[T h.HasID, R any, L IsAPIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, L]) (QueryParameterList, error) {
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
		Results:    shownResources,
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