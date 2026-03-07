package api

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
	sectionNames := []string{}

	for section := range sectionMap {
		sectionNames = append(sectionNames, section)
	}
	slices.Sort(sectionNames)

	listParams, sections, err := createPaginatedList(cfg, r, sectionNames)
	if err != nil {
		return SectionList{}, err
	}

	list := SectionList{
		ListParams: listParams,
		Results:    sections,
	}

	return list, nil
}
