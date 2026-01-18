package main

import (
	"errors"
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type QueryParameterList struct {
	ListParams
	Results []QueryType `json:"results"`
}

func getQueryParamList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (QueryParameterList, error) {
	section := r.URL.Query().Get("section")
	queryParams := queryMapToSlice(i.queryLookup)
	queryParams, err := filterParamsOnSection(queryParams, section, i)
	if err != nil {
		return QueryParameterList{}, err
	}

	queryParams = getAllowedResources(cfg, i, queryParams)

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

func getAllowedResources[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], params []QueryType) []QueryType {
	for idx, param := range params {
		for _, id := range param.AllowedIDs {
			allowedRes := i.idToResFunc(cfg, i, id)
			param.AllowedResources = append(param.AllowedResources, allowedRes)
		}
		params[idx] = param
	}

	return params
}

func filterParamsOnSection[T h.HasID, R any, A APIResource, L APIResourceList](params []QueryType, section string, i handlerInput[T, R, A, L]) ([]QueryType, error) {
	section, err := verifySectionParam(section, i.endpoint, i.subsections)
	if errors.Is(err, errEmptyQuery) {
		return params, nil
	}
	if err != nil {
		return nil, err
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

	return filteredParams, nil
}

func verifySectionParam(section, endpoint string, sectionMap map[string]func(string) (APIResourceList, error)) (string, error) {
	if section == "" {
		return "", errEmptyQuery
	}

	if section == "self" {
		return endpoint, nil
	}

	_, ok := sectionMap[section]
	if !ok {
		return "nil", newHTTPError(http.StatusBadRequest, fmt.Sprintf("subsection '%s' is not available for endpoint /%s.", section, endpoint), nil)
	}

	return section, nil
}
