package main

import (
	"net/http"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type SectionList struct {
	ListParams
	Results []string `json:"results"`
}

func (l SectionList) getListParams() ListParams {
	return l.ListParams
}

func getSectionList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (SectionList, error) {
	sectionMap := i.subsections
	sectionURLs := []string{}

	for section := range sectionMap {
		url := createSectionURL(cfg, i.endpoint, section)
		sectionURLs = append(sectionURLs, url)
	}
	slices.Sort(sectionURLs)

	listParams, sections, err := createPaginatedList(cfg, r, sectionURLs)
	if err != nil {
		return SectionList{}, err
	}

	list := SectionList{
		ListParams: listParams,
		Results:    sections,
	}

	return list, nil
}
