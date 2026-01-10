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

func getSectionList[T h.HasID, R any, L IsAPIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, L]) (SectionList, error) {
	sectionMap := i.subsections
	sectionURLs := []string{}

	for section := range sectionMap {
		url := cfg.createSectionURL(i.endpoint, section)
		sectionURLs = append(sectionURLs, url)
	}
	slices.Sort(sectionURLs)

	listParams, sections, err := createPaginatedList(cfg, r, i, sectionURLs)
	if err != nil {
		return SectionList{}, err
	}

	list := SectionList{
		ListParams: listParams,
		Results: sections,
	}

	return list, nil
}